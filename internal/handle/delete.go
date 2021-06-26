package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
)

type DeleteF F

func Delete(fs afero.Fs) DeleteF {
	return func(c *gin.Context) (int, error) {
		var request struct {
			Path string `form:"path" binding:"required"`
		}
		err := c.ShouldBindQuery(&request)
		if err != nil {
			return http.StatusBadRequest, err
		}
		err = fs.RemoveAll(request.Path)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return 0, nil
	}
}
