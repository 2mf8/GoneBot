package pbbot

import (
	"strconv"

	"github.com/2mf8/GoPbBot/onebot"
)

type Msg struct {
	IMessageList []*onebot.IMessage
}

func NewMsg() *Msg {
	return &Msg{
		IMessageList: make([]*onebot.IMessage, 0),
	}
}

func (msg *Msg) Text(text string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "text",
		Data: map[string]string{
			"text": text,
		},
	})
	return msg
}

func (msg *Msg) Face(id int) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "face",
		Data: map[string]string{
			"id": strconv.Itoa(id),
		},
	})
	return msg
}

func (msg *Msg) Image(url string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "image",
		Data: map[string]string{
			"file": url,
		},
	})
	return msg
}

func (msg *Msg) At(qq int64, display string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "at",
		Data: map[string]string{
			"qq":      strconv.FormatInt(qq, 10),
			"display": display,
		},
	})
	return msg
}

func (msg *Msg) AtAll() *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "at",
		Data: map[string]string{
			"qq": "all",
		},
	})
	return msg
}

func (msg *Msg) Record(url string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "record",
		Data: map[string]string{
			"url": url,
		},
	})
	return msg
}

func (msg *Msg) LightApp(content string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "light_app",
		Data: map[string]string{
			"content": content,
		},
	})
	return msg
}

func (msg *Msg) TTS(text string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "tts",
		Data: map[string]string{
			"text": text,
		},
	})
	return msg
}

func (msg *Msg) Poke(qq int64) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "poke",
		Data: map[string]string{
			"qq": strconv.FormatInt(qq, 10),
		},
	})
	return msg
}

// GMC专用
func (msg *Msg) Reply(IMessageId int32) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "reply",
		Data: map[string]string{
			"IMessage_id": strconv.Itoa(int(IMessageId)),
		},
	})
	return msg
}

func (msg *Msg) Dice(value int64) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "dice",
		Data: map[string]string{
			"value": strconv.FormatInt(value, 10),
		},
	})
	return msg
}

func (msg *Msg) SignIn() *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "sign_in",
		Data: map[string]string{},
	})
	return msg
}

func (msg *Msg) Flash(url string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "image",
		Data: map[string]string{
			"url":  url,
			"type": "flash",
		},
	})
	return msg
}

func (msg *Msg) Share(url string, title string, content string, image string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "share",
		Data: map[string]string{
			"url":     url,
			"title":   title,
			"content": content,
			"image":   image,
		},
	})
	return msg
}

func (msg *Msg) Json(id int, content string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "service",
		Data: map[string]string{
			"sub_type": "json",
			"id":       strconv.Itoa(id),
			"content":  content,
		},
	})
	return msg
}

func (msg *Msg) Xml(id int, content string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "service",
		Data: map[string]string{
			"sub_type": "xml",
			"id":       strconv.Itoa(id),
			"content":  content,
		},
	})
	return msg
}

func (msg *Msg) Video(url string, cover string, cache bool) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "video",
		Data: map[string]string{
			"url":   url,
			"cover": cover,
			"cache": strconv.FormatBool(cache),
		},
	})
	return msg
}

func (msg *Msg) Sleep(time int64) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "sleep",
		Data: map[string]string{
			"time": strconv.FormatInt(time, 10),
		},
	})
	return msg
}

func (msg *Msg) ChannelAt(qq uint64, display string) *Msg {
	msg.IMessageList = append(msg.IMessageList, &onebot.IMessage{
		Type: "at",
		Data: map[string]string{
			"qq":      strconv.FormatUint(qq, 10),
			"display": display,
		},
	})
	return msg
}
