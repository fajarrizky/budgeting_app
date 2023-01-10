package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ChiRouter struct {
	router *chi.Mux
}

func NewChiRouter() Router {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)

	return &ChiRouter{
		router: r,
	}
}

func (cr *ChiRouter) AddMiddleWare(m func(http.Handler) http.Handler) {
	cr.router.Use(m)
}

func (cr *ChiRouter) registerRoutes(r chi.Router, routes []Route) {
	for _, route := range routes {

		switch route.Method {
		case http.MethodGet:
			r.With(route.Middlewares...).Get(route.Path, route.Handler)
		case http.MethodPost:
			r.With(route.Middlewares...).Post(route.Path, route.Handler)
		case http.MethodPut:
			r.With(route.Middlewares...).Put(route.Path, route.Handler)
		case http.MethodDelete:
			r.With(route.Middlewares...).Delete(route.Path, route.Handler)
		default:
			panic(fmt.Sprintf("cannot register route:- method: %v path: %v\n", route.Method, route.Path))
		}
	}
}

func (cr *ChiRouter) RegisterRouteGroup(rg RouteGroup) {
	cr.router.Route(rg.BasePath, func(r chi.Router) {

		if rg.Middlewares != nil && len(rg.Middlewares) > 0 {

			r.Use(rg.Middlewares...)
		}

		cr.registerRoutes(r, rg.Routes)

	})
}

func (cr *ChiRouter) GetMux() http.Handler {
	return cr.router
}
