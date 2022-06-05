package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/2mf8/go-pbbot-for-rq"
	"github.com/2mf8/go-pbbot-for-rq/proto_gen/onebot"
)

func TestBotServer(t *testing.T) {
	pbbot.HandleConnect = func(bot *pbbot.Bot) {
		fmt.Printf("新机器人已连接：%d\n", bot.BotId)
		fmt.Println("所有机器人列表：")
		for botId, _ := range pbbot.Bots {
			println(botId)
		}
	}

	pbbot.HandleGroupMessage = func(bot *pbbot.Bot, event *onebot.GroupMessageEvent) {
		rawMsg := event.RawMessage
		groupId := event.GroupId
		userId := event.UserId
		replyMsg := pbbot.NewMsg().Text("hello world").At(userId).Text("你发送了:" + rawMsg)
		_, _ = bot.SendGroupMessage(groupId, replyMsg, false)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if err := pbbot.UpgradeWebsocket(w, req); err != nil {
			fmt.Println("创建机器人失败")
		}
	})
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
	select {}
}
