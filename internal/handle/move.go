package handle

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
)

type MoveF F

func Move(fs afero.Fs) MoveF {
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
			return http.StatusInternalServerError, fmt.Errorf("move: %w", err)
		}
		if exists {
			return http.StatusForbidden, fmt.Errorf("move: %w", ErrExists)
		}
		err = fs.Rename(request.OldPath, request.NewPath)
		if err != nil {
			return http.StatusInternalServerError, fmt.Errorf("move: %w", err)
		}
		return 0, nil
	}
}
