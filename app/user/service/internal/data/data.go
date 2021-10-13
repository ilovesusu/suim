package data

import (
	"database/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/ilovesusu/suim/app/user/service/internal/biz"
	"github.com/ilovesusu/suim/app/user/service/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	db *gorm.DB
	rc *redis.Client
}

// NewData 数据库连接 .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	sqlDB, err := sql.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		_ = logger.Log(log.LevelError, err)
		return nil, cleanup, err
	}
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "suim_", //表前缀
			SingularTable: true,    //禁用表名复数
		},
	})
	if err != nil {
		_ = logger.Log(log.LevelError, "gorm", err)
		return nil, cleanup, err
	}
	if c.Database.Debug {
		db = db.Debug()
	}
	rc := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Pwd,
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})
	_ = logger.Log(log.LevelInfo, "DATABASE-REDIS", "connect init success!")
	if err = db.AutoMigrate(biz.MigratorTable...); err != nil {
		_ = logger.Log(log.LevelError, "MIGRATOR TABLE", err)
		return nil, nil, err
	}
	_ = logger.Log(log.LevelInfo, "MIGRATOR TABLE", "AutoMigrate table success!")
	return &Data{db: db, rc: rc}, cleanup, nil
}
