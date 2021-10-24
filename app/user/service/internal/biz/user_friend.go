package biz

// UserFriend 用户好友
type UserFriend struct {
	BaseModel
	Uid          int64   `gorm:"index;not null;comment:用户id"`
	Fid          int64   `gorm:"index;not null;comment:好友id"`
	FriendStatus *int32  `gorm:"index;not null;comment:好友状态(1-申请,2-同意,3-拒绝,4-拉黑,5-删除)"`
	FriendRemark *string `gorm:"comment:好友备注"`
}
