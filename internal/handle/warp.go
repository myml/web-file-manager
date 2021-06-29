package handle

import "github.com/gin-gonic/gin"

func WarpF(f func(c *gin.Context) (int, error)) gin.HandlerFunc {
	if f == nil {
		panic("warp nil")
	}
	return func(c *gin.Context) {
		code, err := f(c)
		if err != nil {
			c.AbortWithError(code, err)
		}
		if code != 0 {
			c.Status(200)
		}
	}
}

type F func(c *gin.Context) (int, error)
