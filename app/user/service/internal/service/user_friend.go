package service

import (
	"context"
	"github.com/ilovesusu/suim/api/user/service/v1/friend"
	"github.com/ilovesusu/suim/app/user/service/internal/biz"
	"github.com/ilovesusu/suim/pkg"
)

// UpdateFriendStatus 修改好友状态
func (f *FriendService) UpdateFriendStatus(ctx context.Context, req *friend.UpdateFriendStatusReq) (*friend.UpdateFriendStatusRsp, error) {
	if err := f.uc.UpdateFriendStatus(ctx, &biz.UpdateFriendStatusReq{
		UserId:        req.UserId,
		FriendId:      req.FriendId,
		FriendStatus:  req.FriendStatus,
		VerifyMessage: pkg.GetFromStringValue(req.VerifyMessage),
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
			AvatarUrl:    pkg.CreateStringValuePtr(v.AvatarUrl),
			PersonalSign: pkg.CreateStringValuePtr(v.PersonalSign),
			Remark:       pkg.CreateStringValuePtr(v.Remark),
		})
	}
	rsp.List = rspList
	return &rsp, nil
}
