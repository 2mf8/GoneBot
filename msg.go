package gonebot

import (
	"strconv"

	"github.com/2mf8/GoneBot/onebot"
)

type Msg struct {
	IMessageList []*onebot.IMessage
}

type MutiMsg struct {
	IMutiMsgList []*onebot.ForwardMsg
}

func NewMsg() *Msg {
	return &Msg{
		IMessageList: make([]*onebot.IMessage, 0),
	}
}

func (msg *Msg) AnyMsg(anyMsg []*onebot.IMessage) *Msg {
	msg.IMessageList = append(msg.IMessageList, anyMsg...)
	return msg
}

func (msg *Msg) LongMsg(id string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "longmsg",
		Data: map[string]any{
			"id": id,
		},
	})
	return msg
}

func (msg *Msg) Text(text string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "text",
		Data: map[string]any{
			"text": text,
		},
	})
	return msg
}

func (msg *Msg) Face(id int) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "face",
		Data: map[string]any{
			"id": strconv.Itoa(id),
		},
	})
	return msg
}

func (msg *Msg) Image(url string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "image",
		Data: map[string]any{
			"file": url,
		},
	})
	return msg
}

func (msg *Msg) At(qq int64, display string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "at",
		Data: map[string]any{
			"qq":      strconv.FormatInt(qq, 10),
		},
	})
	return msg
}

func (msg *Msg) AtAll() *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "at",
		Data: map[string]any{
			"qq": "all",
		},
	})
	return msg
}

func (msg *Msg) Record(url string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "record",
		Data: map[string]any{
			"file": url,
		},
	})
	return msg
}

func (msg *Msg) LightApp(content string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "light_app",
		Data: map[string]any{
			"content": content,
		},
	})
	return msg
}

func (msg *Msg) TTS(text string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "tts",
		Data: map[string]any{
			"text": text,
		},
	})
	return msg
}

func (msg *Msg) Poke(qq int64) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "poke",
		Data: map[string]any{
			"qq": strconv.FormatInt(qq, 10),
		},
	})
	return msg
}

// GMC专用
func (msg *Msg) Reply(msgId int32) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "reply",
		Data: map[string]any{
			"id": strconv.Itoa(int(msgId)),
		},
	})
	return msg
}

func (msg *Msg) Dice(value int64) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "dice",
		Data: map[string]any{},
	})
	return msg
}

func (msg *Msg) SignIn() *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "sign_in",
		Data: map[string]any{},
	})
	return msg
}

func (msg *Msg) Flash(url string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "image",
		Data: map[string]any{
			"url":  url,
			"type": "flash",
		},
	})
	return msg
}

func (msg *Msg) Share(url string, title string, content string, image string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "share",
		Data: map[string]any{
			"url":     url,
			"title":   title,
			"content": content,
			"image":   image,
		},
	})
	return msg
}

func (msg *Msg) Json(content string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "json",
		Data: map[string]any{
			"data": content,
		},
	})
	return msg
}

func (msg *Msg) Xml(content string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "xml",
		Data: map[string]any{
			"data": content,
		},
	})
	return msg
}

func (msg *Msg) Video(url string, cover string, cache bool) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "video",
		Data: map[string]any{
			"file":  url,
			"cover": cover,
			"cache": strconv.FormatBool(cache),
		},
	})
	return msg
}

func (msg *Msg) Sleep(time int64) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "sleep",
		Data: map[string]any{
			"time": strconv.FormatInt(time, 10),
		},
	})
	return msg
}

func (msg *Msg) Forward(id string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "forward",
		Data: map[string]any{
			"id": id,
		},
	})
	return msg
}

func (msg *Msg) Node(id string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "node",
		Data: map[string]any{
			"id": id,
		},
	})
	return msg
}

func NewForwardMsg() *MutiMsg {
	return &MutiMsg{
		IMutiMsgList: make([]*onebot.ForwardMsg, 0),
	}
}

func (mutiMsg *MutiMsg) Add(fMsg *onebot.ForwardMsg) *MutiMsg {
	mutiMsg.IMutiMsgList = append(mutiMsg.IMutiMsgList, fMsg)
	return mutiMsg
}
