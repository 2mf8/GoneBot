package pbbot

import (
	"strconv"

	"github.com/2mf8/go-pbbot-for-rq/proto_gen/onebot"
)

type Msg struct {
	MessageList []*onebot.Message
}

func NewMsg() *Msg {
	return &Msg{
		MessageList: make([]*onebot.Message, 0),
	}
}

func (msg *Msg) Text(text string) *Msg {
	msg.MessageList = append(msg.MessageList, &onebot.Message{
		Type: "text",
		Data: map[string]string{
			"text": text,
		},
	})
	return msg
}

func (msg *Msg) Face(id int) *Msg {
	msg.MessageList = append(msg.MessageList, &onebot.Message{
		Type: "face",
		Data: map[string]string{
			"id": strconv.Itoa(id),
		},
	})
	return msg
}

func (msg *Msg) Image(url string) *Msg {
	msg.MessageList = append(msg.MessageList, &onebot.Message{
		Type: "image",
		Data: map[string]string{
			"url": url,
		},
	})
	return msg
}

func (msg *Msg) At(qq int64, dispaly string) *Msg {
	msg.MessageList = append(msg.MessageList, &onebot.Message{
		Type: "at",
		Data: map[string]string{
			"qq": strconv.FormatInt(qq, 10),
			"display": dispaly,
		},
	})
	return msg
}

func (msg *Msg) LightApp(content string) *Msg{
	msg.MessageList = append(msg.MessageList, &onebot.Message{
		Type: "light_app",
		Data: map[string]string{
			"content": content,
		},
	})
	return msg
}
