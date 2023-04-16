package user

import (
	"context"
	"github.com/sarailQAQ/faased-netpan/internal/result"
	"github.com/sarailQAQ/faased-netpan/internal/svc"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"github.com/sarailQAQ/faased-netpan/internal/utils"
	"github.com/sarailQAQ/faased-netpan/models"
)

type UserDeleteFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteFileLogic {
	return &UserDeleteFileLogic{
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
