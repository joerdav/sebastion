package sebastion

import (
	"context"
	"log"
)

func NewContext(ctx context.Context) Context {
	return Context{
		Logger: log.Default(),
		ctx:    &ctx,
	}
}

type Context struct {
	ctx    *context.Context
	Logger *log.Logger
}

func (c *Context) Context() context.Context {
	if c.ctx != nil {
		return *c.ctx
	}
	return context.Background()
}
