package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	ConnectTimeout int
	MaxRetry       int
	ReadTimeout    int
}

type Option func(*Config)

func WithConnectTimeout(timeout int) Option {
	return func(config *Config) {
		config.ConnectTimeout = timeout
	}
}

func WithMaxRetry(maxRetry int) Option {
	return func(config *Config) {
		config.MaxRetry = maxRetry
	}
}

func WithReadTimeout(timeout int) Option {
	return func(config *Config) {
		config.ReadTimeout = timeout
	}
}

type Client struct {
	Conf *Config
}

func NewClient(opts ...Option) *Client {
	// set default value
	config := &Config{
		ConnectTimeout: 1000,
		MaxRetry:       3,
		ReadTimeout:    2000,
	}
	// set specified options
	for _, opt := range opts {
		opt(config)
	}
	// build client
	client := &Client{
		Conf: config,
	}
	return client
}

func main() {
	cli := NewClient(
		WithConnectTimeout(500),
		WithMaxRetry(5),
	)
	b, _ := json.Marshal(cli)
	fmt.Println(string(b))
}
