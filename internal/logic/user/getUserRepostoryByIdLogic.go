package user

import (
	"cloud-disk/models"
	"cloud-disk/result"
	"context"

	"cloud-disk/internal/svc"
	"cloud-disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRepostoryByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRepostoryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRepostoryByIdLogic {
	return &GetUserRepostoryByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRepostoryByIdLogic) GetUserRepostoryById(req *types.GetUserRepostoryByIdRequest, userIdentity string) (resp *types.GetUserRepostoryByIdResponse, err error) {
	resp = &types.GetUserRepostoryByIdResponse{}
	ur := new(models.UserRepository)
	ur.Id = req.Id
	ur.UserIdentity = userIdentity
	ur, _ = ur.GetUserById(l.svcCtx.Engine)
	resp.Result = result.OK(ur)
	return
}
