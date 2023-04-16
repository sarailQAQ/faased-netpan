package user

import (
	"cloud-disk/internal/svc"
	"cloud-disk/internal/types"
	"cloud-disk/models"
	"cloud-disk/result"
	"cloud-disk/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
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
	if req.Type != "all" {
		session = session.Table("user_repository").Where("parent_id=? and user_identity=? and type = ?", req.Id, userIdentity, req.Type)

	}
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
