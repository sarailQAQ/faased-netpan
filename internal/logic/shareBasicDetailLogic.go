package logic

import (
	"context"

	"github.com/sarailQAQ/faased-netpan/internal/result"
	"github.com/sarailQAQ/faased-netpan/internal/svc"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"github.com/sarailQAQ/faased-netpan/internal/utils"
	"github.com/sarailQAQ/faased-netpan/models"
)

type ShareBasicDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailResponse, err error) {
	resp = &types.ShareBasicDetailResponse{}
	//点击计数
	shareBasic := new(models.ShareBasic)
	err = shareBasic.CountUp(req.Identity, l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	//获取资源详情数据
	basicDetail, err := shareBasic.GetBasicDetail(req.Identity, l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	resp.Result = result.OK(basicDetail)
	return
}
