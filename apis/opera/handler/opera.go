package handler

import (
	"context"
	"encoding/json"
	log "github.com/micro/go-micro/v2/logger"

	"opera/client"
	"github.com/micro/go-micro/v2/errors"
	api "github.com/micro/go-micro/v2/api/proto"
	opera "path/to/service/proto/opera"
)

type Opera struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Opera.Call is called by the API as /opera/call with post body {"name": "foo"}
func (e *Opera) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Received Opera.Call request")

	// extract the client from the context
	operaClient, ok := client.OperaFromContext(ctx)
	if !ok {
		return errors.InternalServerError("fuckopera.api.opera.opera.call", "opera client not found")
	}

	// make request
	response, err := operaClient.Call(ctx, &opera.Request{
		Name: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("fuckopera.api.opera.opera.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
