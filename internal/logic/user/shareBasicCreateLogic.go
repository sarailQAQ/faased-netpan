package user

import (
	"context"
	"github.com/sarailQAQ/faased-netpan/internal/result"
	"github.com/sarailQAQ/faased-netpan/internal/svc"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"github.com/sarailQAQ/faased-netpan/internal/utils"
	"github.com/sarailQAQ/faased-netpan/models"
)

type ShareBasicCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
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
