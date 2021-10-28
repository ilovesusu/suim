package biz

// ChannelInfo 频道信息
type ChannelInfo struct {
	BaseModel
	Number        string  `gorm:"index;size:50;not null;comment:频道号"`
	Name          *string `gorm:"index;size:50;not null;comment:频道名称"`
	AvatarUrl     *string `gorm:"not null;size:255;comment:频道头像链接"`
	Introduce     *string `gorm:"size:255;comment:频道简介"`
	ChannelType   *int32  `gorm:"index;not null;comment:频道类型(1-科技,2-知识,3-美食,4-技术,5-网络...)"`
	ChannelStatus *int32  `gorm:"index;not null;comment:频道状态(1-正常,2-关闭...)"`
	AttentionNum  int32   `gorm:"index;not null;comment:关注数量"`
	MsgNum        int32   `gorm:"not null;comment:文章数量"`
}
