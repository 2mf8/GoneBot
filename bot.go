package pbbot

import (
	"encoding/json"
	"errors"

	"github.com/2mf8/go-pbbot-for-rq/proto_gen/onebot"
	"github.com/2mf8/go-pbbot-for-rq/util"
	"github.com/fanliao/go-promise"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var Bots = make(map[int64]*Bot)

type Bot struct {
	BotId         int64
	Session       *SafeWebSocket
	WaitingFrames map[string]*promise.Promise
}

func NewBot(botId int64, conn *websocket.Conn) *Bot {
	messageHandler := func(messageType int, data []byte) {
		var frame onebot.Frame
		if messageType == websocket.BinaryMessage {
			err := proto.Unmarshal(data, &frame)
			if err != nil {
				log.Errorf("failed to unmarshal websocket binary message, err: %+v", err)
				return
			}
		} else if messageType == websocket.TextMessage {
			err := json.Unmarshal(data, &frame)
			if err != nil {
				log.Errorf("failed to unmarshal websocket text message, err: %+v", err)
				return
			}
		} else {
			log.Errorf("invalid websocket messageType: %+v", messageType)
			return
		}

		bot, ok := Bots[botId]
		if !ok {
			_ = conn.Close()
			return
		}
		util.SafeGo(func() {
			bot.handleFrame(&frame)
		})
	}
	closeHandler := func(code int, message string) {
		HandleDisconnect(Bots[botId])
		delete(Bots, botId)
	}
	safeWs := NewSafeWebSocket(conn, messageHandler, closeHandler)
	bot := &Bot{
		BotId:         botId,
		Session:       safeWs,
		WaitingFrames: make(map[string]*promise.Promise),
	}
	Bots[botId] = bot
	HandleConnect(bot)
	return bot
}

func (bot *Bot) handleFrame(frame *onebot.Frame) {
	if event := frame.GetPrivateMessageEvent(); event != nil {
		HandlePrivateMessage(bot, event)
		return
	}
	if event := frame.GetGroupMessageEvent(); event != nil {
		HandleGroupMessage(bot, event)
		return
	}
	if event := frame.GetGroupUploadNoticeEvent(); event != nil {
		HandleGroupUploadNotice(bot, event)
		return
	}
	if event := frame.GetGroupAdminNoticeEvent(); event != nil {
		HandleGroupAdminNotice(bot, event)
		return
	}
	if event := frame.GetGroupDecreaseNoticeEvent(); event != nil {
		HandleGroupDecreaseNotice(bot, event)
		return
	}
	if event := frame.GetGroupIncreaseNoticeEvent(); event != nil {
		HandleGroupIncreaseNotice(bot, event)
		return
	}
	if event := frame.GetGroupBanNoticeEvent(); event != nil {
		HandleGroupBanNotice(bot, event)
		return
	}
	if event := frame.GetFriendAddNoticeEvent(); event != nil {
		HandleFriendAddNotice(bot, event)
		return
	}
	if event := frame.GetFriendRecallNoticeEvent(); event != nil {
		HandleFriendRecallNotice(bot, event)
		return
	}
	if event := frame.GetGroupRecallNoticeEvent(); event != nil {
		HandleGroupRecallNotice(bot, event)
		return
	}
	if event := frame.GetFriendRequestEvent(); event != nil {
		HandleFriendRequest(bot, event)
		return
	}
	if event := frame.GetGroupRequestEvent(); event != nil {
		HandleGroupRequest(bot, event)
		return
	}

	if frame.FrameType < 300 {
		log.Errorf("unknown frame type: %+v", frame.FrameType)
		return
	}
	p, ok := bot.WaitingFrames[frame.Echo]
	if !ok {
		log.Errorf("failed to find waiting frame")
		return
	}
	if err := p.Resolve(frame); err != nil {
		log.Errorf("failed to resolve waiting frame promise")
		return
	}
}

func (bot *Bot) sendFrameAndWait(frame *onebot.Frame) (*onebot.Frame, error) {
	frame.BotId = bot.BotId
	frame.Echo = util.GenerateIdStr()
	frame.Ok = true
	data, err := proto.Marshal(frame)
	if err != nil {
		return nil, err
	}
	bot.Session.Send(websocket.BinaryMessage, data)
	p := promise.NewPromise()
	bot.WaitingFrames[frame.Echo] = p
	defer delete(bot.WaitingFrames, frame.Echo)
	resp, err, timeout := p.GetOrTimeout(120000)
	if err != nil || timeout {
		return nil, err
	}
	respFrame, ok := resp.(*onebot.Frame)
	if !ok {
		return nil, errors.New("failed to convert promise result to resp frame")
	}
	return respFrame, nil
}

func (bot *Bot) SendPrivateMessage(userId int64, msg *Msg, autoEscape bool) (*onebot.SendPrivateMsgResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSendPrivateMsgReq,
		Data: &onebot.Frame_SendPrivateMsgReq{
			SendPrivateMsgReq: &onebot.SendPrivateMsgReq{
				UserId:     userId,
				Message:    msg.MessageList,
				AutoEscape: autoEscape,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetSendPrivateMsgResp(), nil
	}
}

func (bot *Bot) SendGroupMessage(groupId int64, msg *Msg, autoEscape bool) (*onebot.SendGroupMsgResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSendGroupMsgReq,
		Data: &onebot.Frame_SendGroupMsgReq{
			SendGroupMsgReq: &onebot.SendGroupMsgReq{
				GroupId:    groupId,
				Message:    msg.MessageList,
				AutoEscape: autoEscape,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetSendGroupMsgResp(), nil
	}
}

func (bot *Bot) DeleteMsg(messageId []byte) (*onebot.DeleteMsgResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TDeleteMsgReq,
		Data: &onebot.Frame_DeleteMsgReq{
			DeleteMsgReq: &onebot.DeleteMsgReq{
				MessageId: messageId,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetDeleteMsgResp(), nil
	}
}

func (bot *Bot) GetMsg(messageId int32) (*onebot.GetMsgResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TGetMsgReq,
		Data: &onebot.Frame_GetMsgReq{
			GetMsgReq: &onebot.GetMsgReq{
				MessageId: messageId,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetGetMsgResp(), nil
	}
}

func (bot *Bot) SetGroupKick(groupId int64, userId int64, rejectAddRequest bool) (*onebot.SetGroupKickResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSetGroupKickReq,
		Data: &onebot.Frame_SetGroupKickReq{
			SetGroupKickReq: &onebot.SetGroupKickReq{
				GroupId:          groupId,
				UserId:           userId,
				RejectAddRequest: rejectAddRequest,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetSetGroupKickResp(), nil
	}
}

func (bot *Bot) SetGroupBan(groupId int64, userId int64, duration int32) (*onebot.SetGroupBanResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSetGroupBanReq,
		Data: &onebot.Frame_SetGroupBanReq{
			SetGroupBanReq: &onebot.SetGroupBanReq{
				GroupId:  groupId,
				UserId:   userId,
				Duration: duration,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetSetGroupBanResp(), nil
	}
}

func (bot *Bot) SetGroupWholeBan(groupId int64, enable bool) (*onebot.SetGroupWholeBanResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSetGroupWholeBanReq,
		Data: &onebot.Frame_SetGroupWholeBanReq{
			SetGroupWholeBanReq: &onebot.SetGroupWholeBanReq{
				GroupId: groupId,
				Enable:  enable,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetSetGroupWholeBanResp(), nil
	}
}

func (bot *Bot) SetGroupCard(groupId int64, userId int64, card string) (*onebot.SetGroupCardResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSetGroupCardReq,
		Data: &onebot.Frame_SetGroupCardReq{
			SetGroupCardReq: &onebot.SetGroupCardReq{
				GroupId: groupId,
				UserId:  userId,
				Card:    card,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetSetGroupCardResp(), nil
	}
}

func (bot *Bot) SetGroupLeave(groupId int64, isDismiss bool) (*onebot.SetGroupLeaveResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSetGroupLeaveReq,
		Data: &onebot.Frame_SetGroupLeaveReq{
			SetGroupLeaveReq: &onebot.SetGroupLeaveReq{
				GroupId:   groupId,
				IsDismiss: isDismiss,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetSetGroupLeaveResp(), nil
	}
}

func (bot *Bot) SetGroupSpecialTitle(groupId int64, userId int64, specialTitle string, duration int64) (*onebot.SetGroupSpecialTitleResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSetGroupSpecialTitleReq,
		Data: &onebot.Frame_SetGroupSpecialTitleReq{
			SetGroupSpecialTitleReq: &onebot.SetGroupSpecialTitleReq{
				GroupId:      groupId,
				UserId:       userId,
				SpecialTitle: specialTitle,
				Duration:     duration,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetSetGroupSpecialTitleResp(), nil
	}
}

func (bot *Bot) SetFriendAddRequest(flag string, approve bool, remark string) (*onebot.SetFriendAddRequestResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSetFriendAddRequestReq,
		Data: &onebot.Frame_SetFriendAddRequestReq{
			SetFriendAddRequestReq: &onebot.SetFriendAddRequestReq{
				Flag:    flag,
				Approve: approve,
				Remark:  remark,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetSetFriendAddRequestResp(), nil
	}
}

func (bot *Bot) SetGroupAddRequest(flag string, subType string, approve bool, reason string) (*onebot.SetGroupAddRequestResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSetGroupAddRequestReq,
		Data: &onebot.Frame_SetGroupAddRequestReq{
			SetGroupAddRequestReq: &onebot.SetGroupAddRequestReq{
				Flag:    flag,
				SubType: subType,
				Approve: approve,
				Reason:  reason,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetSetGroupAddRequestResp(), nil
	}
}

func (bot *Bot) GetLoginInfo() (*onebot.GetLoginInfoResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TGetLoginInfoReq,
		Data: &onebot.Frame_GetLoginInfoReq{
			GetLoginInfoReq: &onebot.GetLoginInfoReq{},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetGetLoginInfoResp(), nil
	}
}

func (bot *Bot) GetStrangerInfo(userId int64, noCache bool) (*onebot.GetStrangerInfoResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TGetStrangerInfoReq,
		Data: &onebot.Frame_GetStrangerInfoReq{
			GetStrangerInfoReq: &onebot.GetStrangerInfoReq{
				UserId:  userId,
				NoCache: noCache,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetGetStrangerInfoResp(), nil
	}
}

func (bot *Bot) GetFriendList() (*onebot.GetFriendListResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TGetFriendListReq,
		Data: &onebot.Frame_GetFriendListReq{
			GetFriendListReq: &onebot.GetFriendListReq{},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetGetFriendListResp(), nil
	}
}

func (bot *Bot) GetGroupList() (*onebot.GetGroupListResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TGetGroupListReq,
		Data: &onebot.Frame_GetGroupListReq{
			GetGroupListReq: &onebot.GetGroupListReq{},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetGetGroupListResp(), nil
	}
}

func (bot *Bot) GetGroupInfo(groupId int64, noCache bool) (*onebot.GetGroupInfoResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TGetGroupInfoReq,
		Data: &onebot.Frame_GetGroupInfoReq{
			GetGroupInfoReq: &onebot.GetGroupInfoReq{
				GroupId: groupId,
				NoCache: noCache,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetGetGroupInfoResp(), nil
	}
}

func (bot *Bot) GetGroupMemberInfo(groupId int64, userId int64, noCache bool) (*onebot.GetGroupMemberInfoResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TGetGroupMemberInfoReq,
		Data: &onebot.Frame_GetGroupMemberInfoReq{
			GetGroupMemberInfoReq: &onebot.GetGroupMemberInfoReq{
				GroupId: groupId,
				UserId:  userId,
				NoCache: noCache,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetGetGroupMemberInfoResp(), nil
	}
}

func (bot *Bot) GetGroupMemberList(groupId int64) (*onebot.GetGroupMemberListResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TGetGroupMemberListReq,
		Data: &onebot.Frame_GetGroupMemberListReq{
			GetGroupMemberListReq: &onebot.GetGroupMemberListReq{
				GroupId: groupId,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetGetGroupMemberListResp(), nil
	}
}
