package http

import (
	"net/http"
)

type Context struct {

	Request *http.Request
	Writer  http.ResponseWriter
	Params map[string]interface{}

	index int8
	handlers []HandlerFunc

	method     string
	RouterPath string
}

func (ctx *Context) Next()  {
	for ctx.index < int8(len(ctx.handlers)) {
		ctx.handlers[ctx.index](ctx)
		ctx.index++
	}
}