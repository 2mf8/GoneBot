package pbbot

import (
	"github.com/2mf8/GoPbBot/proto_gen/onebot"
)

// HandleConnect 机器人连接
var HandleConnect = func(bot *Bot) {

}

// HandleDisconnect 机器人断开
var HandleDisconnect = func(bot *Bot) {

}

// HandlePrivateMessage 收到私聊消息
var HandlePrivateMessage = func(bot *Bot, event *onebot.PrivateMessageEvent) {

}

// HandleGroupMessage 收到群聊消息
var HandleGroupMessage = func(bot *Bot, event *onebot.GroupMessageEvent) {

}

// HandleGroupUploadNotice 有人上传群文件
var HandleGroupUploadNotice = func(bot *Bot, event *onebot.GroupUploadNoticeEvent) {

}

// HandleGroupAdminNotice 群管理员变动
var HandleGroupAdminNotice = func(bot *Bot, event *onebot.GroupAdminNoticeEvent) {

}

// HandleGroupDecreaseNotice 群人数减少 有人退群或被踢
var HandleGroupDecreaseNotice = func(bot *Bot, event *onebot.GroupDecreaseNoticeEvent) {

}

// HandleGroupIncreaseNotice 群人数增加
var HandleGroupIncreaseNotice = func(bot *Bot, event *onebot.GroupIncreaseNoticeEvent) {

}

// HandleGroupBanNotice 有人被禁言
var HandleGroupBanNotice = func(bot *Bot, event *onebot.GroupBanNoticeEvent) {

}

// HandleFriendAddNotice 新好友添加
var HandleFriendAddNotice = func(bot *Bot, event *onebot.FriendAddNoticeEvent) {

}

// HandleGroupRecallNotice 群消息撤回
var HandleGroupRecallNotice = func(bot *Bot, event *onebot.GroupRecallNoticeEvent) {

}

// HandleFriendRecallNotice 好友消息撤回
var HandleFriendRecallNotice = func(bot *Bot, event *onebot.FriendRecallNoticeEvent) {

}

// HandleFriendRequest 收到好友请求
var HandleFriendRequest = func(bot *Bot, event *onebot.FriendRequestEvent) {

}

// HandleGroupRequest 收到加群请求
var HandleGroupRequest = func(bot *Bot, event *onebot.GroupRequestEvent) {

}
