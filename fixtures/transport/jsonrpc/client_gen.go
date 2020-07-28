//+build !swipe

// Code generated by Swipe v1.24.0. DO NOT EDIT.

//go:generate swipe
package jsonrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/l-vitaly/go-kit/transport/http/jsonrpc"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/swipe-io/swipe/fixtures/service"
	"github.com/swipe-io/swipe/fixtures/user"
	"net/url"
)

type SwipeClientOption func(*clientSwipe)

func SwipeGenericClientOptions(opt ...jsonrpc.ClientOption) (_ SwipeClientOption) {
	return func(c *clientSwipe) { c.genericClientOption = opt }
}

func SwipeGenericClientEndpointMiddlewares(opt ...endpoint.Middleware) (_ SwipeClientOption) {
	return func(c *clientSwipe) { c.genericEndpointMiddleware = opt }
}

func SwipeCreateClientOptions(opt ...jsonrpc.ClientOption) (_ SwipeClientOption) {
	return func(c *clientSwipe) { c.createClientOption = opt }
}

func SwipeCreateClientEndpointMiddlewares(opt ...endpoint.Middleware) (_ SwipeClientOption) {
	return func(c *clientSwipe) { c.createEndpointMiddleware = opt }
}

func SwipeDeleteClientOptions(opt ...jsonrpc.ClientOption) (_ SwipeClientOption) {
	return func(c *clientSwipe) { c.deleteClientOption = opt }
}

func SwipeDeleteClientEndpointMiddlewares(opt ...endpoint.Middleware) (_ SwipeClientOption) {
	return func(c *clientSwipe) { c.deleteEndpointMiddleware = opt }
}

func SwipeGetClientOptions(opt ...jsonrpc.ClientOption) (_ SwipeClientOption) {
	return func(c *clientSwipe) { c.getClientOption = opt }
}

func SwipeGetClientEndpointMiddlewares(opt ...endpoint.Middleware) (_ SwipeClientOption) {
	return func(c *clientSwipe) { c.getEndpointMiddleware = opt }
}

func SwipeGetAllClientOptions(opt ...jsonrpc.ClientOption) (_ SwipeClientOption) {
	return func(c *clientSwipe) { c.getAllClientOption = opt }
}

func SwipeGetAllClientEndpointMiddlewares(opt ...endpoint.Middleware) (_ SwipeClientOption) {
	return func(c *clientSwipe) { c.getAllEndpointMiddleware = opt }
}

func SwipeTestMethodClientOptions(opt ...jsonrpc.ClientOption) (_ SwipeClientOption) {
	return func(c *clientSwipe) { c.testMethodClientOption = opt }
}

func SwipeTestMethodClientEndpointMiddlewares(opt ...endpoint.Middleware) (_ SwipeClientOption) {
	return func(c *clientSwipe) { c.testMethodEndpointMiddleware = opt }
}

func SwipeTestMethod2ClientOptions(opt ...jsonrpc.ClientOption) (_ SwipeClientOption) {
	return func(c *clientSwipe) { c.testMethod2ClientOption = opt }
}

func SwipeTestMethod2ClientEndpointMiddlewares(opt ...endpoint.Middleware) (_ SwipeClientOption) {
	return func(c *clientSwipe) { c.testMethod2EndpointMiddleware = opt }
}

type clientSwipe struct {
	createEndpoint                endpoint.Endpoint
	createClientOption            []jsonrpc.ClientOption
	createEndpointMiddleware      []endpoint.Middleware
	deleteEndpoint                endpoint.Endpoint
	deleteClientOption            []jsonrpc.ClientOption
	deleteEndpointMiddleware      []endpoint.Middleware
	getEndpoint                   endpoint.Endpoint
	getClientOption               []jsonrpc.ClientOption
	getEndpointMiddleware         []endpoint.Middleware
	getAllEndpoint                endpoint.Endpoint
	getAllClientOption            []jsonrpc.ClientOption
	getAllEndpointMiddleware      []endpoint.Middleware
	testMethodEndpoint            endpoint.Endpoint
	testMethodClientOption        []jsonrpc.ClientOption
	testMethodEndpointMiddleware  []endpoint.Middleware
	testMethod2Endpoint           endpoint.Endpoint
	testMethod2ClientOption       []jsonrpc.ClientOption
	testMethod2EndpointMiddleware []endpoint.Middleware
	genericClientOption           []jsonrpc.ClientOption
	genericEndpointMiddleware     []endpoint.Middleware
}

func (c *clientSwipe) Create(ctx context.Context, name string, data []byte) (_ error) {
	_, err := c.createEndpoint(ctx, createRequestSwipe{Name: name, Data: data})
	if err != nil {
		return err
	}
	return nil
}

