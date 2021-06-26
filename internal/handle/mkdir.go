package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
)

type MkdirF F

func Mkdir(fs afero.Fs) MkdirF {
	return func(c *gin.Context) (int, error) {
		var request struct {
			Path string `json:"path" binding:"required"`
		}

		err := c.ShouldBindJSON(&request)
		if err != nil {
			return http.StatusBadRequest, err
		}
		err = fs.MkdirAll(request.Path, 0700)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return 0, nil
	}
}
