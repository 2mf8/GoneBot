package onebot

type Frame struct {
	BotId         int64          `json:"bot_id,omitempty"`
	Echo          string         `json:"echo,omitempty"`
	Ok            bool           `json:"ok,omitempty"`
	Time          int64          `json:"time,omitempty"`
	SelfId        int64          `json:"self_id,omitempty"`
	PostType      PostType       `json:"post_type,omitempty"`
	SubType       string         `json:"sub_type,omitempty"`
	MessageType   string         `json:"message_type,omitempty"`
	NoticeType    string         `json:"notice_type,omitempty"`
	RequestType   string         `json:"request_type,omitempty"`
	MetaEventType string         `json:"meta_event_type"`
	Status        any            `json:"status,omitempty"`
	Code          int32          `json:"retcode,omitempty"`
	Data          map[string]any `json:"data,omitempty"`
	*API
}
