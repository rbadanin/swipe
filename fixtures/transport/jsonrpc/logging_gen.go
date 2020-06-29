//+build !swipe

// Code generated by Swipe v1.20.1. DO NOT EDIT.

//go:generate swipe
package jsonrpc

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/swipe-io/swipe/fixtures/service"
	"github.com/swipe-io/swipe/fixtures/user"
	"time"
)

type loggingMiddlewareServiceInterface struct {
	next   service.Interface
	logger log.Logger
}

func (s *loggingMiddlewareServiceInterface) GetAll(ctx context.Context) (result []*user.User, err error) {
	defer func(now time.Time) {
		s.logger.Log("method", "GetAll", "took", time.Since(now), "result", len(result), "err", err)
	}(time.Now())
	return s.next.GetAll(ctx)
}

func (s *loggingMiddlewareServiceInterface) TestMethod(data map[string]interface{}, ss interface{}) (states map[string]map[int][]string, err error) {
	defer func(now time.Time) {
		s.logger.Log("method", "TestMethod", "took", time.Since(now), "data", len(data), "ss", ss, "states", len(states), "err", err)
	}(time.Now())
	return s.next.TestMethod(data, ss)
}

func (s *loggingMiddlewareServiceInterface) Create(ctx context.Context, name string, data []byte) (err error) {
	defer func(now time.Time) {
		s.logger.Log("method", "Create", "took", time.Since(now), "name", name, "data", len(data), "err", err)
	}(time.Now())
	return s.next.Create(ctx, name, data)
}

func (s *loggingMiddlewareServiceInterface) Delete(ctx context.Context, id uint) (a string, b string, err error) {
	defer func(now time.Time) {
		s.logger.Log("method", "Delete", "took", time.Since(now), "id", id, "a", a, "b", b, "err", err)
	}(time.Now())
	return s.next.Delete(ctx, id)
}

func (s *loggingMiddlewareServiceInterface) Get(ctx context.Context, id int, name string, fname string, price float32, n int) (data user.User, err error) {
	defer func(now time.Time) {
		s.logger.Log("method", "Get", "took", time.Since(now), "id", id, "name", name, "fname", fname, "price", price, "n", n, "data", data, "err", err)
	}(time.Now())
	return s.next.Get(ctx, id, name, fname, price, n)
}

func NewLoggingMiddlewareServiceInterface(s service.Interface, logger log.Logger) service.Interface {
	return &loggingMiddlewareServiceInterface{next: s, logger: logger}
}
