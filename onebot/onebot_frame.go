package onebot

type Frame struct {
	BotId       int64          `json:"bot_id,omitempty"`
	Echo        string         `json:"echo,omitempty"`
	Ok          bool           `json:"ok,omitempty"`
	Time        int64          `json:"time,omitempty"`
	SelfId      int64          `json:"self_id,omitempty"`
	PostType    PostType       `json:"post_type,omitempty"`
	MessageType string         `json:"message_type,omitempty"`
	Status      any            `json:"status,omitempty"`
	Code        int32          `json:"retcode,omitempty"`
	Data        map[string]any `json:"data,omitempty"`
	*API
}
