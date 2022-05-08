package src

import "github.com/gin-gonic/gin"

func GetRoute(r *gin.Engine) {
	r.POST("/", HandlePDF)
	r.GET("/", DisplayBanner)
}
