package redis

import (
	"github.com/caioeverest/fed-its/internal/config"
	"github.com/caioeverest/fed-its/internal/logger"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	*redis.Client
	cfg *config.Config
	log *logger.Logger
}

func New(cfg *config.Config, log *logger.Logger) *Client {
	r := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	return &Client{r, cfg, log}
}
