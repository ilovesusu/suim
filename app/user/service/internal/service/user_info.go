package service

import (
	"context"
	"github.com/ilovesusu/suim/api/user/service/v1/user"
	"github.com/ilovesusu/suim/app/user/service/internal/biz"
	"github.com/ilovesusu/suim/pkg"
)

// Hello hello测试
func (us *UserService) Hello(ctx context.Context, req *user.HelloReq) (*user.HelloRsp, error) {
	return &user.HelloRsp{Hello: "你好!" + req.Name}, nil
}

// CreateUser 创建用户
func (us *UserService) CreateUser(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserRsp, error) {
	info := &biz.UserInfo{
		Phone:             req.Phone,
		Password:          req.Password,
		Name:              pkg.GetFromStringValue(req.Name),
		IdCard:            pkg.GetFromStringValue(req.IdCard),
		Nickname:          req.Nickname,
		Sex:               req.Sex,
		AvatarUrl:         pkg.GetFromStringValue(req.AvatarUrl),
		PersonalSign:      pkg.GetFromStringValue(req.PersonalSign),
		Introduce:         pkg.GetFromStringValue(req.Introduce),
		SnapCall:          req.SnapCall,
		FriendPassType:    req.AddFriendType,
		FriendPassProblem: pkg.GetFromStringValue(req.FriendPassProblem),
		FriendPassAnswer:  pkg.GetFromStringValue(req.FriendPassAnswer),
	}
	if err := us.uc.CreateUser(ctx, info); err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdateIdCard 修改身份信息
func (us *UserService) UpdateIdCard(ctx context.Context, req *user.UpdateIdCardReq) (*user.UpdateIdCardRsp, error) {
	if err := us.uc.UpdateAccount(ctx, &biz.UpdateIdCardReq{
		Id:     req.Id,
		Name:   req.Name,
		IdCard: req.IdCard,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdatePhone 修改电话号码
func (us *UserService) UpdatePhone(ctx context.Context, req *user.UpdatePhoneReq) (*user.UpdatePhoneRsp, error) {
	if err := us.uc.UpdatePhone(ctx, &biz.UpdatePhoneReq{
		Id:    req.Id,
		Phone: req.Phone,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdatePassword 修改密码
func (us *UserService) UpdatePassword(ctx context.Context, req *user.UpdatePasswordReq) (*user.UpdatePasswordRsp, error) {
	if err := us.uc.UpdatePassword(ctx, &biz.UpdatePasswordReq{
		Id:          req.Id,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

// ForgetPassword 忘记密码
func (us *UserService) ForgetPassword(ctx context.Context, req *user.ForgetPasswordReq) (*user.ForgetPasswordRsp, error) {
	if err := us.uc.ForgetPassword(ctx, &biz.ForgetPasswordReq{
		Phone:    req.Phone,
		Password: req.Password,
		Code:     req.Code,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdateNickname 修改昵称
func (us *UserService) UpdateNickname(ctx context.Context, req *user.UpdateNicknameReq) (*user.UpdateNicknameRsp, error) {
	if err := us.uc.UpdateNickname(ctx, &biz.UpdateNicknameReq{
		Id:       req.Id,
		Nickname: req.Nickname,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdateSex 修改性别
func (us *UserService) UpdateSex(ctx context.Context, req *user.UpdateSexReq) (*user.UpdateSexRsp, error) {
	if err := us.uc.UpdateSex(ctx, &biz.UpdateSexReq{
		Id:  req.Id,
		Sex: req.Sex,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdateAvatarUrl 修改头像
func (us *UserService) UpdateAvatarUrl(ctx context.Context, req *user.UpdateAvatarUrlReq) (*user.UpdateAvatarUrlRsp, error) {
	if err := us.uc.UpdateAvatarUrl(ctx, &biz.UpdateAvatarUrlReq{
		Id:        req.Id,
		AvatarUrl: req.AvatarUrl,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdatePersonalSign 修改个性签名
func (us *UserService) UpdatePersonalSign(ctx context.Context, req *user.UpdatePersonalSignReq) (*user.UpdatePersonalSignRsp, error) {
	if err := us.uc.UpdatePersonalSign(ctx, &biz.UpdatePersonalSignReq{
		Id:           req.Id,
		PersonalSign: req.PersonalSign,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdateIntroduce 修改个人介绍
func (us *UserService) UpdateIntroduce(ctx context.Context, req *user.UpdateIntroduceReq) (*user.UpdateIntroduceRsp, error) {
	if err := us.uc.UpdateIntroduce(ctx, &biz.UpdateIntroduceReq{
		Id:        req.Id,
		Introduce: req.Introduce,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdateSnapCall 修改是否允许临时会话
func (us *UserService) UpdateSnapCall(ctx context.Context, req *user.UpdateSnapCallReq) (*user.UpdateSnapCallRsp, error) {
	if err := us.uc.UpdateSnapCall(ctx, &biz.UpdateSnapCallReq{
		Id:       req.Id,
		SnapCall: req.SnapCall,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdateFriendPass 修改用户添加好友方式
func (us *UserService) UpdateFriendPass(ctx context.Context, req *user.UpdateFriendPassReq) (*user.UpdateFriendPassRsp, error) {
	if err := us.uc.UpdateFriendPass(ctx, &biz.UpdateFriendPassReq{
		Id:                req.Id,
		FriendPassType:    req.AddFriendType,
		FriendPassProblem: pkg.GetFromStringValue(req.FriendPassProblem),
		FriendPassAnswer:  pkg.GetFromStringValue(req.FriendPassAnswer),
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

// DeleteUser 删除帐号
func (us *UserService) DeleteUser(ctx context.Context, req *user.DeleteUserReq) (*user.DeleteUserRsp, error) {
	if err := us.uc.DeleteUser(ctx, &biz.DeleteUserReq{
		Id:       req.Id,
		Password: req.Password,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

// InfoUserBase 查询用户基本信息
func (us *UserService) InfoUserBase(ctx context.Context, req *user.InfoUserBaseReq) (*user.InfoUserBaseRsp, error) {
	info, err := us.uc.InfoUserBase(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &user.InfoUserBaseRsp{
		Number:       info.Number,
		Nickname:     info.Nickname,
		Sex:          info.Sex,
		AvatarUrl:    pkg.CreateStringValuePtr(info.AvatarUrl),
		PersonalSign: pkg.CreateStringValuePtr(info.PersonalSign),
		Introduce:    pkg.CreateStringValuePtr(info.Introduce),
	}, nil
}

// InfoAccount 查询用户身份信息
func (us *UserService) InfoAccount(ctx context.Context, req *user.InfoAccountReq) (*user.InfoAccountRsp, error) {
	info, err := us.uc.InfoAccount(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &user.InfoAccountRsp{
		Phone:  info.Phone,
		Name:   pkg.CreateStringValuePtr(info.Name),
		IdCard: pkg.CreateStringValuePtr(info.IdCard),
	}, nil
}

// InfoSnapCall 查询用户是否允许临时会话
func (us *UserService) InfoSnapCall(ctx context.Context, req *user.InfoSnapCallReq) (*user.InfoSnapCallRsp, error) {
	info, err := us.uc.InfoSnapCall(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &user.InfoSnapCallRsp{SnapCall: *info}, nil
}

// InfoFriendPass 查询用户添加好友方式
func (us *UserService) InfoFriendPass(ctx context.Context, req *user.InfoFriendPassReq) (*user.InfoFriendPassRsp, error) {
	info, err := us.uc.InfoFriendPass(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &user.InfoFriendPassRsp{
		FriendPassType:    info.FriendPassType,
		FriendPassProblem: pkg.CreateStringValuePtr(info.FriendPassProblem),
		FriendPassAnswer:  pkg.CreateStringValuePtr(info.FriendPassAnswer),
	}, nil
}
