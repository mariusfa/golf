package middleware

import (
	"net/http"
)

func privateRoute(endpoint http.Handler, authParams AuthParams) http.Handler {
	middlewareWrapper := Auth(endpoint, authParams)
	return middlewareWrapper
}
