package svc

import (
	"cloud-disk/internal/config"
	"cloud-disk/internal/middleware"
	"cloud-disk/models"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	Rdb    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c),
		Rdb:    models.InitRedis(c),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
