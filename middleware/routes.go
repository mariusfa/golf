package middleware

import (
	"net/http"

	"github.com/mariusfa/golf/logging/accesslog"
)

func privateRoute(endpoint http.Handler, authParams AuthParams) http.Handler {
	middlewareWrapper := Auth(endpoint, authParams)
	middlewareWrapper = RequestIdMiddleware(middlewareWrapper)
	middlewareWrapper = AccessLogMiddleware(middlewareWrapper, accesslog.GetLogger())
	return middlewareWrapper
}

func publicRoute(endpoint http.Handler) http.Handler {
	middlewareWrapper := RequestIdMiddleware(endpoint)
	middlewareWrapper = AccessLogMiddleware(middlewareWrapper, accesslog.GetLogger())
	return middlewareWrapper
}
