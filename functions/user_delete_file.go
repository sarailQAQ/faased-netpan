package functions

import (
	"github.com/sarailQAQ/faased-netpan/internal/logic/user"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserDeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	var req types.UserDeleteFileRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := user.NewUserDeleteFileLogic(r.Context(), svcCtx)
	req.Identity = r.URL.Query().Get("identity")
	resp, err := l.UserDeleteFile(&req, r.Header.Get("UserIdentity"))
	if err != nil {
		httpx.Error(w, err)
	} else {
		httpx.OkJson(w, resp)
	}

}
