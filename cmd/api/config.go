package main

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type config struct {
	Port string `env:"PORT, default=8888"`
}

func getConfig(ctx context.Context) (*config, error) {
	var c config
	if err := envconfig.Process(ctx, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
