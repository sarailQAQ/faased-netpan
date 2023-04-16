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

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest, userIdentity string) (resp *types.ShareBasicSaveResponse, err error) {
	resp = &types.ShareBasicSaveResponse{}
	//获取资源数据
	rp := new(models.RepositoryPool)
	rp, err = rp.GetByIdentity(req.RepositoryIdentity, l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	if rp.Id == 0 {
		resp.Result = result.ERROR("资源不存在")
		return resp, nil
	}
	//保存资源
	ur := &models.UserRepository{
		Identity:           utils.GetUUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: rp.Identity,
		Ext:                rp.Ext,
		Name:               rp.Name,
	}
	_, err = ur.Insert(l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	m := make(map[string]interface{})
	m["identity"] = ur.Identity
	resp.Result = result.OK(m)
	return
}
