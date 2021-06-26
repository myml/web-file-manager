package main

import (
	"context"
	"embed"
	"flag"
	"io/fs"
	"os"

	"github.com/google/wire"
	"github.com/myml/web-file-manager/internal/controller"
	"github.com/myml/web-file-manager/internal/handle"
	"github.com/myml/web-file-manager/internal/mfs"
)

//go:embed ui/dist/web
var web embed.FS

func uiFS() (fs.FS, error) {
	return fs.Sub(web, "ui/dist/web")
}

var root string
var addr string

func rootDir() mfs.RootDIR {
	return mfs.RootDIR(root)
}

var Set = wire.NewSet(uiFS, rootDir, handle.Set, mfs.Set, controller.Set)

func main() {
	pwd, _ := os.Getwd()
	flag.StringVar(&root, "d", pwd, "root dir")
	flag.StringVar(&addr, "l", ":8080", "listen addr")
	flag.Parse()

	app, err := NewApp(context.Background())
	if err != nil {
		panic(err)
	}
	app.Run(addr)
}
