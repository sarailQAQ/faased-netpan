package user

import (
	"cloud-disk/models"
	"cloud-disk/result"
	"cloud-disk/utils"
	"context"

	"cloud-disk/internal/svc"
	"cloud-disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMovedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMovedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMovedLogic {
	return &UserFileMovedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMovedLogic) UserFileMoved(req *types.UserFileMovedRequest, userIdentity string) (resp *types.UserFileMovedResponse, err error) {
	resp = &types.UserFileMovedResponse{}
	//获取ParentId
	ur := new(models.UserRepository)
	ur.Identity = req.ParentIdentity
	ur.UserIdentity = userIdentity
	userRepository, err := ur.GetByIdentityAndUserIdentity(l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	if userRepository.Id == 0 {
		resp.Result = result.ERROR("文件夹不存在")
		return resp, nil
	}
	//修改记录
	ur = new(models.UserRepository)
	ur.Identity = req.Identity
	ur.UserIdentity = userIdentity
	ur.ParentId = int64(rune(userRepository.Id))
	_, err = ur.Edit(l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	resp.Result = result.OK()
	return
}
