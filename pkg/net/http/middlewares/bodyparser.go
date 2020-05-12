package middlewares

import (
	"fmt"
	"kama/pkg/net/http"
	"time"
)

func BodyParser() http.HandlerFunc {
	return func(ctx *http.Context) {
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
