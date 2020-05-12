package middlewares

import (
	"fmt"
	"kama/pkg/net/http"
)

func QueryParser() http.HandlerFunc {
	return func(ctx *http.Context) {
		fmt.Print(ctx.Request.URL.Query())
		ctx.Next()
	}
}
