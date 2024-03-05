package pbbot

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/2mf8/GoPbBot/onebot"
	"github.com/2mf8/GoPbBot/util"
	"github.com/fanliao/go-promise"
	"github.com/gorilla/websocket"
	"github.com/jefferyjob/go-easy-utils/v2/anyUtil"
	log "github.com/sirupsen/logrus"
)

var Bots = make(map[int64]*Bot)
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
		if messageType == websocket.TextMessage {
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
	p, ok := bot.getWaitingFrame(frame.Echo)
	if !ok {
		return
	}
	if err := p.Resolve(frame); err != nil {
		log.Errorf("failed to resolve waiting frame promise")
		return
	}
}

func (bot *Bot) sendFrameAndWait(frame *onebot.Frame) (*onebot.Frame, error) {
	frame.BotId = bot.BotId
	frame.Echo = fmt.Sprintf("%v", time.Now().UnixNano())
	frame.Ok = true
	data, err := json.Marshal(frame)
	if err != nil {
		return nil, err
	}
	bot.Session.Send(websocket.BinaryMessage, data)
	p := promise.NewPromise()
	bot.setWaitingFrame(frame.Echo, p)
	defer bot.delWaitingFrame(frame.Echo)
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

func (bot *Bot) SendGroupMessage(groupId int64, msg *Msg, autoEscape bool) (*onebot.SendMsgResponse, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		API: &onebot.API{
			Action: string(onebot.SendGroupMsg),
			Params: &onebot.Params{
				GroupId:    groupId,
				Message:    msg.IMessageList,
				AutoEscape: autoEscape,
			},
			Echo: echo,
		},
	}); err != nil {
		return nil, err
	} else {
		i, err := anyUtil.AnyToInt32(resp.Data["message_id"])
		if err != nil {
			return nil, err
		}
		sr := &onebot.SendMsgResponse{
			Status:  anyUtil.AnyToStr(resp.Status),
			RetCode: resp.Code,
			Data: &onebot.SendMsgResponseData{
				MessageId: i,
				Echo:      resp.Echo,
			},
		}
		return sr, nil
	}
}

func (bot *Bot) SetGroupBan(groupId int64, userId int64, duration int32) (*onebot.SetGroupBanResp, error) {
	if _, err := bot.sendFrameAndWait(&onebot.Frame{
		API: &onebot.API{
			Action: string(onebot.SetGroupBan),
			Params: &onebot.Params{
				GroupId:  groupId,
				UserId:   userId,
				Duration: duration,
			},
			Echo: echo,
		},
	}); err != nil {
		return nil, err
	} else {
		sgbr := &onebot.SetGroupBanResp{}
		return sgbr, nil
	}
}

func (bot *Bot) SetGroupKick(groupId int64, userId int64, rejectAddRequest bool) (*onebot.SetGroupKickResp, error) {
	if _, err := bot.sendFrameAndWait(&onebot.Frame{
		API: &onebot.API{
			Action: string(onebot.SetGroupKick),
			Params: &onebot.Params{
				GroupId:          groupId,
				UserId:           userId,
				RejectAddRequest: rejectAddRequest,
			},
			Echo: echo,
		},
	}); err != nil {
		return nil, err
	} else {
		sgkr := &onebot.SetGroupKickResp{}
		return sgkr, nil
	}
}

func (bot *Bot) SetGroupLeave(groupId int64, isDismiss bool) (*onebot.SetGroupLeaveResp, error) {
	if _, err := bot.sendFrameAndWait(&onebot.Frame{
		API: &onebot.API{
			Action: string(onebot.SetGroupLeave),
			Params: &onebot.Params{
				GroupId:   groupId,
				IsDismiss: isDismiss,
			},
			Echo: echo,
		},
	}); err != nil {
		return nil, err
	} else {
		sglr := &onebot.SetGroupLeaveResp{}
		return sglr, nil
	}
}

func (bot *Bot) SetGroupWholeBan(groupId int64, enable bool) (*onebot.SetGroupWholeBanResp, error) {
	if _, err := bot.sendFrameAndWait(&onebot.Frame{
		API: &onebot.API{
			Action: string(onebot.SetGroupWholeBan),
			Params: &onebot.Params{
				GroupId: groupId,
				Enable:  enable,
			},
			Echo: echo,
		},
	}); err != nil {
		return nil, err
	} else {
		sgwbr := &onebot.SetGroupWholeBanResp{}
		return sgwbr, nil
	}
}

