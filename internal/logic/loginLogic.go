package logic

import (
	"context"
	"github.com/sarailQAQ/faased-netpan/internal/define"
	"github.com/sarailQAQ/faased-netpan/internal/result"
	"github.com/sarailQAQ/faased-netpan/internal/svc"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"github.com/sarailQAQ/faased-netpan/internal/utils"
	"github.com/sarailQAQ/faased-netpan/models"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	m := make(map[string]interface{}, 0)
	resp = new(types.LoginResponse)
	// 登录逻辑
	user := new(models.User)
	// 读取数据库数据
	user, err = user.GetUserByUsername(req.UserName, l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(err.Error())
		return resp, nil
	}
	if user.Password != utils.Md5ToString(req.Password) {
		resp.Result = result.ERROR("密码错误")
		return resp, nil
	}
	//生成token
	err, s := utils.GenerateToken(user.Id, user.Identity, user.UserName, define.TokenExpire)
	if err != nil {
		return nil, err
	}
	//生成一个刷新token的GenerateToken
	err, refreshToke := utils.GenerateToken(user.Id, user.Identity, user.UserName, define.RefreshTokenExpire)
	m["token"] = s
	m["refreshToke"] = refreshToke
	resp.Result = result.OK(m)
	return
}
