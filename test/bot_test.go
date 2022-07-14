package test

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/2mf8/GoPbBot"
	"github.com/2mf8/GoPbBot/proto_gen/onebot"
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
		messageId := event.MessageId
		log.Println(rawMsg)
		if IsAdmin(bot, groupId, userId) && groupId == 706194673 {
			replyMsg := pbbot.NewMsg().Text(rawMsg)
			_, _ = bot.SendGroupMessage(groupId, replyMsg, false)
		}
	}
	http.HandleFunc("/ws/rq/", func(w http.ResponseWriter, req *http.Request) {
		if err := pbbot.UpgradeWebsocket(w, req); err != nil {
			fmt.Println("创建机器人失败")
		}
	})
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
	select {}
}

func IsAdmin(bot *pbbot.Bot, groupId, userId int64) bool {
	memberInfo, _ := bot.GetGroupMemberInfo(groupId, userId, true)
	if strings.ToLower(memberInfo.Role) == "admin" || strings.ToLower(memberInfo.Role) == "owner" {
		return true
	}
	return false
}