package pbbot

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/2mf8/GoPbBot/onebot"
	"github.com/2mf8/GoPbBot/util"
	"github.com/fanliao/go-promise"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var Bots = make(map[int64]*Bot)
var mresp = &onebot.SendMsgResponse{}
var echo = ""

type Bot struct {
	BotId         int64
	Session       *SafeWebSocket
	mux           sync.RWMutex
	WaitingFrames map[string]*promise.Promise
}

func NewBot(botId int64, conn *websocket.Conn) *Bot {
	messageHandler := func(messageType int, data []byte) {
		var frame onebot.Frame
		nullframe := onebot.Frame{}
		if messageType == websocket.TextMessage {
			err := json.Unmarshal(data, &frame)
			if err != nil {
				log.Errorf("failed to unmarshal websocket text message, err: %+v", err)
				return
			}
			if frame == nullframe {
				err = json.Unmarshal(data, &mresp)
				if err != nil {
					log.Errorf("failed to unmarshal websocket text message, err: %+v", err)
					return
				}
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
			bot.handleFrame(&frame, data)
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

func (bot *Bot) setWaitingFrame(key string, value *promise.Promise) {
	bot.mux.Lock()
	defer bot.mux.Unlock()
	bot.WaitingFrames[key] = value
}

func (bot *Bot) getWaitingFrame(key string) (*promise.Promise, bool) {
	bot.mux.RLock()
	defer bot.mux.RUnlock()
	value, ok := bot.WaitingFrames[key]
	return value, ok
}

func (bot *Bot) delWaitingFrame(key string) {
	bot.mux.Lock()
	defer bot.mux.Unlock()
	delete(bot.WaitingFrames, key)
}

func (bot *Bot) handleFrame(frame *onebot.Frame, data []byte) {
	if frame.PostType == onebot.Message && frame.MessageType == string(onebot.Private) {
		pm := &onebot.PrivateMsgEvent{}
		err := json.Unmarshal(data, pm)
		fmt.Println(err)
		if err == nil {
			HandlePrivateMessage(bot, pm)
		}
		return
	}
	if frame.PostType == onebot.Message && frame.MessageType == string(onebot.Group) {
		gm := &onebot.GroupMsgEvent{}
		err := json.Unmarshal(data, gm)
		fmt.Println(err)
		if err == nil {
			HandleGroupMessage(bot, gm)
		}
		return
	}
	p, ok := bot.getWaitingFrame(echo)
	if !ok {
		return
	}
	if err := p.Resolve(frame); err != nil {
		log.Errorf("failed to resolve waiting frame promise")
		return
	}
}

func (bot *Bot) sendFrameAndWait(frame *onebot.Frame, api *onebot.API) (*onebot.SendMsgResponse, error) {
	frame.BotId = bot.BotId
	echo = fmt.Sprintf("%v", time.Now().UnixNano())
	frame.Ok = true
	data, err := json.Marshal(api)
	if err != nil {
		return nil, err
	}
	bot.Session.Send(websocket.BinaryMessage, data)
	p := promise.NewPromise()
	bot.setWaitingFrame(echo, p)
	defer bot.delWaitingFrame(echo)
	_, err, timeout := p.GetOrTimeout(120000)
	if err != nil || timeout {
		return nil, err
	}
	if mresp.Status == "ok" {
		return mresp, nil
	}
	return nil, nil
}

func (bot *Bot) SendGroupMessage(groupId int64, msg *Msg, autoEscape bool) (*onebot.SendMsgResponse, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{}, &onebot.API{
		Action: string(onebot.SendGroupMsg),
		Params: &onebot.Params{
			GroupId:    groupId,
			Message:    msg.IMessageList,
			AutoEscape: autoEscape,
		},
		Echo: echo,
	}); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (bot *Bot) SetGroupBan(groupId int64, userId int64, duration int32) (*onebot.SendMsgResponse, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{}, &onebot.API{
		Action: string(onebot.SetGroupBan),
		Params: &onebot.Params{
			GroupId:  groupId,
			UserId:   userId,
			Duration: duration,
		},
		Echo: echo,
	}); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (bot *Bot) SetGroupKick(groupId int64, userId int64, rejectAddRequest bool) (*onebot.SendMsgResponse, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{}, &onebot.API{
		Action: string(onebot.SetGroupKick),
		Params: &onebot.Params{
			GroupId:          groupId,
			UserId:           userId,
			RejectAddRequest: rejectAddRequest,
		},
		Echo: echo,
	}); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (bot *Bot) SetGroupLeave(groupId int64, isDismiss bool) (*onebot.SendMsgResponse, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{}, &onebot.API{
		Action: string(onebot.SetGroupLeave),
		Params: &onebot.Params{
			GroupId:   groupId,
			IsDismiss: isDismiss,
		},
		Echo: echo,
	}); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (bot *Bot) SetWholeBan(groupId int64, enable bool) (*onebot.SendMsgResponse, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{}, &onebot.API{
		Action: string(onebot.SetGroupWholeBan),
		Params: &onebot.Params{
			GroupId: groupId,
			Enable:  enable,
		},
		Echo: echo,
	}); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (bot *Bot) DeleteMsg(msgId int32) (*onebot.SendMsgResponse, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{}, &onebot.API{
		Action: string(onebot.DeleteMsg),
		Params: &onebot.Params{
			MessageId: msgId,
		},
		Echo: echo,
	}); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

/*
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
	if event := frame.GetChannelMessageEvent(); event != nil {
		HandleChannelMessage(bot, event)
		return
	}
	if event := frame.GetGroupNotifyEvent(); event != nil {
		HandleGroupNotify(bot, event)
		return
	}
	if event := frame.GetGroupTempMessageEvent(); event != nil {
		HandleGroupTempMessage(bot, event)
		return
	}

	if frame.FrameType < 300 {
		log.Errorf("unknown frame type: %+v", frame.FrameType)
		return
	}
	p, ok := bot.getWaitingFrame(frame.Echo)
	if !ok {
		log.Errorf("failed to find waiting frame")
		return
	}
	if err := p.Resolve(frame); err != nil {
		log.Errorf("failed to resolve waiting frame promise")
		return
	}
}
/*
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



// GMC专用
func (bot *Bot) DeleteMsg(messageId int32) (*onebot.DeleteMsgResp, error) {
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

// GMC 1.1.0 以上版本和pbrq皆可用
func (bot *Bot) DeleteMsgByReceipt(messageReceipt *onebot.MessageReceipt) (*onebot.DeleteMsgResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TDeleteMsgReq,
		Data: &onebot.Frame_DeleteMsgReq{
			DeleteMsgReq: &onebot.DeleteMsgReq{
				MessageReceipt: messageReceipt,
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

// rq 设置不了
func (bot *Bot) SetGroupSpecialTitle(groupId int64, userId int64, specialTitle string) (*onebot.SetGroupSpecialTitleResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSetGroupSpecialTitleReq,
		Data: &onebot.Frame_SetGroupSpecialTitleReq{
			SetGroupSpecialTitleReq: &onebot.SetGroupSpecialTitleReq{
				GroupId:      groupId,
				UserId:       userId,
				SpecialTitle: specialTitle,
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

func (bot *Bot) SetGroupSignIn(groupId int64) (*onebot.SetGroupSignInResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSetGroupSignInReq,
		Data: &onebot.Frame_SetGroupSignInReq{
			SetGroupSignInReq: &onebot.SetGroupSignInReq{
				GroupId: groupId,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetSetGroupSignInResp(), nil
	}
}

func (bot *Bot) SendGroupPoke(toUin, groupId int64) (*onebot.SendGroupPokeResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSendGroupPokeReq,
		Data: &onebot.Frame_SendGroupPokeReq{
			SendGroupPokeReq: &onebot.SendGroupPokeReq{
				ToUin:   toUin,
				GroupId: groupId,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetSendGroupPokeResp(), nil
	}
}

func (bot *Bot) SendFriendPoke(toUin int64) (*onebot.SendFriendPokeResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSendFriendPokeReq,
		Data: &onebot.Frame_SendFriendPokeReq{
			SendFriendPokeReq: &onebot.SendFriendPokeReq{
				ToUin: toUin,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetSendFriendPokeResp(), nil
	}
}

func (bot *Bot) SendChannelMessage(guildId, channelId uint64, msg *Msg, autoEscape bool) (*onebot.SendChannelMsgResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSendChannelMsgReq,
		Data: &onebot.Frame_SendChannelMsgReq{
			SendChannelMsgReq: &onebot.SendChannelMsgReq{
				GuildId:    guildId,
				ChannelId:  channelId,
				Message:    msg.MessageList,
				AutoEscape: autoEscape,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetSendChannelMsgResp(), nil
	}
}

/*
	 *  发送群音乐
     *
     *  groupId 群号
     *  musicType 音乐类型 qq、164、migu、kugou、kuwo
     *  title 标题
     *  brief 简介
     *  summary 概览
     *  url 链接
     *  pictureUrl 图片链接
     *  musicUrl 音乐链接
*/ /*
func (bot *Bot) SendGroupMusic(groupId int64, musicType string, title string, brief string, summary string, url string, pictureUrl string, musicUrl string) (*onebot.SendMusicResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSendMusicReq,
		Data: &onebot.Frame_SendMusicReq{
			SendMusicReq: &onebot.SendMusicReq{
				GroupId:    groupId,
				Type:       musicType,
				Title:      title,
				Brief:      brief,
				Summary:    summary,
				Url:        url,
				PictureUrl: pictureUrl,
				MusicUrl:   musicUrl,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetSendMusicResp(), nil
	}
}

/*
	 *  发送好友音乐
     *
     *  userId 用户Id
     *  musicType 音乐类型 qq、164、migu、kugou、kuwo
     *  title 标题
     *  brief 简介
     *  summary 概览
     *  url 链接
     *  pictureUrl 图片链接
     *  musicUrl 音乐链接
*/ /*
func (bot *Bot) SendFriendMusic(userId int64, musicType string, title string, brief string, summary string, url string, pictureUrl string, musicUrl string) (*onebot.SendMusicResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		FrameType: onebot.Frame_TSendMusicReq,
		Data: &onebot.Frame_SendMusicReq{
			SendMusicReq: &onebot.SendMusicReq{
				UserId:     userId,
				Type:       musicType,
				Title:      title,
				Brief:      brief,
				Summary:    summary,
				Url:        url,
				PictureUrl: pictureUrl,
				MusicUrl:   musicUrl,
			},
		},
	}); err != nil {
		return nil, err
	} else {
		return resp.GetSendMusicResp(), nil
	}
}*/
