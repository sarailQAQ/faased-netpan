package functions

import (
	"net/http"

	"github.com/sarailQAQ/faased-netpan/internal/logic"
	"github.com/sarailQAQ/faased-netpan/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShareBasicDetailHandler(w http.ResponseWriter, r *http.Request) {
	var req types.ShareBasicDetailRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewShareBasicDetailLogic(r.Context(), svcCtx)
	resp, err := l.ShareBasicDetail(&req)
	if err != nil {
		httpx.Error(w, err)
	} else {
		httpx.OkJson(w, resp)
	}
}
