package client

import "context"

type Client struct {
	config Config
	ctx    context.Context
}

func New(cnfg Config) *Client {
	return &Client{
		config: cnfg,
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
