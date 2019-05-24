package uuid

import (
	"github.com/gin-gonic/gin"
	uuidtil "github.com/satori/go.uuid"
)

// JWT is jwt middleware
func GetUuid() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Request.Header.Get("uuid")
		if uuid == "" {
			uuid = uuidtil.Must(uuidtil.NewV4()).String()
			c.Writer.Header().Set("uuid", uuid)
		}

		c.Set("uuid", uuid)

		c.Next()
	}
}
