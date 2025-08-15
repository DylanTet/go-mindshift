package router

import (
  "net/http"
)

type Route interface {

}

type Router struct {
  routes map[string]func(http.Request) http.Response
}

// For middleware stuff
func (r Router) Use(func lambda) {

}

func (r Router) Route(string endpoint, func lambda) {
  
}

func NewRouter() *Router {
  return &Router{}
}
