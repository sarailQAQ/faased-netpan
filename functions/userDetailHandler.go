package functions

import (
	"net/http"

	"github.com/sarailQAQ/faased-netpan/internal/logic/user"
	"github.com/sarailQAQ/faased-netpan/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserDetailHandler(w http.ResponseWriter, r *http.Request) {
	var req types.UserInfoRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := user.NewUserDetailLogic(r.Context(), svcCtx)
	resp, err := l.UserDetail(&req, r.Header.Get("UserIdentity"))
	if err != nil {
		httpx.Error(w, err)
	} else {
		httpx.OkJson(w, resp)
	}

}
