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

type UserDirCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDirCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDirCreateLogic {
	return &UserDirCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDirCreateLogic) UserDirCreate(req *types.UserDirCreateRequest, userIdentity string) (resp *types.UserDirCreateResponse, err error) {
	resp = &types.UserDirCreateResponse{}
	ur := new(models.UserRepository)
	ur.Name = req.Name
	ur.ParentId = req.ParentId
	//判断文件夹是否存在
	userRepository, err := ur.GetByName(l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	if userRepository.Id > 0 {
		resp.Result = result.ERROR("文件夹名称已存在")
		return resp, nil
	}
	ur.Identity = utils.GetUUID()
	ur.UserIdentity = userIdentity
	ur.ParentId = req.ParentId
	ur.Name = req.Name
	_, err = ur.Insert(l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	resp.Result = result.OK()
	return
}
