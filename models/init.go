package models

import (
	"cloud-disk/internal/config"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

func Init(c config.Config) *xorm.Engine {
	engine, err := xorm.NewEngine(c.DataBase.Type, c.DataBase.Url)
	if err != nil {
		log.Panicf("Xorm创建错误%v", err)
		return nil
	}
	engine.ShowSQL(c.DataBase.ShowSql)
	engine.SetMaxIdleConns(c.DataBase.MaxIdleConns)
	engine.SetMaxOpenConns(c.DataBase.MaxOpenConns)
	return engine
}

func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password,
		DB:       0,
		PoolSize: c.Redis.PoolSize,
	})
}
