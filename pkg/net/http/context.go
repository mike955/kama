package http

import (
	"fmt"
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
	ctx.index++
	for ctx.index < int8(len(ctx.handlers)) {
		fmt.Println(" ctx.index: ", ctx.index)
		ctx.handlers[ctx.index](ctx)
		ctx.index++
		//fmt.Println(" ctx.index: ", ctx.index)
	}
}

func (ctx *Context) getParams(key string) (interface{}, bool)  {
	value, ok := ctx.Params[key]
	if !ok {
		return nil, false
	}
	return value, true
}

func (ctx *Context) setParams(key string, value interface{}) (bool)  {
	ctx.Params[key] = value
	return true
}

