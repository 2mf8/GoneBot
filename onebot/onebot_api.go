package onebot

type ActionType string

const (
	SendPrivateMsg       ActionType = "send_private_msg"
	SendGroupMsg         ActionType = "send_group_msg"
	SendMsg              ActionType = "send_msg"
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
	SetGoupAddRequest    ActionType = "set_group_add_request"
	GetLoginInfo         ActionType = "get_login_info"
	GetStrangerInfo      ActionType = "get_stranger_info"
	GetFriendList        ActionType = "get_friend_list"
	GetFroupInfo         ActionType = "get_group_info"
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
	MessageId        int32       `json:"message_id,omitempty"`
	Id               string      `json:"id,omitempty"`
	RejectAddRequest bool        `json:"reject_add_request,omitempty"`
	Duration         int32       `json:"duration,omitempty"`
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
}

// {"status":"ok","retcode":0,"data":{"message_id":528925193},"echo":"1709455334"}
type SendMsgResponse struct {
	Status  string               `json:"status,omitempty"`
	RetCode int32                `json:"retcode,omitempty"`
	Data    *SendMsgResponseData `json:"data,omitempty"`
}
type SendMsgResponseData struct {
	MessageId int32  `json:"message_id,omitempty"`
	Echo      string `json:"echo,omitempty"`
}
