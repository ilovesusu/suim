package biz

// ListFriendReq 列表请求参数
type ListFriendReq struct {
	UserId       int64
	FriendStatus int32
}

// ListUserFriendRsp 列表响应参数
type ListUserFriendRsp struct {
	Id           int64   //好友id
	Nickname     string  //好友昵称
	Sex          int32   //好友性别
	AvatarUrl    *string //好友头像
	PersonalSign *string //好友个性签名
	Remark       *string //好友备注
}

// UpdateFriendRemarkReq 修改好友备注请求参数
type UpdateFriendRemarkReq struct {
	UserId       int64  //用户id
	FriendId     int64  //好友id
	FriendRemark string //好友备注
}

// UpdateFriendStatusReq 修改好友状态请求参数
type UpdateFriendStatusReq struct {
	UserId        int64   //用户id
	FriendId      int64   //好友id
	FriendStatus  int32   //好友状态
	VerifyMessage *string //验证消息
}
