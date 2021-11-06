package service

import (
	"context"
	"github.com/ilovesusu/suim/api/user/service/v1/friend"
	"github.com/ilovesusu/suim/app/user/service/internal/biz"
)

// CreateFriend 创建好友
func (f *FriendService) CreateFriend(ctx context.Context, req *friend.CreateFriendReq) (*friend.CreateFriendRsp, error) {
	if err := f.uc.CreateFriend(ctx, &biz.CreateFriendReq{
		UserId:        req.UserId,
		FriendId:      req.FriendId,
		VerifyMessage: req.VerifyMessage,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdateFriendStatus 修改好友状态
func (f *FriendService) UpdateFriendStatus(ctx context.Context, req *friend.UpdateFriendStatusReq) (*friend.UpdateFriendStatusRsp, error) {
	if err := f.uc.UpdateFriendStatus(ctx, &biz.UpdateFriendStatusReq{
		UserId:       req.UserId,
		FriendId:     req.FriendId,
		FriendStatus: req.FriendStatus,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdateFriendRemark 修改好友备注
func (f *FriendService) UpdateFriendRemark(ctx context.Context, req *friend.UpdateFriendRemarkReq) (*friend.UpdateFriendRemarkRsp, error) {
	if err := f.uc.UpdateFriendRemark(ctx, &biz.UpdateFriendRemarkReq{
		UserId:       req.UserId,
		FriendId:     req.FriendId,
		FriendRemark: req.FriendRemark,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

// ListUserFriend 好友列表
func (f *FriendService) ListUserFriend(ctx context.Context, req *friend.ListUserFriendReq) (*friend.ListUserFriendRsp, error) {
	list, err := f.uc.ListUserFriend(ctx, &biz.ListFriendReq{
		UserId:       req.UserId,
		FriendStatus: req.FriendStatus,
	})
	if err != nil {
		return nil, err
	}
	rsp := friend.ListUserFriendRsp{}
	rspList := make([]*friend.ListUserFriendRsp_List, 0)
	for _, v := range list {
		rspList = append(rspList, &friend.ListUserFriendRsp_List{
			Id:           v.Id,
			Nickname:     v.Nickname,
			Sex:          v.Sex,
			AvatarUrl:    v.AvatarUrl,
			PersonalSign: v.PersonalSign,
			Remark:       v.FriendRemark,
		})
	}
	rsp.List = rspList
	return &rsp, nil
}
