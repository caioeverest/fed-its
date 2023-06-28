package service

import "github.com/caioeverest/fedits/model"

type Response struct {
	StatusCode int
	Version    string
	Body       any
}

type Request struct {
	Method string
	Params []any
}

type Requester interface {
	Send(ctx, provider model.Provider, method model.Method, message Request) (Response, error)
}
