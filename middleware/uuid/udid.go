package uuid

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// JWT is jwt middleware
func GetUuid() gin.HandlerFunc {
	return func(c *gin.Context) {
		udid := c.Query("uuid")
		if udid == "" {
			udid = uuid.Must(uuid.NewV4()).String()
		}

		c.Set("uuid", udid)

		c.Next()
	}
}
