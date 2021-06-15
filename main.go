package main

import (
	"embed"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
)

//go:embed dist/web
var web embed.FS

var rootDir string

func rootFS() afero.Fs {
	return afero.NewBasePathFs(afero.NewOsFs(), rootDir)
}

func main() {
	pwd, _ := os.Getwd()
	flag.StringVar(&rootDir, "d", pwd, "root dir")
	var addr string
	flag.StringVar(&addr, "l", ":8080", "listen addr")
	flag.Parse()
	log.Println("root dir:", rootDir)

	engine := gin.Default()
	static, err := fs.Sub(web, "dist/web")
	if err != nil {
		panic(err)
	}
	api := engine.Group("api")
	api.GET("/dl/:name", download)
	api.GET("/file", list)
	api.POST("/file", upload)
	api.POST("/file/move", move)
	api.POST("/file/mkdir", mkdir)
	api.DELETE("/file", delete)

	engine.StaticFS("app", http.FS(static))
	engine.NoRoute(func(c *gin.Context) {
		c.FileFromFS("/", http.FS(static))
	})
	engine.Run(addr)
}

func download(c *gin.Context) {
	var request struct {
		Path string `form:"path" binding:"required"`
	}
	err := c.ShouldBindQuery(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fs := rootFS()
	info, err := fs.Stat(request.Path)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	if info.IsDir() {
		c.AbortWithStatus(http.StatusNotImplemented)
		return
	}
	c.FileFromFS(request.Path, afero.NewHttpFs(fs))
}

func list(c *gin.Context) {
	var request struct {
		Path     string `form:"path" validate:"required"`
		ShowHide bool   `form:"show_hide"`
	}
	err := c.ShouldBindQuery(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fs := rootFS()
	info, err := fs.Stat(request.Path)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	if !info.IsDir() {
		c.FileFromFS(request.Path, afero.NewHttpFs(fs))
		return
	}
	infos, err := afero.ReadDir(fs, request.Path)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
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
}

func move(c *gin.Context) {
	var request struct {
		OldPath string `json:"old_path" binding:"required"`
		NewPath string `json:"new_path" binding:"required"`
	}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fs := rootFS()
	err = fs.Rename(request.OldPath, request.NewPath)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func delete(c *gin.Context) {
	var request struct {
		Path string `form:"path" binding:"required"`
	}
	err := c.ShouldBindQuery(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fs := rootFS()
	err = fs.RemoveAll(request.Path)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func mkdir(c *gin.Context) {
	var request struct {
		Path string `json:"path" binding:"required"`
	}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fs := rootFS()
	err = fs.MkdirAll(request.Path, 0700)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func upload(c *gin.Context) {
	var request struct {
		Path string                `form:"path" binding:"required"`
		File *multipart.FileHeader `form:"file" binding:"required"`
	}
	err := c.ShouldBind(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	uf, err := request.File.Open()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer uf.Close()

	fs := rootFS()
	tmpPath := filepath.Join(request.Path, fmt.Sprintf(".%s.tmp", request.File.Filename))
	defer func() {
		fs.Remove(tmpPath)
	}()
	f, err := fs.Create(tmpPath)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, uf)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = f.Close()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = fs.Rename(tmpPath, filepath.Join(request.Path, request.File.Filename))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
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
