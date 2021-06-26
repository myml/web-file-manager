package controller

import "github.com/myml/web-file-manager/internal/handle"

type API struct {
	handle.MoveF
	handle.DownloadF
	handle.ListF
	handle.UploadF
	handle.MkdirF
	handle.DeleteF
}
