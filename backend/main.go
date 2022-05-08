package main

import (
	"wkbackend/src"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	src.GetRoute(r)
	r.Run(":8090")
}
