package functions

import (
	"github.com/sarailQAQ/faased-netpan/internal/logic/user"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFileMovedHandler(w http.ResponseWriter, r *http.Request) {
	var req types.UserFileMovedRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := user.NewUserFileMovedLogic(r.Context(), svcCtx)
	resp, err := l.UserFileMoved(&req, r.Header.Get("UserIdentity"))
	if err != nil {
		httpx.Error(w, err)
	} else {
		httpx.OkJson(w, resp)
	}

}
