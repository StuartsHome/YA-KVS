package client

import (
	"context"

	"github.com/StuartsHome/YA-KVS/store"
)

type Client struct {
	config Config
	ctx    context.Context

	cache *store.StoreImpl
}

func New(cnfg Config) *Client {
	return &Client{
		config: cnfg,
		cache:  store.NewStore(),
	}
}

func newClient(cnfg *Config) *Client {
	if cnfg == nil {
		cnfg = &Config{}
	}

	client := &Client{
		config: Config{},
		ctx:    nil,
	}

	return client
}
