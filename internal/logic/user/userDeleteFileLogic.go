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

type UserDeleteFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteFileLogic {
	return &UserDeleteFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteFileLogic) UserDeleteFile(req *types.UserDeleteFileRequest, userIdentity string) (resp *types.UserDeleteFileResponse, err error) {
	resp = &types.UserDeleteFileResponse{}
	ur := new(models.UserRepository)
	ur.Identity = req.Identity
	ur.UserIdentity = userIdentity
	//查询判断文件夹下面是否有文件
	u, err := ur.GetByIdentityAndUserIdentity(l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	if u.Id != 0 {
		count, err := ur.GetParentIdCount(u.Id, l.svcCtx.Engine)
		if err != nil {
			resp.Result = result.ERROR(utils.FormatErrorLog(err))
			return resp, nil
		}
		if count > 0 {
			resp.Result = result.ERROR("该文件夹下存在文件，请先删除或者转移")
			return resp, nil

		}
	}
	//删除用户文件
	_, err = ur.Delete(l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	resp.Result = result.OK()
	return
}
