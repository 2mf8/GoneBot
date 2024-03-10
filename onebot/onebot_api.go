package onebot

type ActionType string

const (
	SendPrivateMsg       ActionType = "send_private_msg"
	SendGroupMsg         ActionType = "send_group_msg"
	SendMsg              ActionType = "send_msg"
	SendForwardMsg       ActionType = "send_forward_msg"
	DeleteMsg            ActionType = "delete_msg"
	GetMsg               ActionType = "get_msg"
	GetForwardMsg        ActionType = "get_forward_msg"
	SendLike             ActionType = "send_like"
	SetGroupKick         ActionType = "set_group_kick"
	SetGroupBan          ActionType = "set_group_ban"
	SetGroupAnonymousBan ActionType = "set_group_anonymous_ban"
	SetGroupWholeBan     ActionType = "set_group_whole_ban"
	SetGroupAdmin        ActionType = "set_group_admin"
	SetGroupAnonymous    ActionType = "set_group_anonymous"
	SetGroupCard         ActionType = "set_group_card"
	SetGroupName         ActionType = "set_group_name"
	SetGroupLeave        ActionType = "set_group_leave"
	SetGroupSpecialTitle ActionType = "set_group_special_title"
	SetFriendAddRequest  ActionType = "set_friend_add_request"
	SetGroupAddRequest   ActionType = "set_group_add_request"
	GetLoginInfo         ActionType = "get_login_info"
	GetStrangerInfo      ActionType = "get_stranger_info"
	GetFriendList        ActionType = "get_friend_list"
	GetGroupInfo         ActionType = "get_group_info"
	GetGroupList         ActionType = "get_group_list"
	GetGroupMemberInfo   ActionType = "get_group_member_info"
	GetGroupMemberList   ActionType = "get_group_member_list"
	GetGroupHonorInfo    ActionType = "get_group_honor_info"
	GetGookies           ActionType = "get_cookies"
	GetGsrfToken         ActionType = "get_csrf_token"
	GetGredentials       ActionType = "get_credentials"
	GetGecord            ActionType = "get_record"
	GetImage             ActionType = "get_image"
	CanSendImage         ActionType = "can_send_image"
	CanSendRecord        ActionType = "can_send_record"
	GetStatus            ActionType = "get_status"
	GetVersioInfo        ActionType = "get_version_info"
	SetRestart           ActionType = "set_restart"
	CleanCache           ActionType = "clean_cache"
)

type API struct {
	Action string  `json:"action,omitempty"`
	Params *Params `json:"params,omitempty"`
	Echo   string  `json:"echo,omitempty"`
}

type Params struct {
	UserId           int64       `json:"user_id,omitempty"`
	GroupId          int64       `json:"group_id,omitempty"`
	Message          []*IMessage `json:"message,omitempty"`
	MessageType      string      `json:"message_type,omitempty"`
	AutoEscape       bool        `json:"auto_escape,omitempty"`
	MessageId        int64       `json:"message_id,omitempty"`
	Id               string      `json:"id,omitempty"`
	RejectAddRequest bool        `json:"reject_add_request,omitempty"`
	Duration         int64       `json:"duration,omitempty"`
	Enable           bool        `json:"enable,omitempty"`
	Card             string      `json:"card,omitempty"`
	GroupName        string      `json:"group_name,omitempty"`
	Flag             string      `json:"flag,omitempty"`
	Approve          bool        `json:"approve,omitempty"`
	Remark           string      `json:"remark,omitempty"`
	IsDismiss        bool        `json:"is_dismiss,omitempty"`
	SpecialTitle     string      `json:"special_title,omitempty"`
	SubType          string      `json:"sub_type,omitempty"`
	Type             string      `json:"type,omitempty"`
	Reason           string      `json:"reason,omitempty"`
	Nickname         string      `json:"nickname,omitempty"`
	NoCache          bool        `json:"no_cache,omitempty"`
	Domain           string      `json:"domain,omitempty"`
	File             string      `json:"file,omitempty"`
	Delay            int32       `json:"delay,omitempty"`
	OutFormat        string      `json:"out_format,omitempty"`
	Times            int32       `json:"times,omitempty"`
	Anonymous        *Anonymous  `json:"anonymous,omitempty"`
	AnonymousFlag    string      `json:"anonymous_flag,omitempty"`
}

type SendMsgResponse struct {
	Status  string               `json:"status,omitempty"`
	RetCode int32                `json:"retcode,omitempty"`
	Data    *SendMsgResponseData `json:"data,omitempty"`
	Echo    string               `json:"echo,omitempty"`
}
type SendMsgResponseData struct {
	MessageId int64 `json:"message_id,omitempty"`
}

