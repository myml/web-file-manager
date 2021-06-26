package handle

import "github.com/google/wire"

var Set = wire.NewSet(Move, Download, List, Upload, Mkdir, Delete)
