package src

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

type HandlePDFBody struct {
	Url string `json:"url"`
}

func HandlePDF(c *gin.Context) {
	var requestBody HandlePDFBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	urlPayload := string(requestBody.Url)
	res1 := strings.ReplaceAll(urlPayload, "&&", "")
	res2 := strings.ReplaceAll(res1, "|", "")

	cmd := exec.Command("wkhtmltopdf", res2, "public/report.pdf")
	stdout, err := cmd.Output()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": string(stdout), "success": true})
}

func DisplayBanner(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"usage": "POST / {'url': 'http://google.com'}", "success": true})
}