type GetGroupMemberInfoResp struct {
	Status  string           `json:"status,omitempty"`
	RetCode int32            `json:"retcode,omitempty"`
	Data    *GroupMemberInfo `json:"data,omitempty"`
	Echo    string           `json:"echo,omitempty"`
}

type GroupMemberInfo struct {
	GroupId         int64   `json:"group_id,omitempty"`
	UserId          int64   `json:"user_id,omitempty"`
	Nickname        string  `json:"nickname,omitempty"`
	Card            string  `json:"card,omitempty"`
	Sex             SexType `json:"sex,omitempty"`
	Age             int32   `json:"age,omitempty"`
	Area            string  `json:"area,omitempty"`
	JoinTime        int64   `json:"join_time,omitempty"`
	LastSentTime    int64   `json:"last_sent_time,omitempty"`
	Level           string  `json:"level,omitempty"`
	Role            string  `json:"role,omitempty"`
	UnFriendly      bool    `json:"unfriendly,omitempty"`
	Title           string  `json:"title,omitempty"`
	TitleExpireTime int64   `json:"title_expire_time,omitempty"`
	CardChangeable  bool    `json:"card_changeable,omitempty"`
}

type GetGroupInfoResp struct {
	Status  string     `json:"status,omitempty"`
	RetCode int32      `json:"retcode,omitempty"`
	Data    *GroupInfo `json:"data,omitempty"`
	Echo    string     `json:"echo,omitempty"`
}

type GroupInfo struct {
	GroupId        int64  `json:"group_id,omitempty"`
	GroupName      string `json:"group_name,omitempty"`
	MemberCount    int32  `json:"member_count,omitempty"`
	MaxMemberCount int32  `json:"max_member_count,omitempty"`
}

type SetGroupBanResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    any    `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type SetGroupKickResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    any    `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type SetGroupLeaveResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    any    `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type SetGroupWholeBanResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    any    `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type DeleteMsgResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    any    `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type GetMsgResp struct {
	Status  string   `json:"status,omitempty"`
	RetCode int32    `json:"retcode,omitempty"`
	Data    *IGetMsg `json:"data,omitempty"`
	Echo    string   `json:"echo,omitempty"`
}

type IGetMsg struct {
	Time        int64        `json:"time,omitempty"`
	MessageType string       `json:"message_type,omitempty"`
	MessageId   int64        `json:"message_id,omitempty"`
	RealId      int64        `json:"real_id,omitempty"`
	Sender      *GroupSender `json:"sender,omitempty"`
	Message     []*IMessage  `json:"message,omitempty"`
}

type GetForwardMsgResp struct {
	Status  string      `json:"status,omitempty"`
	RetCode int32       `json:"retcode,omitempty"`
	Data    []*IMessage `json:"data,omitempty"`
	Echo    string      `json:"echo,omitempty"`
}

type SendLikeResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    any    `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type SetGroupAnonymousBanResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    any    `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type SetGroupAdminResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    any    `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type SetGroupAnonymousResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    any    `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type SetGroupCardResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    any    `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type SetGroupNameResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    any    `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type SetGroupSpecialTitleResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    any    `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type SetFriendAddRequestResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    any    `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type SetGroupAddRequestResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    any    `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type GetLoginInfoResp struct {
	Status  string         `json:"status,omitempty"`
	RetCode int32          `json:"retcode,omitempty"`
	Data    *IGetLoginInfo `json:"data,omitempty"`
	Echo    string         `json:"echo,omitempty"`
}

type IGetLoginInfo struct {
	UserId   int64  `json:"user_id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
}

type GetStrangerInfoResp struct {
	Status  string            `json:"status,omitempty"`
	RetCode int32             `json:"retcode,omitempty"`
	Data    *IGetStrangerInfo `json:"data,omitempty"`
	Echo    string            `json:"echo,omitempty"`
}

type IGetStrangerInfo struct {
	UserId   int64   `json:"user_id,omitempty"`
	Nickname string  `json:"nickname,omitempty"`
	Sex      SexType `json:"sex,omitempty"`
	Age      int32   `json:"age,omitempty"`
}

type GetFriendListResp struct {
	Status  string          `json:"status,omitempty"`
	RetCode int32           `json:"retcode,omitempty"`
	Data    *IGetFriendList `json:"data,omitempty"`
	Echo    string          `json:"echo,omitempty"`
}

type IGetFriendList struct {
	UserId   int64  `json:"user_id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Remark   string `json:"remark,omitempty"`
}

type GetGroupListResp struct {
	Status  string       `json:"status,omitempty"`
	RetCode int32        `json:"retcode,omitempty"`
	Data    []*GroupInfo `json:"data,omitempty"`
	Echo    string       `json:"echo,omitempty"`
}

