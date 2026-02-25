package ai

import "context"

type Client interface {
	Generate(ctx context.Context, promt string) (string, error)
}
