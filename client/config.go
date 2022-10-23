package client

import "context"

type Config struct {
	Context context.Context
	// TODO: decide on a logging package.
	Logger string
}
