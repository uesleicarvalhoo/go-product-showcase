package middleware

import (
	"context"
	"time"
)

type Storage interface {
	HSet(ctx context.Context, key string, values ...any) error
	Set(ctx context.Context, key, value string, expiration time.Duration) error
	HSetExp(ctx context.Context, key string, expiration time.Duration, values ...any) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}
