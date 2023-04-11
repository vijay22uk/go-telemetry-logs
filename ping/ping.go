package ping

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type ping struct {
	logger log.Logger
}

func NewPingService(logger log.Logger) Service {
	return &ping{
		logger: logger,
	}
}

//	func (ping *ping) Ping(ctx context.Context, name string) string {
//		logger := log.With(ping.logger, "method", "Ping")
//		logger.Log("ping", "from", name)
//		return "pong"
//	}
func (ping *ping) SayHello(ctx context.Context, name string) (string, string, error) {
	tr := otel.Tracer("SayHello")
	ctx, span := tr.Start(ctx, "SayingHello")
	span.SetAttributes(attribute.Key("name").String(name))
	defer span.End()
	logger := log.With(ping.logger, "method", "SayHello")
	logger.Log("from", name)
	uuid := generateUniqueUUID(ctx)
	return uuid.String(), fmt.Sprintf("hello, %s", name), nil
}

func generateUniqueUUID(ctx context.Context) uuid.UUID {
	tr := otel.Tracer("SayHello")
	_, span := tr.Start(ctx, "UUID")
	defer span.End()
	uuid, _ := uuid.NewV4()
	return uuid
}
