package router

import "net/http"

type RouteGroup struct {
	BasePath    string
	Middlewares []func(http.Handler) http.Handler
	Routes      []Route
}

type Route struct {
	Path          string
	Method        string
	Middlewares   []func(http.Handler) http.Handler
	Handler       http.HandlerFunc
	SubRoute      *Route
	SubRouteGroup *RouteGroup
}
