package handler_users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	c.Writer.WriteString("Implement me")
	//_, _ = fmt.Fprint(c.Writer, "Implement me")
}

func Search(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me")
}
