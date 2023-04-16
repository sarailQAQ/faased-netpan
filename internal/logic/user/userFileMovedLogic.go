package user

import (
	"context"

	"github.com/sarailQAQ/faased-netpan/internal/result"
	"github.com/sarailQAQ/faased-netpan/internal/svc"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"github.com/sarailQAQ/faased-netpan/internal/utils"
	"github.com/sarailQAQ/faased-netpan/models"
)

type UserFileMovedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMovedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMovedLogic {
	return &UserFileMovedLogic{
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
