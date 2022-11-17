package gin

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func Test_main(t *testing.T) {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Example": "Hello Gin",
		})
	})

	r.Run()
}
