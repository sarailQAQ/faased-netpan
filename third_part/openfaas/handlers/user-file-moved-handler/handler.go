package function

import (
	"github.com/sarailQAQ/faased-netpan/functions"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	functions.UserFileMovedHandler(w, r)
}
