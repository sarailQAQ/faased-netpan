package functions

import (
	"github.com/sarailQAQ/faased-netpan/internal/logic"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterCodeHandler(w http.ResponseWriter, r *http.Request) {
	var req types.GetCodeRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}
	l := logic.NewRegisterCodeLogic(r.Context(), svcCtx)
	resp, err := l.RegisterCode(&req, r.URL.Query().Get("email"))
	if err != nil {
		httpx.Error(w, err)
	} else {
		httpx.OkJson(w, resp)
	}
}
