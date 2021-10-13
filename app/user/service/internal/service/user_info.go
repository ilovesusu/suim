package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/ilovesusu/suim/api/user/service/v1"
	"github.com/ilovesusu/suim/app/user/service/internal/biz"
	"github.com/ilovesusu/suim/pkg"
)

// UserService is a greeter service.
type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewUserService 创建一个用户服务
func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (us *UserService) Hello(ctx context.Context, req *v1.HelloReq) (*v1.HelloRsp, error) {
	return &v1.HelloRsp{Hello: "你好!" + req.Name}, nil
}

// UserCreate 添加用户
func (us *UserService) UserCreate(ctx context.Context, req *v1.UserCreateReq) (*v1.UserCreateRsp, error) {
	user := &biz.UserInfo{
		Phone:             &req.Phone,
		Password:          &req.Password,
		Name:              &req.Name,
		IdCard:            pkg.GetFromStringValue(req.IdCard),
		Nickname:          &req.Nickname,
		Sex:               &req.Sex,
		AvatarUrl:         pkg.GetFromStringValue(req.AvatarUrl),
		Introduce:         pkg.GetFromStringValue(req.Introduce),
		SnapCall:          &req.SnapCall,
		AddFriendType:     &req.AddFriendType,
		FriendPassProblem: pkg.GetFromStringValue(req.FriendPassProblem),
		FriendPassAnswer:  pkg.GetFromStringValue(req.FriendPassAnswer),
	}
	if err := us.uc.UserCreate(ctx, user); err != nil {
		return nil, err
	}
	return &v1.UserCreateRsp{}, nil
}

// UserUpdate 修改用户信息
func (us *UserService) UserUpdate(ctx context.Context, req *v1.UserUpdateReq) (*v1.UserUpdateRsp, error) {
	user := biz.UserInfo{
		BaseModel:         biz.BaseModel{Id: req.Id},
		Number:            "",
		Phone:             pkg.GetFromStringValue(req.Phone),
		Password:          pkg.GetFromStringValue(req.Password),
		Name:              pkg.GetFromStringValue(req.Name),
		IdCard:            pkg.GetFromStringValue(req.IdCard),
		Nickname:          pkg.GetFromStringValue(req.Nickname),
		Sex:               pkg.GetFromInt32Value(req.Sex),
		AvatarUrl:         pkg.GetFromStringValue(req.AvatarUrl),
		Introduce:         pkg.GetFromStringValue(req.Introduce),
		SnapCall:          pkg.GetFromBoolValue(req.SnapCall),
		AddFriendType:     pkg.GetFromInt32Value(req.AddFriendType),
		FriendPassProblem: pkg.GetFromStringValue(req.FriendPassProblem),
		FriendPassAnswer:  pkg.GetFromStringValue(req.FriendPassAnswer),
	}
	if err := us.uc.UserUpdate(ctx, &user); err != nil {
		return nil, err
	}
	return &v1.UserUpdateRsp{}, nil
}

// UserDelete 注销用户
func (us *UserService) UserDelete(ctx context.Context, req *v1.UserDeleteReq) (*v1.UserDeleteRsp, error) {
	if err := us.uc.UserDelete(ctx, &biz.UserInfo{BaseModel: biz.BaseModel{Id: req.Id}}); err != nil {
		return nil, err
	}
	return &v1.UserDeleteRsp{}, nil
}

// UserInfo 用户信息
func (us *UserService) UserInfo(ctx context.Context, req *v1.UserInfoReq) (*v1.UserInfoRsp, error) {
	info, err := us.uc.UserInfo(ctx, &biz.UserInfo{BaseModel: biz.BaseModel{Id: req.Id}})
	if err != nil {
		return nil, err
	}
	var rsp = &v1.UserInfoRsp{
		Id:                info.Id,
		Name:              *info.Name,
		Phone:             *info.Phone,
		Nickname:          *info.Nickname,
		Sex:               *info.Sex,
		AvatarUrl:         pkg.CreateStringValuePtr(info.AvatarUrl),
		Introduce:         pkg.CreateStringValuePtr(info.Introduce),
		SnapCall:          *info.SnapCall,
		AddFriendType:     *info.AddFriendType,
		FriendPassProblem: pkg.CreateStringValuePtr(info.FriendPassProblem),
		FriendPassAnswer:  pkg.CreateStringValuePtr(info.FriendPassAnswer),
	}
	return rsp, nil
}

// UserList 用户好友列表
func (us *UserService) UserList(ctx context.Context, req *v1.UserListReq) (*v1.UserListRsp, error) {
	list, err := us.uc.UserList(ctx, &biz.UserFriend{Uid: req.Uid, FriendStatus: &req.FriendStatus})
	if err != nil {
		return nil, err
	}
	var rsp v1.UserListRsp
	rsp.Total = int32(len(list))
	for _, v := range list {
		rsp.List = append(rsp.List, &v1.UserListRsp_List{
			UserId:       v.Id,
			Nickname:     v.Nickname,
			Sex:          v.Sex,
			AvatarUrl:    pkg.CreateStringValuePtr(v.AvatarUrl),
			Introduce:    pkg.CreateStringValuePtr(v.Introduce),
			FriendRemake: pkg.CreateStringValuePtr(v.FriendRemark),
		})
	}
	return &rsp, nil
}
