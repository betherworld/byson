package commonhandlers

import (
	"fin4-core/server/datatype"
	"github.com/gin-gonic/gin"
)

// Index index route
func Index(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.File("public/index.html")
	}
}