func (c *clientSwipe) Delete(ctx context.Context, id uint) (_ string, _ string, _ error) {
	resp, err := c.deleteEndpoint(ctx, deleteRequestSwipe{Id: id})
	if err != nil {
		return "", "", err
	}
	response := resp.(deleteResponseSwipe)
	return response.A, response.B, nil
}

func (c *clientSwipe) Get(ctx context.Context, id int, name string, fname string, price float32, n int, b int, c int) (_ user.User, _ error) {
	resp, err := c.getEndpoint(ctx, getRequestSwipe{Id: id, Name: name, Fname: fname, Price: price, N: n, B: b, C: c})
	if err != nil {
		return user.User{}, err
	}
	response := resp.(getResponseSwipe)
	return response.Data, nil
}

func (c *clientSwipe) GetAll(ctx context.Context) (_ []*user.User, _ error) {
	resp, err := c.getAllEndpoint(ctx, nil)
	if err != nil {
		return nil, err
	}
	response := resp.([]*user.User)
	return response, nil
}

func (c *clientSwipe) TestMethod(data map[string]interface{}, ss interface{}) (_ map[string]map[int][]string, _ error) {
	resp, err := c.testMethodEndpoint(context.Background(), testMethodRequestSwipe{Data: data, Ss: ss})
	if err != nil {
		return nil, err
	}
	response := resp.(testMethodResponseSwipe)
	return response.States, nil
}

func (c *clientSwipe) TestMethod2(ctx context.Context, ns string, utype string, user string, restype string, resource string, permission string) (_ error) {
	_, err := c.testMethod2Endpoint(ctx, testMethod2RequestSwipe{Ns: ns, Utype: utype, User: user, Restype: restype, Resource: resource, Permission: permission})
	if err != nil {
		return err
	}
	return nil
}

