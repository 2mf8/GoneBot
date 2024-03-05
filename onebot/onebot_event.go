package onebot

type PostType string
type MsgEventType string
type PrivateSubEventType string
type GroupSubEventType string
type SexType string
type NoticeType string

const (
	Message   PostType = "message"
	Notice    PostType = "notice"
	Request   PostType = "request"
	MetaEvent PostType = "meta_event"

	Private MsgEventType = "private"
	Group   MsgEventType = "group"

	SubFriend PrivateSubEventType = "friend"
	SubGroup  PrivateSubEventType = "group"
	SubOther  PrivateSubEventType = "other"

	SubNormal    GroupSubEventType = "normal"
	SubAnonymous GroupSubEventType = "anonymous"
	SubNotice    GroupSubEventType = "notice"

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
	LuckyKing     NoticeType = "lucky_king"
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
	MessageId   int32          `json:"message_id,omitempty"`
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
	MessageId   int32        `json:"message_id,omitempty"`
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
	BanDuration int32  `json:"ban_duration,omitempty"`
}

type File struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Size int64  `json:"size,omitempty"`
}

type GroupUploadEvent struct {
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

type IMessage struct {
	Type string            `json:"type,omitempty"`
	Data map[string]string `json:"data,omitempty"`
}
