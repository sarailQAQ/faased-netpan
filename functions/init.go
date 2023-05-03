package functions

import (
	"github.com/sarailQAQ/faased-netpan/internal/config"
	"github.com/sarailQAQ/faased-netpan/internal/svc"
)

var svcCtx *svc.ServiceContext

func InitSvc(c *config.Config) {
	svcCtx = svc.NewServiceContext(*c)
}
