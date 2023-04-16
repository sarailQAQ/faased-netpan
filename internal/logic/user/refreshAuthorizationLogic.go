package user

import (
	"context"
	"github.com/sarailQAQ/faased-netpan/internal/define"
	"github.com/sarailQAQ/faased-netpan/internal/result"
	"github.com/sarailQAQ/faased-netpan/internal/svc"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"github.com/sarailQAQ/faased-netpan/internal/utils"
)

type RefreshAuthorizationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthorizationLogic {
	return &RefreshAuthorizationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthorizationLogic) RefreshAuthorization(authorization string) (resp *types.RefreshAuthorizationResponse, err error) {
	resp = &types.RefreshAuthorizationResponse{}
	//刷新GenerateToken
	token, err := utils.AnalyzeToken(authorization)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	err, newToken := utils.GenerateToken(token.Id, token.Identity, token.Name, define.TokenExpire)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	err, newRefreshToken := utils.GenerateToken(token.Id, token.Identity, token.Name, define.RefreshTokenExpire)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	m := make(map[string]interface{})
	m["token"] = newToken
	m["refreshToke"] = newRefreshToken
	resp.Result = result.OK(m)
	return
}
