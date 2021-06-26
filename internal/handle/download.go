package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
)

type DownloadF F

func Download(fs afero.Fs) DownloadF {
	return func(c *gin.Context) (int, error) {
		var request struct {
			Path string `form:"path" binding:"required"`
		}
		err := c.ShouldBindQuery(&request)
		if err != nil {
			return http.StatusBadRequest, err
		}
		info, err := fs.Stat(request.Path)
		if err != nil {
			return http.StatusNotFound, err
		}
		if info.IsDir() {
			return http.StatusNotImplemented, nil
		}
		c.FileFromFS(request.Path, afero.NewHttpFs(fs))
		return 0, nil
	}
}
