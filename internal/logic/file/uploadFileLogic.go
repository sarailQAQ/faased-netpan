package file

import (
	"cloud-disk/models"
	"cloud-disk/result"
	"cloud-disk/utils"
	"context"

	"cloud-disk/internal/svc"
	"cloud-disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileLogic) UploadFile(req *types.UploadFileRequest) (resp *types.UploadFileResponse, err error) {
	rp := &models.RepositoryPool{
		Identity: utils.GetUUID(),
		Hash:     req.Hash,
		Name:     req.Name,
		Ext:      req.Ext,
		Size:     req.Size,
		Path:     req.Path,
		Type:     utils.GetFileType(req.Ext),
	}
	_, err = rp.Insert(l.svcCtx.Engine)
	if err != nil {
		return nil, err
	}
	resp = new(types.UploadFileResponse)
	m := make(map[string]string)
	m["identity"] = rp.Identity
	m["ext"] = rp.Ext
	m["name"] = rp.Name
	m["type"] = rp.Type
	resp.Result = result.OK(m)
	return
}
