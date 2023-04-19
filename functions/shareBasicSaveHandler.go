package functions

import (
	"github.com/sarailQAQ/faased-netpan/internal/logic/user"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShareBasicSaveHandler(w http.ResponseWriter, r *http.Request) {
	var req types.ShareBasicSaveRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := user.NewShareBasicSaveLogic(r.Context(), svcCtx)
	resp, err := l.ShareBasicSave(&req, r.Header.Get("UserIdentity"))
	if err != nil {
		httpx.Error(w, err)
	} else {
		httpx.OkJson(w, resp)
	}
	
}
