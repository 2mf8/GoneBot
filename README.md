# GoneBot

[![QQ群](https://img.shields.io/static/v1?label=QQ%E7%BE%A4&message=901125207&color=blue)](http://qm.qq.com/cgi-bin/qm/qr?_wv=1027&k=tuBRPtCvexFULgMSjtty0RsbVk2HrCIj&authKey=JuIiTnyDbxaL2Ts6m7rY0kAZ0aJmyNETF97MmDSlLWME7Q4e%2F9%2Bg30r5l4zeWPiA&noverify=0&group_code=901125207)
[![MIT](https://img.shields.io/github/license/2mf8/GoneBot)](https://github.com/2mf8/GoneBot?tab=MIT-1-ov-file)
[![OneBot](https://img.shields.io/badge/OneBot-v11-black?style=social&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABABAMAAABYR2ztAAAAIVBMVEUAAAAAAAADAwMHBwceHh4UFBQNDQ0ZGRkoKCgvLy8iIiLWSdWYAAAAAXRSTlMAQObYZgAAAQVJREFUSMftlM0RgjAQhV+0ATYK6i1Xb+iMd0qgBEqgBEuwBOxU2QDKsjvojQPvkJ/ZL5sXkgWrFirK4MibYUdE3OR2nEpuKz1/q8CdNxNQgthZCXYVLjyoDQftaKuniHHWRnPh2GCUetR2/9HsMAXyUT4/3UHwtQT2AggSCGKeSAsFnxBIOuAggdh3AKTL7pDuCyABcMb0aQP7aM4AnAbc/wHwA5D2wDHTTe56gIIOUA/4YYV2e1sg713PXdZJAuncdZMAGkAukU9OAn40O849+0ornPwT93rphWF0mgAbauUrEOthlX8Zu7P5A6kZyKCJy75hhw1Mgr9RAUvX7A3csGqZegEdniCx30c3agAAAABJRU5ErkJggg==)](https://onebot.dev/)
[![GoneBot](https://img.shields.io/badge/GoneBot-SDK-black?style=social&logo=data:image/png;base64,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)](https://github.com/2mf8/GoneBot)

使用方法

```go
package main

import (
	"fmt"

	"github.com/2mf8/GoneBot"
	"github.com/2mf8/GoneBot/onebot"
	"github.com/gin-gonic/gin"
)

func main() {

	gonebot.HandleConnect = func(bot *gonebot.Bot) {
		fmt.Printf("新机器人已连接：%d\n", bot.BotId)
		fmt.Println("所有机器人列表：")
		for botId, _ := range gonebot.Bots {
			println(botId)
		}
	}

	gonebot.HandleGroupMessage = func(bot *gonebot.Bot, event *onebot.GroupMsgEvent) {
		rawMsg := event.RawMessage
		groupId := event.GroupId
		userId := event.UserId
		display := event.Sender.Card
		replyMsg := gonebot.NewMsg().Text("hello world").At(userId, display).Text("你发送了:" + rawMsg)
		_, _ = bot.SendGroupMessage(groupId, replyMsg, false)
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		if err := gonebot.UpgradeWebsocket(c.Writer, c.Request); err != nil {
			fmt.Println("创建机器人失败")
		}
	})

	if err := router.Run(":8081"); err != nil {
		panic(err)
	}
	select {}
}
```