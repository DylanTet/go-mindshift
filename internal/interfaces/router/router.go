package router

import (
  "net/http"
)

type Router struct {
  routes map[string]map[string]http.HandlerFunc // method -> path -> handler
}

func NewRouter() *Router {
  r := &Router{
    routes: make(map[string]map[string]http.HandlerFunc),
  }

  // Add all backend routes
  r.POST("/create-user", func(w http.ResponseWriter, req *http.Request) {

  })

  r.GET("/get-user", func(w http.ResponseWriter, req *http.Request) {

  })

  return r
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  method := req.Method
  path := req.URL.Path

  if methodRoutes, exists := r.routes[method]; exists {
    if handler, exists := methodRoutes[path]; exists {
      handler(w, req)
      return 
    }
  }

  http.NotFound(w, req)
}

func (r *Router) GET(path string, handler http.HandlerFunc) {
  r.addRoute("GET", path, handler)
}

func (r *Router) POST(path string, handler http.HandlerFunc) {
  r.addRoute("POST", path, handler)
}

// For middleware stuff
func (r *Router) Use(lambda func()) {

}

func (r *Router) addRoute(method string, path string, handler http.HandlerFunc) {
  if r.routes[method] == nil {
    r.routes[method] = make(map[string]http.HandlerFunc)
  }

  r.routes[method][path] = handler
}

