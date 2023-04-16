package user

import (
	"cloud-disk/models"
	"cloud-disk/result"
	"cloud-disk/utils"
	"context"
	"fmt"

	"cloud-disk/internal/svc"
	"cloud-disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositoryLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositoryLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositoryLinkLogic {
	return &UserRepositoryLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositoryLinkLogic) UserRepositoryLink(req *types.UserRepositoryLinkRequest, userIdentity string) (resp *types.UserRepositoryLinkResponse, err error) {
	resp = &types.UserRepositoryLinkResponse{}
	ur := &models.UserRepository{
		Identity:           utils.GetUUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                req.Ext,
		Name:               req.Name,
		Type:               req.Type,
	}
	//先查询文件是否已经存在关联
	identity, _ := ur.GetByRepositoryIdentityAndUserIdentity(l.svcCtx.Engine)
	if identity > 0 {
		resp.Result = result.ERROR(fmt.Sprintf("%s该文件已经存在！", ur.Name))
		return resp, nil
	}
	_, err = ur.Insert(l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(fmt.Sprintf("数据库发生异常%s", err.Error()))
		return resp, nil
	}
	resp.Result = result.OK("操作成功")
	return
}
