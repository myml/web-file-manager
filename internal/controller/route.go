package controller

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/myml/web-file-manager/internal/handle"
)

var Set = wire.NewSet(NewEngine, wire.Struct(new(API), "*"))

func NewEngine(uiFS fs.FS, api API) *gin.Engine {
	engine := gin.Default()
	r := engine.Group("api")
	r.GET("/dl/:name", handle.WarpF(api.DownloadF))
	r.GET("/file", handle.WarpF(api.ListF))
	r.POST("/file", handle.WarpF(api.UploadF))
	r.POST("/file/move", handle.WarpF(api.MoveF))
	r.POST("/file/copy", handle.WarpF(api.CopyFileF))
	r.POST("/file/mkdir", handle.WarpF(api.MkdirF))
	r.DELETE("/file", handle.WarpF(api.DeleteF))

	engine.StaticFS("app", http.FS(uiFS))
	engine.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.String(http.StatusNotFound, "no route")
			return
		}
		c.FileFromFS("/", http.FS(uiFS))
	})
	return engine
}
