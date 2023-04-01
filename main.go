package main

import (
	"github.com/deeper-x/gogpt/web"
)

func main() {
	s := web.NewServer()

	s.GET("/input/:key/:query", web.Query)

	s.Run()
}
