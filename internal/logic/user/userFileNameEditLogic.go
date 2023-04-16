package user

import (
	"context"

	"github.com/sarailQAQ/faased-netpan/internal/result"
	"github.com/sarailQAQ/faased-netpan/internal/svc"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"github.com/sarailQAQ/faased-netpan/internal/utils"
	"github.com/sarailQAQ/faased-netpan/models"
)

type UserFileNameEditLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameEditLogic {
	return &UserFileNameEditLogic{
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
