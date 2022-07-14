# GoPbBot

使用方法

```go
package main

import (
	"fmt"

	"github.com/2mf8/GoPbBot"
	"github.com/2mf8/GoPbBot/proto_gen/onebot"
	"github.com/gin-gonic/gin"
)

func main() {

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
		display := event.Sender.Card
		replyMsg := pbbot.NewMsg().Text("hello world").At(userId, display).Text("你发送了:" + rawMsg)
		_, _ = bot.SendGroupMessage(groupId, replyMsg, false)
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		if err := pbbot.UpgradeWebsocket(c.Writer, c.Request); err != nil {
			fmt.Println("创建机器人失败")
		}
	})

	if err := router.Run(":8081"); err != nil {
		panic(err)
	}
	select {}
}
```