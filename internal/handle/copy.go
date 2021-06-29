package handle

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
)

type CopyFileF F

func CopyFile(fs afero.Fs) CopyFileF {
	return func(c *gin.Context) (int, error) {
		var request struct {
			OldPath string `json:"old_path" binding:"required"`
			NewPath string `json:"new_path" binding:"required"`
		}
		err := c.ShouldBindJSON(&request)
		if err != nil {
			return http.StatusBadRequest, err
		}
		exists, err := afero.Exists(fs, request.NewPath)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		if exists {
			return http.StatusForbidden, fmt.Errorf("move: %w", ErrExists)
		}
		dst, err := os.Create(request.NewPath)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		defer dst.Close()
		src, err := fs.Open(request.OldPath)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		defer src.Close()
		_, err = io.Copy(dst, src)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		err = dst.Close()
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return 0, nil
	}
}
