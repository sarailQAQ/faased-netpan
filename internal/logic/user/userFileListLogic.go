package user

import (
	"context"
	"github.com/sarailQAQ/faased-netpan/internal/result"
	"github.com/sarailQAQ/faased-netpan/internal/svc"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"github.com/sarailQAQ/faased-netpan/internal/utils"
	"github.com/sarailQAQ/faased-netpan/models"
)

type UserFileListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListResponse, err error) {
	resp = &types.UserFileListResponse{}
	//获取分页数据
	list, err := models.UserRepository{}.UserFileList(req, userIdentity, l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	//获取条数
	session := l.svcCtx.Engine.NewSession()
	session = session.Table("user_repository").Where("parent_id=? and user_identity=?", req.Id, userIdentity)
	//if req.Type != "all" {
	//	session = session.Table("user_repository").Where("parent_id=? and user_identity=? and type = ?", req.Id, userIdentity, req.Type)
	//
	//}
	count, err := session.Count(new(models.UserRepository))
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	m := make(map[string]interface{})
	m["list"] = list
	m["count"] = count
	resp.Result = result.OK("操作成功", m)
	return
}
