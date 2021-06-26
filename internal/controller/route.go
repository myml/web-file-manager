package controller

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myml/web-file-manager/internal/handle"
)

func NewEngine(uiFS fs.FS, api API) *gin.Engine {
	engine := gin.Default()
	r := engine.Group("api")
	r.GET("/dl/:name", handle.WarpF(api.DownloadF))
	r.GET("/file", handle.WarpF(api.ListF))
	r.POST("/file", handle.WarpF(api.UploadF))
	r.POST("/file/move", handle.WarpF(api.MoveF))
	r.POST("/file/mkdir", handle.WarpF(api.MkdirF))
	r.DELETE("/file", handle.WarpF(api.DeleteF))

	engine.StaticFS("app", http.FS(uiFS))
	engine.NoRoute(func(c *gin.Context) {
		c.FileFromFS("/", http.FS(uiFS))
	})
	return engine
}
