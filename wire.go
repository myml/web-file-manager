//+build wireinject

package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func NewApp(ctx context.Context) (*gin.Engine, error) {
	panic(wire.Build(Set))
}