type GetGroupMemberListResp struct {
	Status  string             `json:"status,omitempty"`
	RetCode int32              `json:"retcode,omitempty"`
	Data    []*GroupMemberInfo `json:"data,omitempty"`
	Echo    string             `json:"echo,omitempty"`
}

type GetGroupHonorInfoResp struct {
	Status  string              `json:"status,omitempty"`
	RetCode int32               `json:"retcode,omitempty"`
	Data    *IGetGroupHonorInfo `json:"data,omitempty"`
	Echo    string              `json:"echo,omitempty"`
}

type IGetGroupHonorInfo struct {
	GroupId          int64             `json:"group_id,omitempty"`
	CurrentTalkative *CurrentTalkative `json:"current_talkative,omitempty"`  // 当前龙王，仅 type 为 talkative 或 all 时有数据
	TalkativeList    *OtherList        `json:"talkative_list,omitempty"`     // 历史龙王，仅 type 为 talkative 或 all 时有数据
	PerformerList    *OtherList        `json:"performer_list,omitempty"`     // 群聊之火，仅 type 为 performer 或 all 时有数据
	LegendList       *OtherList        `json:"legend_list,omitempty"`        // 群聊炽焰，仅 type 为 legend 或 all 时有数据
	StrongNewbieList *OtherList        `json:"strong_newbie_list,omitempty"` // 冒尖小春笋，仅 type 为 strong_newbie 或 all 时有数据
	EmotionList      *OtherList        `json:"emotion_list,omitempty"`       // 快乐之源，仅 type 为 emotion 或 all 时有数据
}

type CurrentTalkative struct {
	UserId   int64  `json:"user_id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	DayCount int32  `json:"day_count,omitempty"`
}

type OtherList struct {
	UserId      int64  `json:"user_id,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	Description string `json:"description,omitempty"`
}

type GetCookiesResp struct {
	Status  string   `json:"status,omitempty"`
	RetCode int32    `json:"retcode,omitempty"`
	Data    *Cookies `json:"data,omitempty"`
	Echo    string   `json:"echo,omitempty"`
}

type Cookies struct {
	Cookies string `json:"cookies,omitempty"`
}

type GetCRSFTokenResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    *Token `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type Token struct {
	Token int32 `json:"token,omitempty"`
}

type GetCredentialsResp struct {
	Status  string           `json:"status,omitempty"`
	RetCode int32            `json:"retcode,omitempty"`
	Data    *CookiesAndToken `json:"data,omitempty"`
	Echo    string           `json:"echo,omitempty"`
}

type CookiesAndToken struct {
	Cookies string `json:"cookies,omitempty"`
	Token   int32  `json:"token,omitempty"`
}

type GetRecordResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    *IFile `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type IFile struct {
	File string `json:"file,omitempty"`
}

type GetImageResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    *IFile `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type CanSendImageResp struct {
	Status  string                 `json:"status,omitempty"`
	RetCode int32                  `json:"retcode,omitempty"`
	Data    *ICanSendImageOrRecord `json:"data,omitempty"`
	Echo    string                 `json:"echo,omitempty"`
}

type ICanSendImageOrRecord struct {
	Yes bool `json:"yes,omitempty"`
}

type CanSendRecordResp struct {
	Status  string                 `json:"status,omitempty"`
	RetCode int32                  `json:"retcode,omitempty"`
	Data    *ICanSendImageOrRecord `json:"data,omitempty"`
	Echo    string                 `json:"echo,omitempty"`
}

type GetStatusResp struct {
	Status  string      `json:"status,omitempty"`
	RetCode int32       `json:"retcode,omitempty"`
	Data    *IGetStatus `json:"data,omitempty"`
	Echo    string      `json:"echo,omitempty"`
}

type IGetStatus struct {
	Online bool `json:"online,omitempty"`
	Good   bool `json:"good,omitempty"`
}

type GetVersioInfoResp struct {
	Status  string          `json:"status,omitempty"`
	RetCode int32           `json:"retcode,omitempty"`
	Data    *IGetVersioInfo `json:"data,omitempty"`
	Echo    string          `json:"echo,omitempty"`
}

type IGetVersioInfo struct {
	AppName         string `json:"app_name,omitempty"`
	AppVersion      string `json:"app_version,omitempty"`
	ProtocolVersion string `json:"protocol_version,omitempty"` // OneBot 标准版本，如 v11
}

type SetRestartResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    any    `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type CleanCacheResp struct {
	Status  string `json:"status,omitempty"`
	RetCode int32  `json:"retcode,omitempty"`
	Data    any    `json:"data,omitempty"`
	Echo    string `json:"echo,omitempty"`
}
