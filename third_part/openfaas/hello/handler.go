package function

import (
	"fmt"
	handler "github.com/openfaas/templates-sdk/go-http"
	"github.com/sarailQAQ/faased-netpan/functions"
	"net/http"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	message := fmt.Sprintf("Body: %s", string(req.Body))

	functions.Upload()

	return handler.Response{
		Body:       []byte(message),
		StatusCode: http.StatusOK,
	}, err
}
