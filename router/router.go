package router

import "net/http"

type Router interface {
	RegisterRouteGroup(rg RouteGroup)
	GetMux() http.Handler
	AddMiddleWare(m func(http.Handler) http.Handler)
}
