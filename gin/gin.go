package gin

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

// implement the interface of ServeHTTP
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}

// New is the constructor of gin.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(patten string, handlerFunc HandlerFunc) {
	engine.addRoute("POST", patten, handlerFunc)
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}
