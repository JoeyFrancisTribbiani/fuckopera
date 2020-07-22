package client

import (
	"context"

	opera "github.com/joeyfrancistribbiani/fuckopera/microservices/opera/proto/opera"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
)

type operaKey struct{}

// FromContext retrieves the client from the Context
func OperaFromContext(ctx context.Context) (opera.OperaService, bool) {
	c, ok := ctx.Value(operaKey{}).(opera.OperaService)
	return c, ok
}

// Client returns a wrapper for the OperaClient
func OperaWrapper(service micro.Service) server.HandlerWrapper {
	client := opera.NewOperaService("fuckopera.service.opera", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, operaKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
