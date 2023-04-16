package user

import (
	"cloud-disk/internal/svc"
	"cloud-disk/internal/types"
	"cloud-disk/models"
	"cloud-disk/result"
	"cloud-disk/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameEditLogic {
	return &UserFileNameEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameEditLogic) UserFileNameEdit(req *types.UserFileNameEditRequest, userIdentity string) (resp *types.UserFileNameEditResponse, err error) {
	resp = &types.UserFileNameEditResponse{}
	ur := new(models.UserRepository)
	//判断用户文件是否存在
	count := ur.GetFileNameByUser(req, l.svcCtx.Engine)
	if count > 0 {
		resp.Result = result.ERROR("文件名称已存在")
		return resp, nil
	}
	//进行修改
	ur.UserIdentity = userIdentity
	ur.Identity = req.Identity
	ur.Name = req.Name
	_, err = ur.Edit(l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return nil, err
	}
	resp.Result = result.OK()
	return
}
