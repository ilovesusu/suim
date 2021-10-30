package biz

// GroupInfo 群组信息
type GroupInfo struct {
	BaseModel
	Number      string  `gorm:"not null;index;size:50;comment:群组号"`
	Name        *string `gorm:"not null;index;size:50;comment:群组名称"`
	AvatarUrl   *string `gorm:"not null;size:255;comment:群组头像链接"`
	Introduce   *string `gorm:"size:255;comment:群组简介"`
	MemberNum   int32   `gorm:"not null;comment:群组成员人数"`
	GroupStatus *int32  `gorm:"index;not null;comment:群组状态(1-正常,2-解散)"`
}

// TagInfo 标签信息
type TagInfo struct {
	BaseModel
	Name     *string `gorm:"index;not null;size:50;comment:标签名称"`
	OtherUse *bool   `gorm:"index;not null;comment:是否可被其他人使用"`
}

// GroupTag 群组标签
type GroupTag struct {
	BaseModel
	Gid int64 `gorm:"index;not null;comment:群组id"`
	Tid int64 `gorm:"index;not null;comment:标签id"`
}
