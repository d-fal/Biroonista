package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnSite(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"body":  "Biroonista",
		"title": "Biroonista App",
	})
}
