package http

import (
	"fmt"
	"net/http"
	"time"
)

const (
	VERSION = "0.0.1"
)

// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)

type Engine struct {
	handlers []HandlerFunc
	index int8
}

func New() *Engine {
	engine := &Engine{}
	return engine
}

func Default() *Engine {
	engine := New()
	engine.Use(Logger())
	//app.Use(logging, http)
	return engine
}

// Use attaches a global middleware to the router. ie. the middleware attached though Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
// For example, this is the right place for a logger or error management middleware.
func (engine *Engine) Use(middleware HandlerFunc) {
	engine.handlers = append(engine.handlers, middleware)
}

func (engine *Engine) Run(port string) (err error) {
	err = http.ListenAndServe(port, engine)
	return
}

func (engine *Engine) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	ctx := &Context{
		Request:    req,
		Writer:     wr,
		Params:     nil,
		index:      0,
		handlers:   engine.handlers,
		method:     "",
		RouterPath: "",
	}
	engine.handleContext(ctx)
}

func (engine *Engine) handleContext(ctx *Context)  {
	ctx.RouterPath = ctx.Request.URL.Path
	ctx.method = ctx.Request.Method
	fmt.Println(ctx.handlers)
	ctx.Next()
}

func Logger() HandlerFunc {
	return func(ctx *Context) {
		fmt.Print("Logger")
		_ = time.Now()
		time.Sleep(time.Second * 5)
		fmt.Println(ctx)
		_ = time.Now()
		//dura := end - start
		fmt.Printf("spent time :", time.Now())
		ctx.Next()
	}
}



