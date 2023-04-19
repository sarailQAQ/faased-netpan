package functions

import (
	"github.com/sarailQAQ/faased-netpan/internal/logic/user"
	"github.com/sarailQAQ/faased-netpan/internal/types"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserRepostoryByIdHandler(w http.ResponseWriter, r *http.Request) {
	var req types.GetUserRepostoryByIdRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := user.NewGetUserRepostoryByIdLogic(r.Context(), svcCtx)
	resp, err := l.GetUserRepostoryById(&req, r.Header.Get("UserIdentity"))
	atoi, _ := strconv.Atoi(r.URL.Query().Get("id"))
	req.Id = atoi
	if err != nil {
		httpx.Error(w, err)
	} else {
		httpx.OkJson(w, resp)
	}
	
}
