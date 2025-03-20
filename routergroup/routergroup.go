package routergroup

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func chainMiddleware(final http.Handler, mws ...Middleware) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		final = mws[i](final)
	}
	return final
}

type RouterGroup struct {
	Prefix      string
	Middlewares []Middleware
	Mux         *http.ServeMux
}

func NewRouterGroup(prefix string, mux *http.ServeMux) *RouterGroup {
	return &RouterGroup{
		Prefix: prefix,
		Mux:    mux,
	}
}

func (g *RouterGroup) Use(mw Middleware) {
	g.Middlewares = append(g.Middlewares, mw)
}

func (g *RouterGroup) HandleFunc(method string, pattern string, handlerFunc http.HandlerFunc) {
	prefixed_pattern := method + " " + g.Prefix + pattern
	wrappedHandler := chainMiddleware(handlerFunc, g.Middlewares...)

	g.Mux.Handle(prefixed_pattern, wrappedHandler)
}
