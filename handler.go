package gonebot

import (
	"github.com/2mf8/GoneBot/onebot"
)

// HandleConnect 机器人连接
var HandleConnect = func(bot *Bot) {

}

// HandleDisconnect 机器人断开
var HandleDisconnect = func(bot *Bot) {

}

// HandlePrivateMessage 收到私聊消息
var HandlePrivateMessage = func(bot *Bot, event *onebot.PrivateMsgEvent) {

}

// HandleGroupMessage 收到群聊消息
var HandleGroupMessage = func(bot *Bot, event *onebot.GroupMsgEvent) {

}

// HandleGroupUploadNotice 有人上传群文件
var HandleGroupUploadNotice = func(bot *Bot, event *onebot.GroupUploadNoticeEvent) {

}

// HandleGroupAdminNotice 群管理员变动
var HandleGroupAdminNotice = func(bot *Bot, event *onebot.GroupAdminChangeNoticeEvent) {

}

// HandleGroupDecreaseNotice 群人数减少 有人退群或被踢
var HandleGroupDecreaseNotice = func(bot *Bot, event *onebot.GroupMemberDecreaseNoticeEvent) {

}

// HandleGroupIncreaseNotice 群人数增加
var HandleGroupIncreaseNotice = func(bot *Bot, event *onebot.GroupMemberIncreaseNoticeEvent) {

}

// HandleGroupBanNotice 有人被禁言
var HandleGroupBanNotice = func(bot *Bot, event *onebot.GroupBanNoticeEvent) {

}

// HandleFriendAddNotice 新好友添加
var HandleFriendAddNotice = func(bot *Bot, event *onebot.FriendAddNoticeEvent) {

}

// HandleGroupRecallNotice 群消息撤回
var HandleGroupRecallNotice = func(bot *Bot, event *onebot.GroupMsgRecallNoticeEvent) {

}

// HandleFriendRecallNotice 好友消息撤回
var HandleFriendRecallNotice = func(bot *Bot, event *onebot.FriendMsgRecallNoticeEvent) {

}

// HandleFriendRequest 收到好友请求
var HandleFriendRequest = func(bot *Bot, event *onebot.FriendAddRequestEvent) {

}

// HandleGroupRequest 收到加群请求
var HandleGroupRequest = func(bot *Bot, event *onebot.GroupAddOrInviteRequestEvent) {

}

// HandleGroupNotify 收到群通知
var HandleGroupMemberHonorChangeNotify = func(bot *Bot, event *onebot.GroupMemberHonorChangeNoticeEvent) {

}
var HandleGroupLuckyKingNotify = func(bot *Bot, event *onebot.GroupLuckyKingNoticeEvent) {

}
var HandleGroupPokeNotify = func(bot *Bot, event *onebot.GroupPokeNoticeEvent) {

}

// HandleLifeTime 收到生命周期
var HandleLifeTime = func(bot *Bot, event *onebot.LifeTime) {

}

// HandleHeartBeat 收到机器人心跳
var HandleHeartBeat = func(bot *Bot, event *onebot.BotHeartBeat) {

}
