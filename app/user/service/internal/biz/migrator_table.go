package biz

// MigratorTable 此数组专门存储结构体用来迁移表使用
var MigratorTable = []interface{}{
	&UserInfo{},    //用户信息表
	&UserDevice{},  //用户设备表
	&UserFriend{},  //用户好友表
	&UserGroup{},   //用户群组表
	&UserChannel{}, //用户频道表
	&GroupInfo{},   //群组信息表
	&ChannelInfo{}, //频道信息表
	&TagInfo{},     //标签信息表
	&GroupTag{},    //群组标签表
}
