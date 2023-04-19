package functions

import (
	"github.com/sarailQAQ/faased-netpan/internal/logic"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req types.RegisterRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewRegisterLogic(r.Context(), svcCtx)
	resp, err := l.Register(&req)
	if err != nil {
		httpx.Error(w, err)
	} else {
		httpx.OkJson(w, resp)
	}
}
