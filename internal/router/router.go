package router

import (
	"net/http"
	"strings"
)

type Router struct {
	routes map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) Register(path string, handler http.HandlerFunc) {
	r.routes[path] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// 1. Try exact match first
	if handler, exists := r.routes[req.URL.Path]; exists {
		handler(w, req)
		return
	}

	// 2. Try dynamic route matching
	for route, handler := range r.routes {

		if matchDynamicRoute(route, req.URL.Path) {
			handler(w, req)
			return
		}
	}

	// 3. No route matched
	http.NotFound(w, req)
}

func matchDynamicRoute(route string, path string) bool {

	routeParts := strings.Split(route, "/")
	pathParts := strings.Split(path, "/")

	if len(routeParts) != len(pathParts) {
		return false
	}

	for i := range routeParts {

		if strings.HasPrefix(routeParts[i], ":") {
			continue
		}

		if routeParts[i] != pathParts[i] {
			return false
		}
	}

	return true
}