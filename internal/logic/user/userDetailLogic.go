package user

import (
	"cloud-disk/models"
	"cloud-disk/result"
	"context"

	"cloud-disk/internal/svc"
	"cloud-disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
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
