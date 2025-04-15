package service

import "context"

type HelloService struct {
}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (s *HelloService) GetHello(ctx context.Context) (string, error) {
	if ctx.Err() != nil {
		return "", ctx.Err()
	}
	return "Hello World!", nil
}
