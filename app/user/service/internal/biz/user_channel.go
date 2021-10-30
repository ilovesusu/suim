package biz

// UserChannel 用户频道
type UserChannel struct {
	BaseModel
	Uid             int64  `gorm:"index;not null;comment:用户id"`
	Cid             int64  `gorm:"index;not null;comment:频道id"`
	AttentionStatus *bool  `gorm:"index;not null;comment:关注状态"`
	MsgPushStatus   *int32 `gorm:"index;not null;comment:消息推送状态(1-正常接收消息推送,2-静默接收消息推送,3-不接收消息推送)"`
}
