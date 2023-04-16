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

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateResponse, err error) {
	resp = &types.ShareBasicCreateResponse{}
	//查询用户文件池文件
	ur := new(models.UserRepository)
	ur.Identity = req.UserRepositoryIdentity
	ur.UserIdentity = userIdentity
	userRepository, err := ur.GetByIdentityAndUserIdentity(l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	if userRepository.Id == 0 {
		resp.Result = result.ERROR("用户分享文件不存在")
		return resp, nil
	}
	//保存用户分享
	data := &models.ShareBasic{
		Identity:               utils.GetUUID(),
		UserIdentity:           userIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		RepositoryIdentity:     userRepository.RepositoryIdentity,
		ExpiredTime:            req.ExpiredTime,
	}
	_, err = data.Insert(l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	m := make(map[string]interface{})
	m["identity"] = data.Identity
	resp.Result = result.OK(m)
	return
}
