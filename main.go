package main

import (
	"context"
	"embed"
	"flag"
	"io/fs"
	"net/http"
	"os"

	"github.com/coreos/go-systemd/activation"
	"github.com/google/wire"
	"github.com/myml/web-file-manager/internal/controller"
	"github.com/myml/web-file-manager/internal/handle"
	"github.com/myml/web-file-manager/internal/mfs"
)

var Set = wire.NewSet(uiFS, rootDir, handle.Set, mfs.Set, controller.Set)

//go:embed ui/dist/web
var web embed.FS
var root string
var addr string

func main() {
	pwd, _ := os.Getwd()
	flag.StringVar(&root, "d", pwd, "root dir")
	flag.StringVar(&addr, "l", ":8080", "listen addr")
	flag.Parse()

	app, err := NewApp(context.Background())
	if err != nil {
		panic(err)
	}
	listeners, err := activation.Listeners()
	if err != nil {
		panic(err)
	}
	if len(listeners) != 1 {
		panic(app.Run(addr))
	}
	http.Serve(listeners[0], app)
}

func uiFS() (fs.FS, error) {
	return fs.Sub(web, "ui/dist/web")
}

func rootDir() mfs.RootDIR {
	return mfs.RootDIR(root)
}
