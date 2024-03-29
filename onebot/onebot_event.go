package onebot

type PostType string
type MsgEventType string
type PrivateSubEventType string
type GroupSubEventType string
type SexType string
type NoticeType string
type MetaEventType string
type HonorType string
type ReqType string

const (
	Message   PostType = "message"
	Notice    PostType = "notice"
	Request   PostType = "request"
	MetaEvent PostType = "meta_event"

	LifeCycle MetaEventType = "lifecycle"
	HeartBeat MetaEventType = "heartbeat"

	Private MsgEventType = "private"
	Group   MsgEventType = "group"

	SubFriend PrivateSubEventType = "friend"
	SubGroup  PrivateSubEventType = "group"
	SubOther  PrivateSubEventType = "other"

	SubNormal    GroupSubEventType = "normal"
	SubAnonymous GroupSubEventType = "anonymous"
	SubNotice    GroupSubEventType = "notice"
	SetAdmin     GroupSubEventType = "set"
	RemoveAdmin  GroupSubEventType = "unset"
	Leave        GroupSubEventType = "leave"   // 主动退群
	Kick         GroupSubEventType = "kick"    // 成员被踢
	KickMe       GroupSubEventType = "kick_me" // 登录号被踢
	Approve      GroupSubEventType = "approve" // 管理员已同意入群
	Invite       GroupSubEventType = "invite"  // 管理员邀请入群
	Ban          GroupSubEventType = "ban"
	LiftBan      GroupSubEventType = "lift_ban"
	Poke         GroupSubEventType = "poke"
	LuckyKing    GroupSubEventType = "lucky_king"
	Honor        GroupSubEventType = "honor"

	Male    SexType = "male"
	Female  SexType = "female"
	Unknown SexType = "unknown"

	GroupUpload   NoticeType = "group_upload"
	GroupAdmin    NoticeType = "group_admin"
	GroupDecrease NoticeType = "group_decrease"
	GroupIncrease NoticeType = "group_increase"
	GroupBan      NoticeType = "group_ban"
	FriendAdd     NoticeType = "friend_add"
	GroupRecall   NoticeType = "group_recall"
	FriendRecall  NoticeType = "friend_recall"
	Notify        NoticeType = "notify"

	Talkative HonorType = "talkative"
	Performer HonorType = "performer"
	Emotion   HonorType = "emotion"

	FriendAddRequest        ReqType = "friend"
	GroupAddOrInviteRequest ReqType = "group"
)

type Event struct {
	Time     int64  `json:"time,omitempty"`
	SelfId   int64  `json:"self_id,omitempty"`
	PostType string `json:"post_type,omitempty"`
}

type PrivateSender struct {
	UserId   int64   `json:"user_id,omitempty"`
	Nickname string  `json:"nickname,omitempty"`
	Sex      SexType `json:"sex,omitempty"`
	Age      int32   `json:"age,omitempty"`
}

type PrivateMsgEvent struct {
	Event
	MessageType string         `json:"message_type,omitempty"`
	SubType     string         `json:"sub_type,omitempty"`
	MessageId   int64          `json:"message_id,omitempty"`
	UserId      int64          `json:"user_id,omitempty"`
	Message     []*IMessage    `json:"message,omitempty"`
	RawMessage  string         `json:"raw_message,omitempty"`
	Font        int32          `json:"font,omitempty"`
	Sender      *PrivateSender `json:"sender,omitempty"`
}

type PrivateQuickOperate struct {
	Reply      string `json:"reply,omitempty"`
	AutoEscape bool   `json:"auto_escape,omitempty"`
}

type Anonymous struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Flag string `json:"flag,omitempty"`
}

