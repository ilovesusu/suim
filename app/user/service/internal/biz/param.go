package biz

// FriendList 好友列表
type FriendList struct {
	Id           int64   `gorm:"index;not null;comment:用户id"`
	Nickname     string  `gorm:"index;not null;size:100;comment:昵称"`
	Sex          int32   `gorm:"index;not null;size:4;comment:性别(1-保密,2-男,3-女)"`
	AvatarUrl    *string `gorm:"not null;size:255;comment:头像地址链接"`
	Introduce    *string `gorm:"size:255;comment:个人介绍"`
	FriendRemark *string `gorm:"comment:好友备注"`
}
