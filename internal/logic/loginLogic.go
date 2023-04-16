package logic

import (
	"cloud-disk/define"
	"cloud-disk/internal/svc"
	"cloud-disk/internal/types"
	"cloud-disk/models"
	"cloud-disk/result"
	"cloud-disk/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
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
