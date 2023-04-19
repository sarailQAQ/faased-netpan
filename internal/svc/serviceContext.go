package svc

import (
	"github.com/sarailQAQ/faased-netpan/internal/config"
	"github.com/sarailQAQ/faased-netpan/models"

	"github.com/go-redis/redis/v8"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	Rdb    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c),
		Rdb:    models.InitRedis(c),
	}
}
