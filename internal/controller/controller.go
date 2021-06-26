package controller

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(NewEngine, wire.Struct(new(API), "*"))
