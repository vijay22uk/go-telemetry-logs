package ping

import (
	"context"
)

type Service interface {
	// Ping(ctx context.Context, nane string) string
	SayHello(ctx context.Context, name string) (string, string, error)
}
