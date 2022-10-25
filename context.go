package sebastion

import (
	"context"
	"log"
)

type Context struct {
	ctx    *context.Context
	Logger *log.Logger
}

func NewContext(ctx context.Context) Context {
	return Context{
		Logger: log.Default(),
		ctx:    &ctx,
	}
}

func (c *Context) Context() context.Context {
	if c.ctx != nil {
		return *c.ctx
	}
	return context.Background()
}
