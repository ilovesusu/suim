package data

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ilovesusu/suim/app/user/service/internal/biz"
	"github.com/ilovesusu/suim/app/user/service/internal/conf"
	"os"
	"testing"
)

var (
	UserRepos biz.UserRepo
	flagConf  string
)

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		//"service.id", id,
		//"service.name", Name,
		//"service.version", Version,
		//"trace_id", log.TraceID(),
		//"span_id", log.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagConf),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	dataData, cleanup, err := NewData(bc.Data, logger)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cleanup()
	UserRepos = NewUserRepo(dataData, logger)
}

//新增
func TestUserRepo_CreateUser(t *testing.T) {
	phone := "18298883648"
	password := "123456789"
	name := "谢金武"
	card := "622827199810010916"
	nickname := "武邪"
	sex := int32(1)
	snap := false
	addFriendType := int32(1)
	if err := UserRepos.CreateUser(context.Background(), &biz.UserInfo{
		Phone:         &phone,
		Password:      &password,
		Name:          &name,
		IdCard:        &card,
		Nickname:      &nickname,
		Sex:           &sex,
		SnapCall:      &snap,
		AddFriendType: &addFriendType,
	}); err != nil {
		fmt.Println(err)
		return
	}
}

//更新
func BenchmarkUpdateUser(t *testing.B) {
	phone := "18258255157"
	password := "xiejinwu00000"
	if err := UserRepos.UpdateUser(context.Background(), &biz.UserInfo{
		BaseModel: biz.BaseModel{Id: 1},
		Phone:     &phone,
		Password:  &password,
	}); err != nil {
		fmt.Println(err)
		return
	}
}
