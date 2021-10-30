package biz

// UserGroup 用户群组
type UserGroup struct {
	BaseModel
	Uid          int64  `gorm:"index;not null;comment:用户id"`
	Gid          int64  `gorm:"index;not null;comment:群组id"`
	MemberType   *int32 `gorm:"index;not null;comment:成员类型(1-群主,2-管理员,3-普通成员)"`
	MemberStatus *int32 `gorm:"index;not null;comment:成员状态(1-正常,2-禁言,3-移除)"`
}
