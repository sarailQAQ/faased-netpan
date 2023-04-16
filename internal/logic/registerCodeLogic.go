package logic

import (
	"context"
	"time"

	"github.com/sarailQAQ/faased-netpan/internal/define"
	"github.com/sarailQAQ/faased-netpan/internal/result"
	"github.com/sarailQAQ/faased-netpan/internal/svc"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"github.com/sarailQAQ/faased-netpan/internal/utils"
	"github.com/sarailQAQ/faased-netpan/models"
)

type RegisterCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterCodeLogic {
	return &RegisterCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterCodeLogic) RegisterCode(req *types.GetCodeRequest, email string) (resp *types.GetCodeResponse, err error) {
	resp = &types.GetCodeResponse{}
	//查询邮箱是否被注册
	count := models.User{}.GetUserByEmailCount(email, l.svcCtx.Engine)
	if count > 0 {
		resp.Result = result.ERROR("邮箱已经被注册")
		return resp, nil
	}
	//生成随机验证码
	code := utils.RandCode()
	//存储验证码
	l.svcCtx.Rdb.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExpire))
	//发送验证码
	err = utils.MailSendCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	resp.Result = result.OK("操作成功")
	return
}