func (bot *Bot) DeleteMsg(msgId int32) (*onebot.DeleteMsgResp, error) {
	if _, err := bot.sendFrameAndWait(&onebot.Frame{
		API: &onebot.API{
			Action: string(onebot.DeleteMsg),
			Params: &onebot.Params{
				MessageId: msgId,
			},
			Echo: echo,
		},
	}); err != nil {
		return nil, err
	} else {
		dmr := &onebot.DeleteMsgResp{}
		return dmr, nil
	}
}

func (bot *Bot) GetGroupMemberInfo(groupId, userId int64, noCache bool) (*onebot.GetGroupMemberInfoResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		API: &onebot.API{
			Action: string(onebot.GetGroupMemberInfo),
			Params: &onebot.Params{
				GroupId: groupId,
				UserId:  userId,
				NoCache: noCache,
			},
			Echo: echo,
		},
	}); err != nil {
		return nil, err
	} else {
		rc, _ := anyUtil.AnyToInt32(resp.Code)
		gi, _ := anyUtil.AnyToInt64(resp.Data["group_id"])
		ui, _ := anyUtil.AnyToInt64(resp.Data["user_id"])
		ag, _ := anyUtil.AnyToInt32(resp.Data["age"])
		jt, _ := anyUtil.AnyToInt64(resp.Data["join_time"])
		lst, _ := anyUtil.AnyToInt64(resp.Data["last_sent_time"])
		tet, _ := anyUtil.AnyToInt64(resp.Data["title_expire_time"])
		ggmi := &onebot.GetGroupMemberInfoResp{
			Status:  anyUtil.AnyToStr(resp.Status),
			RetCode: rc,
			Data: &onebot.GroupMemberInfo{
				GroupId:         gi,
				UserId:          ui,
				Nickname:        anyUtil.AnyToStr(resp.Data["nickname"]),
				Card:            anyUtil.AnyToStr(resp.Data["card"]),
				Sex:             onebot.SexType(anyUtil.AnyToStr(resp.Data["sex"])),
				Age:             ag,
				Area:            anyUtil.AnyToStr(resp.Data["area"]),
				JoinTime:        jt,
				LastSentTime:    lst,
				Level:           anyUtil.AnyToStr(resp.Data["level"]),
				Role:            anyUtil.AnyToStr(resp.Data["role"]),
				UnFriendly:      anyUtil.AnyToBool(resp.Data["unfriendly"]),
				Title:           anyUtil.AnyToStr(resp.Data["title"]),
				TitleExpireTime: tet,
				CardChangeable:  anyUtil.AnyToBool(resp.Data["card_changeable"]),
			},
			Echo: anyUtil.AnyToStr(resp.Data["echo"]),
		}
		return ggmi, nil
	}
}

func (bot *Bot) GetGroupInfo(groupId int64, noCache bool) (*onebot.GetGroupInfoResp, error) {
	if resp, err := bot.sendFrameAndWait(&onebot.Frame{
		API: &onebot.API{
			Action: string(onebot.GetGroupInfo),
			Params: &onebot.Params{
				GroupId: groupId,
				NoCache: noCache,
			},
			Echo: echo,
		},
	}); err != nil {
		return nil, err
	} else {
		rc, _ := anyUtil.AnyToInt32(resp.Code)
		gi, _ := anyUtil.AnyToInt64(resp.Data["group_id"])
		mc, _ := anyUtil.AnyToInt32(resp.Data["member_count"])
		mmc, _ := anyUtil.AnyToInt32(resp.Data["max_member_count"])
		ggi := &onebot.GetGroupInfoResp{
			Status:  anyUtil.AnyToStr(resp.Status),
			RetCode: rc,
			Data: &onebot.GroupInfo{
				GroupId:        gi,
				GroupName:      anyUtil.AnyToStr(resp.Data["group_name"]),
				MemberCount:    mc,
				MaxMemberCount: mmc,
			},
			Echo: anyUtil.AnyToStr(resp.Data["echo"]),
		}
		return ggi, nil
	}
}
