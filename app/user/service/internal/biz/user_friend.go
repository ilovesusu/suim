package biz

import (
	"context"
	"github.com/ilovesusu/suim/api/user/service/v1/friend"
)

// UserFriend 用户好友
type UserFriend struct {
	BaseModel
	UserId       int64   `gorm:"index;not null;comment:用户id"`
	FriendId     int64   `gorm:"index;not null;comment:好友id"`
	FriendStatus int32   `gorm:"index;not null;comment:好友状态(1-申请,2-同意,3-拒绝,4-拉黑,5-删除)"`
	ApplyMessage *string `gorm:"comment:申请好友消息"`
	FriendRemark *string `gorm:"comment:好友备注"`
}

// UpdateFriendStatus 修改好友状态
func (f *FriendUsecase) UpdateFriendStatus(ctx context.Context, req *UpdateFriendStatusReq) error {
	if req.FriendStatus == int32(friend.FriendStatus_APPLY_FRIEND) {
		//todo 申请好友和同意好友需要通知
	}
	return f.repo.UpdateFriendStatus(ctx, req)
}

// UpdateFriendRemark 修改好友备注
func (f *FriendUsecase) UpdateFriendRemark(ctx context.Context, req *UpdateFriendRemarkReq) error {
	return f.repo.UpdateFriendRemark(ctx, req)
}

// ListUserFriend 好友列表
func (f *FriendUsecase) ListUserFriend(ctx context.Context, req *ListFriendReq) ([]ListUserFriendRsp, error) {
	return f.repo.ListUserFriend(ctx, req)
}
