package functions

import (
	"github.com/sarailQAQ/faased-netpan/internal/logic"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req types.LoginRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewLoginLogic(r.Context(), svcCtx)
	resp, err := l.Login(&req)
	if err != nil {
		httpx.Error(w, err)
	} else {
		httpx.OkJson(w, resp)
	}
}
