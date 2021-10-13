package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ilovesusu/suim/api/user/service/v1"
	"github.com/ilovesusu/suim/app/user/service/internal/biz"
	"github.com/ilovesusu/suim/app/user/service/internal/pkg"
	"time"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateUser 注册用户
func (uc *UserRepo) CreateUser(ctx context.Context, user *biz.UserInfo) error {
	var (
		num int64
		err error
	)
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	//判断手机号是否被使用过
	if err = uc.data.db.Model(user).Where("phone = ? and delete_time is null", user.Phone).Count(&num).Error; err != nil {
		uc.log.Log(log.LevelError, "Error", err)
		return err
	}
	if num > 0 {
		uc.log.Log(log.LevelError, "手机号已被使用", *user.Phone)
		return v1.ErrorPhoneIsUsed("手机号%s已被使用", *user.Phone)
	}
	//生成随机用户号码
	for {
		user.Number, err = pkg.RandNumber(uc.log)
		if err != nil {
			return err
		}
		//判断此用户号码是否已被使用过
		if err = uc.data.db.Model(user).Where("number = ? and delete_time is null", user.Number).Count(&num).Error; err != nil {
			uc.log.Log(log.LevelError, "Error", err)
			return err
		}
		if num > 0 {
			continue
		}
		break
	}
	if err = uc.data.db.Model(user).Create(user).Error; err != nil {
		uc.log.Log(log.LevelError, "Error", err)
		return err
	}
	return nil
}

// UpdateUser 更新用户信息
func (uc *UserRepo) UpdateUser(ctx context.Context, user *biz.UserInfo) error {
	var (
		db  = uc.data.db.Begin()
		err error
	)
	defer func() {
		if err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}()
	if user.Phone != nil {
		if err = db.Model(user).Where("id = ?", user.Id).Update("phone", user.Phone).Error; err != nil {
			uc.log.Log(log.LevelError, "Error", err)
			return err
		}
	}
	if user.Password != nil {
		if err = db.Model(user).Where("id = ?", user.Id).Update("password", user.Password).Error; err != nil {
			uc.log.Log(log.LevelError, "Error", err)
			return err
		}
	}
	if user.Name != nil {
		if err = db.Model(user).Where("id = ?", user.Id).Update("name", user.Name).Error; err != nil {
			uc.log.Log(log.LevelError, "Error", err)
			return err
		}
	}
	if user.IdCard != nil {
		if err = db.Model(user).Where("id = ?", user.Id).Update("id_card", user.IdCard).Error; err != nil {
			uc.log.Log(log.LevelError, "Error", err)
			return err
		}
	}
	if user.Nickname != nil {
		if err = db.Model(user).Where("id = ?", user.Id).Update("nickname", user.Nickname).Error; err != nil {
			uc.log.Log(log.LevelError, "Error", err)
			return err
		}
	}
	if user.Sex != nil {
		if err = db.Model(user).Where("id = ?", user.Id).Update("sex", user.Sex).Error; err != nil {
			uc.log.Log(log.LevelError, "Error", err)
			return err
		}
	}
	if user.AvatarUrl != nil {
		if err = db.Model(user).Where("id = ?", user.Id).Update("avatar_url", user.AvatarUrl).Error; err != nil {
			uc.log.Log(log.LevelError, "Error", err)
			return err
		}
	}
	if user.Introduce != nil {
		if err = db.Model(user).Where("id = ?", user.Id).Update("introduce", user.Introduce).Error; err != nil {
			uc.log.Log(log.LevelError, "Error", err)
			return err
		}
	}
	if user.SnapCall != nil {
		if err = db.Model(user).Where("id = ?", user.Id).Update("snap_call", user.SnapCall).Error; err != nil {
			uc.log.Log(log.LevelError, "Error", err)
			return err
		}
	}
	if user.AddFriendType != nil {
		if err = db.Model(user).Where("id = ?", user.Id).Update("add_friend_type", user.AddFriendType).Error; err != nil {
			uc.log.Log(log.LevelError, "Error", err)
			return err
		}
	}
	if user.FriendPassProblem != nil {
		if err = db.Model(user).Where("id = ?", user.Id).Update("friend_pass_problem", user.FriendPassProblem).Error; err != nil {
			uc.log.Log(log.LevelError, "Error", err)
			return err
		}
	}
	if user.FriendPassAnswer != nil {
		if err = db.Model(user).Where("id = ?", user.Id).Update("friend_pass_answer", user.FriendPassAnswer).Error; err != nil {
			uc.log.Log(log.LevelError, "Error", err)
			return err
		}
	}
	return nil
}

// DeleteUser 注销用户
func (uc *UserRepo) DeleteUser(ctx context.Context, user *biz.UserInfo) error {
	var (
		now = time.Now()
		err error
	)
	user.DeleteTime = &now
	if err = uc.data.db.Model(user).Updates(user).Error; err != nil {
		uc.log.Log(log.LevelError, "Error", err)
		return err
	}
	return nil
}

// ListUser 用户列表
func (uc *UserRepo) ListUser(ctx context.Context, friend *biz.UserFriend) ([]biz.FriendList, error) {
	var (
		list  = make([]biz.FriendList, 0) //好友列表
		ids   = make([]int64, 0)          //好友id
		fr    = make(map[int64]*string)   //好友备注map
		users = make([]biz.UserInfo, 0)   //用户列表
		uf    = make([]biz.UserFriend, 0) //用户好友列表
		err   error
	)
	//查找符合条件的好友
	if err = uc.data.db.Model(friend).Where("uid = ? and friend_status = ? and delete_time is null", friend.Uid, friend.FriendStatus).Find(&uf).Error; err != nil {
		uc.log.Log(log.LevelError, "Error", err)
		return nil, err
	}
	for _, v := range uf {
		fr[v.Fid] = v.FriendRemark
		ids = append(ids, v.Fid)
	}
	//查找好友基本信息
	if err = uc.data.db.Model(&biz.UserInfo{}).Where("id in ? and delete_time is null", ids).Find(&users).Error; err != nil {
		uc.log.Log(log.LevelError, "Error", err)
		return nil, err
	}
	//组装返回信息
	for _, v := range users {
		list = append(list, biz.FriendList{
			Id:           v.Id,
			Nickname:     *v.Nickname,
			Sex:          *v.Sex,
			AvatarUrl:    v.AvatarUrl,
			Introduce:    v.Introduce,
			FriendRemark: fr[v.Id],
		})
	}
	return list, nil
}

// InfoUser 用户信息
func (uc *UserRepo) InfoUser(ctx context.Context, user *biz.UserInfo) (*biz.UserInfo, error) {
	var err error
	//查找用户信息
	if err = uc.data.db.Model(user).Where("id = ?", user.Id).First(user).Error; err != nil {
		uc.log.Log(log.LevelError, "Error", err)
		return nil, err
	}
	return user, nil
}
