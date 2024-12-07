package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	// return c.Render("welcome", nil, "layouts/main")
	c.HTML(http.StatusOK, "welcome.html", nil)
}
