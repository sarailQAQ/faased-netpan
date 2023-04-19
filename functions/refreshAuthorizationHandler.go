package functions

import (
	"net/http"

	"github.com/sarailQAQ/faased-netpan/internal/logic/user"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RefreshAuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	l := user.NewRefreshAuthorizationLogic(r.Context(), svcCtx)
	resp, err := l.RefreshAuthorization(r.Header.Get("Authorization"))
	if err != nil {
		httpx.Error(w, err)
	} else {
		httpx.OkJson(w, resp)
	}

}
