package biz

// CreateUserReq 创建用户请求参数
type CreateUserReq struct {
	Phone     string //手机号
	Password  string //密码
	Nickname  string //昵称
	Sex       int32  //性别
	AvatarUrl string //头像地址
}

// UpdateIdCardReq 修改身份信息请求参数
type UpdateIdCardReq struct {
	Id     int64  //用户id
	Name   string //姓名
	IdCard string //身份证号码
}

// UpdatePhoneReq 修改用户电话号码请求参数
type UpdatePhoneReq struct {
	Id    int64  //用户id
	Phone string //手机号
}

// UpdatePasswordReq 修改用户密码请求参数
type UpdatePasswordReq struct {
	Id          int64  //用户id
	OldPassword string //旧密码
	NewPassword string //新密码
}

// ForgetPasswordReq 忘记密码请求参数
type ForgetPasswordReq struct {
	Phone    string //手机号
	Password string //密码
	Code     string //验证码
}

// UpdateNicknameReq 修改用户昵称请求参数
type UpdateNicknameReq struct {
	Id       int64  //用户id
	Nickname string //昵称
}

// UpdateSexReq 修改用户性别请求参数
type UpdateSexReq struct {
	Id  int64 //用户id
	Sex int32 //性别
}

// UpdateAvatarUrlReq 修改用户头像请求参数
type UpdateAvatarUrlReq struct {
	Id        int64  //用户id
	AvatarUrl string //头像地址
}

// UpdatePersonalSignReq 修改用户个性签名请求参数
type UpdatePersonalSignReq struct {
	Id           int64  //用户id
	PersonalSign string //个性签名
}

// UpdateIntroduceReq 修改用户自我介绍请求参数
type UpdateIntroduceReq struct {
	Id        int64  //用户id
	Introduce string //自我介绍
}

// UpdateSnapCallReq 修改用户是否允许临时会话请求参数
type UpdateSnapCallReq struct {
	Id       int64 //用户id
	SnapCall bool  //是否允许临时会话
}

// UpdateFriendPassReq 修改用户添加好友方式
type UpdateFriendPassReq struct {
	Id                int64  //用户id
	FriendPassType    int32  //添加好友方式(1-直接通过,2-需要验证,3-回答问题通过验证,4-拒绝加好友)
	FriendPassProblem string //问题通过好友请求问题
	FriendPassAnswer  string //问题通过好友答案
}

// DeleteUserReq 删除用户帐号
type DeleteUserReq struct {
	Id       int64  //用户id
	Password string //密码
}

// InfoUserBaseRsp 查询用户基本信息响应参数
type InfoUserBaseRsp struct {
	Number       string //用户号码
	Nickname     string //昵称
	Sex          int32  //性别
	AvatarUrl    string //头像
	PersonalSign string //个性签名
	Introduce    string //个人介绍
}

// InfoAccountRsp 查询用户身份信息响应参数
type InfoAccountRsp struct {
	Phone  string //手机号
	Name   string //姓名
	IdCard string //身份证号码
}

// InfoFriendPassRsp 查询用户添加好友方式响应参数
type InfoFriendPassRsp struct {
	FriendPassType    int32  //添加好友方式(1-直接通过,2-需要验证,3-回答问题通过验证,4-拒绝加好友)
	FriendPassProblem string //问题通过好友请求问题
	FriendPassAnswer  string //问题通过好友答案
}
