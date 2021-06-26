package handle

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
)

type UploadF F

func Upload(fs afero.Fs) UploadF {
	return func(c *gin.Context) (int, error) {
		var request struct {
			Path string                `form:"path" binding:"required"`
			File *multipart.FileHeader `form:"file" binding:"required"`
		}
		err := c.ShouldBind(&request)
		if err != nil {
			return http.StatusBadRequest, err
		}
		uf, err := request.File.Open()
		if err != nil {
			return http.StatusInternalServerError, err
		}
		defer uf.Close()
		tmpPath := filepath.Join(request.Path, fmt.Sprintf(".%s.tmp", request.File.Filename))
		defer func() {
			fs.Remove(tmpPath)
		}()
		f, err := fs.Create(tmpPath)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		defer f.Close()
		_, err = io.Copy(f, uf)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		err = f.Close()
		if err != nil {
			return http.StatusInternalServerError, err
		}
		err = fs.Rename(tmpPath, filepath.Join(request.Path, request.File.Filename))
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return 0, nil
	}
}
