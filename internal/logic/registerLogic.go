package logic

import (
	"context"

	"github.com/sarailQAQ/faased-netpan/internal/result"
	"github.com/sarailQAQ/faased-netpan/internal/svc"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"github.com/sarailQAQ/faased-netpan/internal/utils"
	"github.com/sarailQAQ/faased-netpan/models"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	resp = new(types.RegisterResponse)
	// 判断code是否一致
	r, err := l.svcCtx.Rdb.Get(l.ctx, req.Email).Result()
	if err != nil {
		resp.Result = result.ERROR("该邮箱验证码不存在")
		return resp, nil
	}
	if r != req.Code {
		resp.Result = result.ERROR("验证码错误")
		return resp, nil
	}
	//判断用户名是否存在
	u, _ := models.User{}.GetUserByUsername(req.UserName, l.svcCtx.Engine)
	if u != nil {
		resp.Result = result.ERROR("用户名已存在")
		return resp, nil
	}
	user := &models.User{
		UserName: req.UserName,
		Password: utils.Md5ToString(req.Password),
		Identity: utils.GetUUID(),
		Email:    req.Email,
	}
	insert, err := models.User{}.Insert(user, l.svcCtx.Engine)
	if err != nil {
		return nil, err
	}
	if insert < 0 {
		resp.Result = result.ERROR("保存失败")
		return resp, nil
	}
	resp.Result = result.OK("申请成功")
	return
}