type GroupSender struct {
	UserId   int64  `json:"user_id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Card     string `json:"card,omitempty"`
	Sex      string `json:"sex,omitempty"`
	Age      int32  `json:"age,omitempty"`
	Area     string `json:"area,omitempty"`
	Level    string `json:"level,omitempty"`
	Role     string `json:"role,omitempty"`
	Title    string `json:"title,omitempty"`
}
type GroupMsgEvent struct {
	Event
	MessageType string       `json:"message_type,omitempty"`
	SubType     string       `json:"sub_type,omitempty"`
	MessageId   int64        `json:"message_id,omitempty"`
	GroupId     int64        `json:"group_id,omitempty"`
	UserId      int64        `json:"user_id,omitempty"`
	Anonymous   *Anonymous   `json:"anonymous,omitempty"`
	Message     []*IMessage  `json:"message,omitempty"`
	RawMessage  string       `json:"raw_message,omitempty"`
	Font        int32        `json:"font,omitempty"`
	Sender      *GroupSender `json:"sender,omitempty"`
}
type GroupQuickOperate struct {
	Reply       string `json:"reply,omitempty"`
	AutoEscape  bool   `json:"auto_escape,omitempty"`
	AtSender    bool   `json:"at_sender,omitempty"`
	Delete      bool   `json:"delete,omitempty"`
	Kick        bool   `json:"kick,omitempty"`
	Ban         bool   `json:"ban,omitempty"`
	BanDuration int64  `json:"ban_duration,omitempty"`
}

type File struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Size int64  `json:"size,omitempty"`
}

type IMessage struct {
	Type string         `json:"type,omitempty"`
	Data map[string]any `json:"data,omitempty"`
}

type ForwardMsg struct {
	Name    string `json:"name,omitempty"`
	Uin     string `json:"uin,omitempty"`
	Content *IMessage   `json:"content,omitempty"`
}

type LifeTime struct {
	Event
	MetaEventType string `json:"meta_event_type"`
	SubType       string `json:"sub_type"`
}

type BotHeartBeat struct {
	Event
	MetaEventType string           `json:"meta_event_type"`
	Status        *HeartBeatStatus `json:"status,omitempty"`
}

type HeartBeatStatus struct {
	AppInitialized bool `json:"app_initialized,omitempty"`
	AppEnabled     bool `json:"app_enabled,omitempty"`
	AppGood        bool `json:"app_good,omitempty"`
	Online         bool `json:"online,omitempty"`
	Good           bool `json:"good,omitempty"`
}

type GroupUploadNoticeEvent struct {
	Event
	NoticeType NoticeType  `json:"notice_type,omitempty"`
	GroupId    int64       `json:"group_id,omitempty"`
	UserId     int64       `json:"user_id,omitempty"`
	File       *UploadFile `json:"file,omitempty"`
}

type UploadFile struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Size  int64  `json:"size,omitempty"`
	BusId int64  `json:"busid,omitempty"`
}

type GroupAdminChangeNoticeEvent struct {
	Event
	NoticeType NoticeType `json:"notice_type,omitempty"`
	SubType    string     `json:"sub_type,omitempty"`
	GroupId    int64      `json:"group_id,omitempty"`
	UserId     int64      `json:"user_id,omitempty"`
}

type GroupMemberDecreaseNoticeEvent struct {
	Event
	NoticeType NoticeType `json:"notice_type,omitempty"`
	SubType    string     `json:"sub_type,omitempty"`
	GroupId    int64      `json:"group_id,omitempty"`
	OperatorId int64      `json:"operator_id,omitempty"`
	UserId     int64      `json:"user_id,omitempty"`
}

type GroupMemberIncreaseNoticeEvent struct {
	Event
	NoticeType NoticeType `json:"notice_type,omitempty"`
	SubType    string     `json:"sub_type,omitempty"`
	GroupId    int64      `json:"group_id,omitempty"`
	OperatorId int64      `json:"operator_id,omitempty"`
	UserId     int64      `json:"user_id,omitempty"`
}

type GroupBanNoticeEvent struct {
	Event
	NoticeType NoticeType `json:"notice_type,omitempty"`
	SubType    string     `json:"sub_type,omitempty"`
	GroupId    int64      `json:"group_id,omitempty"`
	OperatorId int64      `json:"operator_id,omitempty"`
	UserId     int64      `json:"user_id,omitempty"`
	Duration   int64      `json:"duration,omitempty"`
}

type FriendAddNoticeEvent struct {
	Event
	NoticeType NoticeType `json:"notice_type,omitempty"`
	UserId     int64      `json:"user_id,omitempty"`
}

type GroupMsgRecallNoticeEvent struct {
	Event
	NoticeType NoticeType `json:"notice_type,omitempty"`
	GroupId    int64      `json:"group_id,omitempty"`
	OperatorId int64      `json:"operator_id,omitempty"`
	UserId     int64      `json:"user_id,omitempty"`
	MessageId  int64      `json:"message_id,omitempty"`
}

type FriendMsgRecallNoticeEvent struct {
	Event
	NoticeType NoticeType `json:"notice_type,omitempty"`
	UserId     int64      `json:"user_id,omitempty"`
	MessageId  int64      `json:"message_id,omitempty"`
}

type GroupPokeNoticeEvent struct {
	Event
	NoticeType NoticeType `json:"notice_type,omitempty"`
	SubType    string     `json:"sub_type,omitempty"`
	GroupId    int64      `json:"group_id,omitempty"`
	UserId     int64      `json:"user_id,omitempty"`
	TargetId   int64      `json:"target_id,omitempty"`
}

type GroupLuckyKingNoticeEvent struct {
	Event
	NoticeType NoticeType `json:"notice_type,omitempty"`
	SubType    string     `json:"sub_type,omitempty"`
	GroupId    int64      `json:"group_id,omitempty"`
	UserId     int64      `json:"user_id,omitempty"`
	TargetId   int64      `json:"target_id,omitempty"`
}

type GroupMemberHonorChangeNoticeEvent struct {
	Event
	NoticeType NoticeType `json:"notice_type,omitempty"`
	SubType    string     `json:"sub_type,omitempty"`
	GroupId    int64      `json:"group_id,omitempty"`
	HonorType  HonorType  `json:"honor_type,omitempty"`
	UserId     int64      `json:"user_id,omitempty"`
}

type FriendAddRequestEvent struct {
	Event
	RequestType string `json:"request_type,omitempty"`
	UserId      int64  `json:"user_id,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Flag        string `json:"flag,omitempty"`
}

type GroupAddOrInviteRequestEvent struct {
	Event
	RequestType string `json:"request_type,omitempty"`
	SubType     string `json:"sub_type,omitempty"`
	GroupId     int64  `json:"group_id,omitempty"`
	UserId      int64  `json:"user_id,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Flag        string `json:"flag,omitempty"`
}
