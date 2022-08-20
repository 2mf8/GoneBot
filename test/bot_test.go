package test

import (
	"fmt"
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

	pbbot.HandleGroupNotify = func(bot *pbbot.Bot, event *onebot.GroupNotifyEvent) {
		fmt.Println(event.GroupId, event.NoticeType, event.Sender, event.TargetId)
	}

	pbbot.HandleChannelMessage = func(bot *pbbot.Bot, event *onebot.ChannelMessageEvent) {
		guildId := event.GuildId
		channelId := event.ChannelId
		userId := event.Sender.TinyId
		display := event.Sender.Nickname
		rawMsg := event.RawMessage
		replyMsg := pbbot.NewMsg().ChannelAt(userId, display).Text("你发送了:" + rawMsg)
		_, err := bot.SendChannelMessage(guildId, channelId, replyMsg, false)
		fmt.Println(err)
	}

	pbbot.HandleGroupMessage = func(bot *pbbot.Bot, event *onebot.GroupMessageEvent) {
		rawMsg := event.RawMessage
		groupId := event.GroupId
		userId := event.UserId
		display := event.Sender.Card
		bot.DeleteMsg(event.MessageId)
		replyMsg := pbbot.NewMsg().Text("hello world").At(userId, display).Text("你发送了:" + rawMsg)
		_, _ = bot.SendGroupMessage(groupId, replyMsg, false)
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