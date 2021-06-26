package handle

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
)

type ListF F

func List(fs afero.Fs) ListF {
	return func(c *gin.Context) (int, error) {
		var request struct {
			Path     string `form:"path" validate:"required"`
			ShowHide bool   `form:"show_hide"`
		}
		err := c.ShouldBindQuery(&request)
		if err != nil {
			return http.StatusBadRequest, err
		}
		info, err := fs.Stat(request.Path)
		if err != nil {
			return http.StatusNotFound, err
		}
		if !info.IsDir() {
			c.FileFromFS(request.Path, afero.NewHttpFs(fs))
			return 0, nil
		}
		infos, err := afero.ReadDir(fs, request.Path)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		type FileInfo struct {
			Name     string      `json:"name"` // base name of the file
			Fullname string      `json:"fullname"`
			Ext      string      `json:"ext"`
			Size     int64       `json:"size"`     // length in bytes for regular files; system-dependent for others
			Mode     os.FileMode `json:"mode"`     // file mode bits
			ModTime  time.Time   `json:"mod_time"` // modification time
			IsDir    bool        `json:"is_dir"`   // abbreviation for Mode().IsDir()
		}
		result := make([]*FileInfo, 0, len(infos))
		for i := range infos {
			info := FileInfo{
				Name:    infos[i].Name(),
				Size:    infos[i].Size(),
				Mode:    infos[i].Mode(),
				ModTime: infos[i].ModTime(),
				IsDir:   infos[i].IsDir(),
			}
			info.Fullname = filepath.Join(request.Path, info.Name)
			info.Ext = filepath.Ext(infos[i].Name())
			if info.Name[0] == '.' && !request.ShowHide {
				continue
			}
			result = append(result, &info)
		}
		c.JSON(http.StatusOK, result)
		return 0, nil
	}
}
