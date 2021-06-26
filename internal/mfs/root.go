package mfs

import (
	"github.com/spf13/afero"
)

type RootDIR string

func RootFS(rootDir RootDIR) afero.Fs {
	return afero.NewBasePathFs(afero.NewOsFs(), string(rootDir))
}
