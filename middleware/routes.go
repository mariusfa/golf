package middleware

import (
	"net/http"
)

func PrivateRoute(endpoint http.Handler, authParams AuthParams) http.Handler {
	middlewareWrapper := Auth(endpoint, authParams)
	middlewareWrapper = RequestIdMiddleware(middlewareWrapper)
	middlewareWrapper = AccessLogMiddleware(middlewareWrapper)
	return middlewareWrapper
}

func PublicRoute(endpoint http.Handler) http.Handler {
	middlewareWrapper := RequestIdMiddleware(endpoint)
	middlewareWrapper = AccessLogMiddleware(middlewareWrapper)
	return middlewareWrapper
}
