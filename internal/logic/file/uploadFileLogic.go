package file

import (
	"context"
	"github.com/sarailQAQ/faased-netpan/internal/result"
	"github.com/sarailQAQ/faased-netpan/internal/svc"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"github.com/sarailQAQ/faased-netpan/internal/utils"
	"github.com/sarailQAQ/faased-netpan/models"
)

type UploadFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
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
