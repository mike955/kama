package main

import "kama/pkg/net/http"

func main() {
	r := http.Default()
	r.Run(":8081")
}
