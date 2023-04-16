package user

import (
	"context"
	"github.com/sarailQAQ/faased-netpan/internal/result"
	"github.com/sarailQAQ/faased-netpan/internal/svc"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"github.com/sarailQAQ/faased-netpan/models"
)

type UserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserInfoRequest, userIdentity string) (resp *types.UserInfoResponse, err error) {
	// 根据userIdentity获取用户详情
	u := new(models.User)
	userinfo := u.GetUserInfo(userIdentity, l.svcCtx.Engine)
	resp = &types.UserInfoResponse{
		Result: result.OK("操作成功", userinfo),
	}
	return
}
