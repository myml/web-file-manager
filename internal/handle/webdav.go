package handle

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myml/web-file-manager/internal/mfs"
	"golang.org/x/net/webdav"
)

type WebDavH gin.HandlerFunc

func Webdav(root mfs.RootDIR) WebDavH {
	dav := &webdav.Handler{
		Prefix:     "/dav/",
		FileSystem: webdav.Dir(root),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				log.Println(r, err)
			}
		},
	}
	return WebDavH(gin.WrapH(dav))
}
