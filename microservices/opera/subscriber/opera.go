package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	opera "opera/proto/opera"
)

type Opera struct{}

func (e *Opera) Handle(ctx context.Context, msg *opera.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *opera.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