func NewClientJSONRPCSwipe(tgt string, opts ...SwipeClientOption) (service.Interface, error) {
	c := &clientSwipe{}
	for _, o := range opts {
		o(c)
	}
	u, err := url.Parse(tgt)
	if err != nil {
		return nil, err
	}
	c.createClientOption = append(
		c.createClientOption,
		jsonrpc.ClientRequestEncoder(func(_ context.Context, obj interface{}) (json.RawMessage, error) {
			req, ok := obj.(createRequestSwipe)
			if !ok {
				return nil, fmt.Errorf("couldn't assert request as createRequestSwipe, got %T", obj)
			}
			b, err := ffjson.Marshal(req)
			if err != nil {
				return nil, fmt.Errorf("couldn't marshal request %T: %s", obj, err)
			}
			return b, nil
		}),
		jsonrpc.ClientResponseDecoder(func(_ context.Context, response jsonrpc.Response) (interface{}, error) {
			if response.Error != nil {
				return nil, ErrorDecode(response.Error.Code)
			}
			return nil, nil
		}),
	)
	c.createEndpoint = jsonrpc.NewClient(
		u,
		"create",
		append(c.genericClientOption, c.createClientOption...)...,
	).Endpoint()
	c.createEndpoint = middlewareChain(append(c.genericEndpointMiddleware, c.createEndpointMiddleware...))(c.createEndpoint)
	c.deleteClientOption = append(
		c.deleteClientOption,
		jsonrpc.ClientRequestEncoder(func(_ context.Context, obj interface{}) (json.RawMessage, error) {
			req, ok := obj.(deleteRequestSwipe)
			if !ok {
				return nil, fmt.Errorf("couldn't assert request as deleteRequestSwipe, got %T", obj)
			}
			b, err := ffjson.Marshal(req)
			if err != nil {
				return nil, fmt.Errorf("couldn't marshal request %T: %s", obj, err)
			}
			return b, nil
		}),
		jsonrpc.ClientResponseDecoder(func(_ context.Context, response jsonrpc.Response) (interface{}, error) {
			if response.Error != nil {
				return nil, ErrorDecode(response.Error.Code)
			}
			var resp deleteResponseSwipe
			err := ffjson.Unmarshal(response.Result, &resp)
			if err != nil {
				return nil, fmt.Errorf("couldn't unmarshal body to deleteResponseSwipe: %s", err)
			}
			return resp, nil
		}),
	)
	c.deleteEndpoint = jsonrpc.NewClient(
		u,
		"delete",
		append(c.genericClientOption, c.deleteClientOption...)...,
	).Endpoint()
	c.deleteEndpoint = middlewareChain(append(c.genericEndpointMiddleware, c.deleteEndpointMiddleware...))(c.deleteEndpoint)
	c.getClientOption = append(
		c.getClientOption,
		jsonrpc.ClientRequestEncoder(func(_ context.Context, obj interface{}) (json.RawMessage, error) {
			req, ok := obj.(getRequestSwipe)
			if !ok {
				return nil, fmt.Errorf("couldn't assert request as getRequestSwipe, got %T", obj)
			}
			b, err := ffjson.Marshal(req)
			if err != nil {
				return nil, fmt.Errorf("couldn't marshal request %T: %s", obj, err)
			}
			return b, nil
		}),
		jsonrpc.ClientResponseDecoder(func(_ context.Context, response jsonrpc.Response) (interface{}, error) {
			if response.Error != nil {
				return nil, ErrorDecode(response.Error.Code)
			}
			var resp getResponseSwipe
			err := ffjson.Unmarshal(response.Result, &resp)
			if err != nil {
				return nil, fmt.Errorf("couldn't unmarshal body to getResponseSwipe: %s", err)
			}
			return resp, nil
		}),
	)
	c.getEndpoint = jsonrpc.NewClient(
		u,
		"get",
		append(c.genericClientOption, c.getClientOption...)...,
	).Endpoint()
	c.getEndpoint = middlewareChain(append(c.genericEndpointMiddleware, c.getEndpointMiddleware...))(c.getEndpoint)
	c.getAllClientOption = append(
		c.getAllClientOption,
		jsonrpc.ClientRequestEncoder(func(_ context.Context, obj interface{}) (json.RawMessage, error) {
			return nil, nil
		}),
		jsonrpc.ClientResponseDecoder(func(_ context.Context, response jsonrpc.Response) (interface{}, error) {
			if response.Error != nil {
				return nil, ErrorDecode(response.Error.Code)
			}
			var resp []*user.User
			err := ffjson.Unmarshal(response.Result, &resp)
			if err != nil {
				return nil, fmt.Errorf("couldn't unmarshal body to getAllResponseSwipe: %s", err)
			}
			return resp, nil
		}),
	)
	c.getAllEndpoint = jsonrpc.NewClient(
		u,
		"getAll",
		append(c.genericClientOption, c.getAllClientOption...)...,
	).Endpoint()
	c.getAllEndpoint = middlewareChain(append(c.genericEndpointMiddleware, c.getAllEndpointMiddleware...))(c.getAllEndpoint)
	c.testMethodClientOption = append(
		c.testMethodClientOption,
		jsonrpc.ClientRequestEncoder(func(_ context.Context, obj interface{}) (json.RawMessage, error) {
			req, ok := obj.(testMethodRequestSwipe)
			if !ok {
				return nil, fmt.Errorf("couldn't assert request as testMethodRequestSwipe, got %T", obj)
			}
			b, err := ffjson.Marshal(req)
			if err != nil {
				return nil, fmt.Errorf("couldn't marshal request %T: %s", obj, err)
			}
			return b, nil
		}),
		jsonrpc.ClientResponseDecoder(func(_ context.Context, response jsonrpc.Response) (interface{}, error) {
			if response.Error != nil {
				return nil, ErrorDecode(response.Error.Code)
			}
			var resp testMethodResponseSwipe
			err := ffjson.Unmarshal(response.Result, &resp)
			if err != nil {
				return nil, fmt.Errorf("couldn't unmarshal body to testMethodResponseSwipe: %s", err)
			}
			return resp, nil
		}),
	)
	c.testMethodEndpoint = jsonrpc.NewClient(
		u,
		"testMethod",
		append(c.genericClientOption, c.testMethodClientOption...)...,
	).Endpoint()
	c.testMethodEndpoint = middlewareChain(append(c.genericEndpointMiddleware, c.testMethodEndpointMiddleware...))(c.testMethodEndpoint)
	c.testMethod2ClientOption = append(
		c.testMethod2ClientOption,
		jsonrpc.ClientRequestEncoder(func(_ context.Context, obj interface{}) (json.RawMessage, error) {
			req, ok := obj.(testMethod2RequestSwipe)
			if !ok {
				return nil, fmt.Errorf("couldn't assert request as testMethod2RequestSwipe, got %T", obj)
			}
			b, err := ffjson.Marshal(req)
			if err != nil {
				return nil, fmt.Errorf("couldn't marshal request %T: %s", obj, err)
			}
			return b, nil
		}),
		jsonrpc.ClientResponseDecoder(func(_ context.Context, response jsonrpc.Response) (interface{}, error) {
			if response.Error != nil {
				return nil, ErrorDecode(response.Error.Code)
			}
			return nil, nil
		}),
	)
	c.testMethod2Endpoint = jsonrpc.NewClient(
		u,
		"testMethod2",
		append(c.genericClientOption, c.testMethod2ClientOption...)...,
	).Endpoint()
	c.testMethod2Endpoint = middlewareChain(append(c.genericEndpointMiddleware, c.testMethod2EndpointMiddleware...))(c.testMethod2Endpoint)
	return c, nil
}
