//+build !swipe

// Code generated by Swipe v1.20.1. DO NOT EDIT.

//go:generate swipe
package rest

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/swipe-io/swipe/fixtures/service"
	"github.com/swipe-io/swipe/fixtures/user"
	"io"
)

type EndpointSet struct {
	GetAllEndpoint     endpoint.Endpoint
	TestMethodEndpoint endpoint.Endpoint
	CreateEndpoint     endpoint.Endpoint
	DeleteEndpoint     endpoint.Endpoint
	GetEndpoint        endpoint.Endpoint
}

func MakeEndpointSet(s service.Interface) EndpointSet {
	return EndpointSet{
		TestMethodEndpoint: makeTestMethodEndpoint(s),
		CreateEndpoint:     makeCreateEndpoint(s),
		DeleteEndpoint:     makeDeleteEndpoint(s),
		GetEndpoint:        makeGetEndpoint(s),
		GetAllEndpoint:     makeGetAllEndpoint(s),
	}
}

type createRequestServiceInterface struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}
type createResponseServiceInterface struct {
}

func makeCreateEndpoint(s service.Interface) endpoint.Endpoint {
	w := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequestServiceInterface)
		err := s.Create(ctx, req.Name, req.Data)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	return w
}

type deleteRequestServiceInterface struct {
	Id uint `json:"id"`
}
type deleteResponseServiceInterface struct {
	A string `json:"a"`
	B string `json:"b"`
}

func makeDeleteEndpoint(s service.Interface) endpoint.Endpoint {
	w := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteRequestServiceInterface)
		a, b, err := s.Delete(ctx, req.Id)
		if err != nil {
			return nil, err
		}
		return deleteResponseServiceInterface{A: a, B: b}, nil
	}
	return w
}

type getRequestServiceInterface struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Fname string  `json:"fname"`
	Price float32 `json:"price"`
	N     int     `json:"n"`
}
type getResponseServiceInterface struct {
	Data user.User `json:"data"`
}

func makeGetEndpoint(s service.Interface) endpoint.Endpoint {
	w := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequestServiceInterface)
		result, err := s.Get(ctx, req.Id, req.Name, req.Fname, req.Price, req.N)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return w
}

func makeGetAllEndpoint(s service.Interface) endpoint.Endpoint {
	w := func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := s.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return w
}

type testMethodRequestServiceInterface struct {
	Data map[string]interface{} `json:"data"`
	Ss   interface{}            `json:"ss"`
}
type testMethodResponseServiceInterface struct {
	States map[string]map[int][]string `json:"states"`
}

func makeTestMethodEndpoint(s service.Interface) endpoint.Endpoint {
	w := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(testMethodRequestServiceInterface)
		result, err := s.TestMethod(req.Data, req.Ss)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return w
}

func TestMethodEndpointFactory(instance string) (endpoint.Endpoint, io.Closer, error) {
	s, err := NewClientRESTServiceInterface(instance)
	if err != nil {
		return nil, nil, err
	}
	return makeTestMethodEndpoint(s), nil, nil

}

func CreateEndpointFactory(instance string) (endpoint.Endpoint, io.Closer, error) {
	s, err := NewClientRESTServiceInterface(instance)
	if err != nil {
		return nil, nil, err
	}
	return makeCreateEndpoint(s), nil, nil

}

func DeleteEndpointFactory(instance string) (endpoint.Endpoint, io.Closer, error) {
	s, err := NewClientRESTServiceInterface(instance)
	if err != nil {
		return nil, nil, err
	}
	return makeDeleteEndpoint(s), nil, nil

}

func GetEndpointFactory(instance string) (endpoint.Endpoint, io.Closer, error) {
	s, err := NewClientRESTServiceInterface(instance)
	if err != nil {
		return nil, nil, err
	}
	return makeGetEndpoint(s), nil, nil

}

func GetAllEndpointFactory(instance string) (endpoint.Endpoint, io.Closer, error) {
	s, err := NewClientRESTServiceInterface(instance)
	if err != nil {
		return nil, nil, err
	}
	return makeGetAllEndpoint(s), nil, nil

}
