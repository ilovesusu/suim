package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// UserInfo 用户信息
type UserInfo struct {
	BaseModel
	Number            string  `gorm:"index;not null;size:50;comment:用户号码"`
	Phone             *string `gorm:"index;not null;size:20;comment:电话号码"`
	Password          *string `gorm:"not null;size:255;comment:密码"`
	Name              *string `gorm:"index;size:50;comment:姓名"`
	IdCard            *string `gorm:"size:20;comment:身份证号"`
	Nickname          *string `gorm:"index;not null;size:100;comment:昵称"`
	Sex               *int32  `gorm:"index;not null;size:4;comment:性别(1-保密,2-男,3-女)"`
	AvatarUrl         *string `gorm:"size:255;comment:头像地址链接"`
	Introduce         *string `gorm:"size:255;comment:个人介绍"`
	SnapCall          *bool   `gorm:"not null;comment:是否允许临时会话"`
	AddFriendType     *int32  `gorm:"not null;comment:添加好友方式(1-直接通过,2-需要验证,3-回答问题通过验证,4-拒绝加好友)"`
	FriendPassProblem *string `gorm:"size:255;comment:问题通过好友请求问题"`
	FriendPassAnswer  *string `gorm:"size:255;comment:问题通过好友答案"`
}

// UserDevice 用户设备
type UserDevice struct {
	BaseModel
	Uid           int64   `gorm:"index;not null;comment:用户id"`
	DeviceType    *int32  `gorm:"not null;comment:用户设备类型(1-Android,2-IOS,3-Windows,4-MacOS,5-Linux,6-Web"`
	DeviceBrand   *string `gorm:"size:50;comment:设备品牌"`
	DeviceModel   *string `gorm:"size:50;comment:设备型号"`
	DeviceVersion *string `gorm:"size:50;comment:设备系统版本"`
	SdkVersion    *string `gorm:"size:30;comment:软件版本"`
	ConnStatus    *bool   `gorm:"index;not null;comment:设备是否在线"`
	ServerAddr    *string `gorm:"index;not null;comment:连接服务器地址"`
}

// UserFriend 用户好友
type UserFriend struct {
	BaseModel
	Uid          int64   `gorm:"index;not null;comment:用户id"`
	Fid          int64   `gorm:"index;not null;comment:好友id"`
	FriendStatus *int32  `gorm:"index;not null;comment:好友状态(1-申请,2-同意,3-拒绝,4-拉黑,5-删除)"`
	FriendRemark *string `gorm:"comment:好友备注"`
}

// UserGroup 用户群组
type UserGroup struct {
	BaseModel
	Uid          int64  `gorm:"index;not null;comment:用户id"`
	Gid          int64  `gorm:"index;not null;comment:群组id"`
	MemberType   *int32 `gorm:"index;not null;comment:成员类型(1-群主,2-管理员,3-普通成员)"`
	MemberStatus *int32 `gorm:"index;not null;comment:成员状态(1-正常,2-禁言,3-移除)"`
}

// UserChannel 用户频道
type UserChannel struct {
	BaseModel
	Uid             int64  `gorm:"index;not null;comment:用户id"`
	Cid             int64  `gorm:"index;not null;comment:频道id"`
	AttentionStatus *bool  `gorm:"index;not null;comment:关注状态"`
	MsgPushStatus   *int32 `gorm:"index;not null;comment:消息推送状态(1-正常接收消息推送,2-静默接收消息推送,3-不接收消息推送)"`
}

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

type UserRepo interface {
	CreateUser(context.Context, *UserInfo) error
	UpdateUser(context.Context, *UserInfo) error
	DeleteUser(context.Context, *UserInfo) error
	ListUser(context.Context, *UserFriend) ([]FriendList, error)
	InfoUser(context.Context, *UserInfo) (*UserInfo, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (u *UserUsecase) UserCreate(ctx context.Context, user *UserInfo) error {
	return u.repo.CreateUser(ctx, user)
}

func (u *UserUsecase) UserUpdate(ctx context.Context, user *UserInfo) error {
	return u.repo.UpdateUser(ctx, user)
}

func (u *UserUsecase) UserDelete(ctx context.Context, user *UserInfo) error {
	return u.repo.DeleteUser(ctx, user)
}

func (u *UserUsecase) UserList(ctx context.Context, user *UserFriend) ([]FriendList, error) {
	return u.repo.ListUser(ctx, user)
}

func (u *UserUsecase) UserInfo(ctx context.Context, user *UserInfo) (*UserInfo, error) {
	return u.repo.InfoUser(ctx, user)
}
