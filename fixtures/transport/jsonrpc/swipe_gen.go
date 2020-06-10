//+build !swipe

// Code generated by Swipe v1.11.4. DO NOT EDIT.

//go:generate swipe
package jsonrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	prometheus2 "github.com/go-kit/kit/metrics/prometheus"
	"github.com/gorilla/mux"
	"github.com/l-vitaly/go-kit/transport/http/jsonrpc"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/swipe-io/swipe/fixtures/service"
	"github.com/swipe-io/swipe/fixtures/user"
	"net/http"
	"net/url"
	"time"
)

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
		data, err := s.Get(ctx, req.Id, req.Name, req.Fname, req.Price, req.N)
		if err != nil {
			return nil, err
		}
		return getResponseServiceInterface{Data: data}, nil
	}
	return w
}

type getAllResponseServiceInterface []user.User

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

func makeTestMethodEndpoint(s service.Interface) endpoint.Endpoint {
	w := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(testMethodRequestServiceInterface)
		err := s.TestMethod(req.Data, req.Ss)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	return w
}

type loggingMiddlewareServiceInterface struct {
	next   service.Interface
	logger log.Logger
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

func (s *loggingMiddlewareServiceInterface) GetAll(ctx context.Context) (result []user.User, err error) {
	defer func(now time.Time) {
		s.logger.Log("method", "GetAll", "took", time.Since(now), "result", len(result), "err", err)
	}(time.Now())
	return s.next.GetAll(ctx)
}

func (s *loggingMiddlewareServiceInterface) TestMethod(data map[string]interface{}, ss interface{}) (err error) {
	defer func(now time.Time) {
		s.logger.Log("method", "TestMethod", "took", time.Since(now), "data", len(data), "ss", ss, "err", err)
	}(time.Now())
	return s.next.TestMethod(data, ss)
}

func (s *loggingMiddlewareServiceInterface) Create(ctx context.Context, name string, data []byte) (err error) {
	defer func(now time.Time) {
		s.logger.Log("method", "Create", "took", time.Since(now), "name", name, "data", len(data), "err", err)
	}(time.Now())
	return s.next.Create(ctx, name, data)
}

type instrumentingMiddlewareServiceInterface struct {
	next           service.Interface
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
}

func (s *instrumentingMiddlewareServiceInterface) TestMethod(data map[string]interface{}, ss interface{}) (_ error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "TestMethod").Add(1)
		s.requestLatency.With("method", "TestMethod").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.TestMethod(data, ss)
}

func (s *instrumentingMiddlewareServiceInterface) Create(ctx context.Context, name string, data []byte) (_ error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Create").Add(1)
		s.requestLatency.With("method", "Create").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.Create(ctx, name, data)
}

func (s *instrumentingMiddlewareServiceInterface) Delete(ctx context.Context, id uint) (a string, b string, _ error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Delete").Add(1)
		s.requestLatency.With("method", "Delete").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.Delete(ctx, id)
}

func (s *instrumentingMiddlewareServiceInterface) Get(ctx context.Context, id int, name string, fname string, price float32, n int) (data user.User, _ error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Get").Add(1)
		s.requestLatency.With("method", "Get").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.Get(ctx, id, name, fname, price, n)
}

func (s *instrumentingMiddlewareServiceInterface) GetAll(ctx context.Context) (_ []user.User, _ error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "GetAll").Add(1)
		s.requestLatency.With("method", "GetAll").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.GetAll(ctx)
}

func ErrorDecode(code int) (_ error) {
	switch code {
	default:
		return fmt.Errorf("error code %d", code)
	case -32001:
		return &service.ErrUnauthorized{}
	}
}

func middlewareChain(middlewares []endpoint.Middleware) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		if len(middlewares) == 0 {
			return next
		}
		outer := middlewares[0]
		others := middlewares[1:]
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return outer(next)
	}
}

type clientServiceInterface struct {
	getAllEndpoint               endpoint.Endpoint
	getAllClientOption           []jsonrpc.ClientOption
	getAllEndpointMiddleware     []endpoint.Middleware
	testMethodEndpoint           endpoint.Endpoint
	testMethodClientOption       []jsonrpc.ClientOption
	testMethodEndpointMiddleware []endpoint.Middleware
	createEndpoint               endpoint.Endpoint
	createClientOption           []jsonrpc.ClientOption
	createEndpointMiddleware     []endpoint.Middleware
	deleteEndpoint               endpoint.Endpoint
	deleteClientOption           []jsonrpc.ClientOption
	deleteEndpointMiddleware     []endpoint.Middleware
	getEndpoint                  endpoint.Endpoint
	getClientOption              []jsonrpc.ClientOption
	getEndpointMiddleware        []endpoint.Middleware
	genericClientOption          []jsonrpc.ClientOption
	genericEndpointMiddleware    []endpoint.Middleware
}

type clientServiceInterfaceOption func(*clientServiceInterface)

func ServiceInterfaceGenericClientOptions(opt ...jsonrpc.ClientOption) (_ clientServiceInterfaceOption) {
	return func(c *clientServiceInterface) { c.genericClientOption = opt }
}

func ServiceInterfaceGenericClientEndpointMiddlewares(opt ...endpoint.Middleware) (_ clientServiceInterfaceOption) {
	return func(c *clientServiceInterface) { c.genericEndpointMiddleware = opt }
}

func ServiceInterfaceTestMethodClientOptions(opt ...jsonrpc.ClientOption) (_ clientServiceInterfaceOption) {
	return func(c *clientServiceInterface) { c.testMethodClientOption = opt }
}

func ServiceInterfaceTestMethodClientEndpointMiddlewares(opt ...endpoint.Middleware) (_ clientServiceInterfaceOption) {
	return func(c *clientServiceInterface) { c.testMethodEndpointMiddleware = opt }
}

func ServiceInterfaceCreateClientOptions(opt ...jsonrpc.ClientOption) (_ clientServiceInterfaceOption) {
	return func(c *clientServiceInterface) { c.createClientOption = opt }
}

func ServiceInterfaceCreateClientEndpointMiddlewares(opt ...endpoint.Middleware) (_ clientServiceInterfaceOption) {
	return func(c *clientServiceInterface) { c.createEndpointMiddleware = opt }
}

func ServiceInterfaceDeleteClientOptions(opt ...jsonrpc.ClientOption) (_ clientServiceInterfaceOption) {
	return func(c *clientServiceInterface) { c.deleteClientOption = opt }
}

func ServiceInterfaceDeleteClientEndpointMiddlewares(opt ...endpoint.Middleware) (_ clientServiceInterfaceOption) {
	return func(c *clientServiceInterface) { c.deleteEndpointMiddleware = opt }
}

func ServiceInterfaceGetClientOptions(opt ...jsonrpc.ClientOption) (_ clientServiceInterfaceOption) {
	return func(c *clientServiceInterface) { c.getClientOption = opt }
}

func ServiceInterfaceGetClientEndpointMiddlewares(opt ...endpoint.Middleware) (_ clientServiceInterfaceOption) {
	return func(c *clientServiceInterface) { c.getEndpointMiddleware = opt }
}

func ServiceInterfaceGetAllClientOptions(opt ...jsonrpc.ClientOption) (_ clientServiceInterfaceOption) {
	return func(c *clientServiceInterface) { c.getAllClientOption = opt }
}

func ServiceInterfaceGetAllClientEndpointMiddlewares(opt ...endpoint.Middleware) (_ clientServiceInterfaceOption) {
	return func(c *clientServiceInterface) { c.getAllEndpointMiddleware = opt }
}

func (c *clientServiceInterface) Get(ctx context.Context, id int, name string, fname string, price float32, n int) (_ user.User, _ error) {
	resp, err := c.getEndpoint(ctx, getRequestServiceInterface{Id: id, Name: name, Fname: fname, Price: price, N: n})
	if err != nil {
		return user.User{}, err
	}
	response := resp.(getResponseServiceInterface)
	return response.Data, nil
}

func (c *clientServiceInterface) GetAll(ctx context.Context) (_ []user.User, _ error) {
	resp, err := c.getAllEndpoint(ctx, nil)
	if err != nil {
		return nil, err
	}
	response := resp.(getAllResponseServiceInterface)
	return response, nil
}

func (c *clientServiceInterface) TestMethod(data map[string]interface{}, ss interface{}) (_ error) {
	_, err := c.testMethodEndpoint(context.Background(), testMethodRequestServiceInterface{Data: data, Ss: ss})
	if err != nil {
		return err
	}
	return nil
}

func (c *clientServiceInterface) Create(ctx context.Context, name string, data []byte) (_ error) {
	_, err := c.createEndpoint(ctx, createRequestServiceInterface{Name: name, Data: data})
	if err != nil {
		return err
	}
	return nil
}

func (c *clientServiceInterface) Delete(ctx context.Context, id uint) (_ string, _ string, _ error) {
	resp, err := c.deleteEndpoint(ctx, deleteRequestServiceInterface{Id: id})
	if err != nil {
		return "", "", err
	}
	response := resp.(deleteResponseServiceInterface)
	return response.A, response.B, nil
}

func NewClientJSONRPCServiceInterface(tgt string, opts ...clientServiceInterfaceOption) (service.Interface, error) {
	c := &clientServiceInterface{}
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
			req, ok := obj.(createRequestServiceInterface)
			if !ok {
				return nil, fmt.Errorf("couldn't assert request as createRequestServiceInterface, got %T", obj)
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
			req, ok := obj.(deleteRequestServiceInterface)
			if !ok {
				return nil, fmt.Errorf("couldn't assert request as deleteRequestServiceInterface, got %T", obj)
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
			var resp deleteResponseServiceInterface
			err := ffjson.Unmarshal(response.Result, &resp)
			if err != nil {
				return nil, fmt.Errorf("couldn't unmarshal body to deleteResponseServiceInterface: %s", err)
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
			req, ok := obj.(getRequestServiceInterface)
			if !ok {
				return nil, fmt.Errorf("couldn't assert request as getRequestServiceInterface, got %T", obj)
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
			var resp getResponseServiceInterface
			err := ffjson.Unmarshal(response.Result, &resp)
			if err != nil {
				return nil, fmt.Errorf("couldn't unmarshal body to getResponseServiceInterface: %s", err)
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
			var resp getAllResponseServiceInterface
			err := ffjson.Unmarshal(response.Result, &resp)
			if err != nil {
				return nil, fmt.Errorf("couldn't unmarshal body to getAllResponseServiceInterface: %s", err)
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
			req, ok := obj.(testMethodRequestServiceInterface)
			if !ok {
				return nil, fmt.Errorf("couldn't assert request as testMethodRequestServiceInterface, got %T", obj)
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
	c.testMethodEndpoint = jsonrpc.NewClient(
		u,
		"testMethod",
		append(c.genericClientOption, c.testMethodClientOption...)...,
	).Endpoint()
	c.testMethodEndpoint = middlewareChain(append(c.genericEndpointMiddleware, c.testMethodEndpointMiddleware...))(c.testMethodEndpoint)
	return c, nil
}

type serverServiceInterfaceOption func(*serverServiceInterfaceOpts)
type serverServiceInterfaceOpts struct {
	genericServerOption          []jsonrpc.ServerOption
	genericEndpointMiddleware    []endpoint.Middleware
	getAllServerOption           []jsonrpc.ServerOption
	getAllEndpointMiddleware     []endpoint.Middleware
	testMethodServerOption       []jsonrpc.ServerOption
	testMethodEndpointMiddleware []endpoint.Middleware
	createServerOption           []jsonrpc.ServerOption
	createEndpointMiddleware     []endpoint.Middleware
	deleteServerOption           []jsonrpc.ServerOption
	deleteEndpointMiddleware     []endpoint.Middleware
	getServerOption              []jsonrpc.ServerOption
	getEndpointMiddleware        []endpoint.Middleware
}

func ServiceInterfaceGenericServerOptions(v ...jsonrpc.ServerOption) (_ serverServiceInterfaceOption) {
	return func(o *serverServiceInterfaceOpts) { o.genericServerOption = v }
}

func ServiceInterfaceGenericServerEndpointMiddlewares(v ...endpoint.Middleware) (_ serverServiceInterfaceOption) {
	return func(o *serverServiceInterfaceOpts) { o.genericEndpointMiddleware = v }
}

func ServiceInterfaceTestMethodServerOptions(opt ...jsonrpc.ServerOption) (_ serverServiceInterfaceOption) {
	return func(c *serverServiceInterfaceOpts) { c.testMethodServerOption = opt }
}

func ServiceInterfaceTestMethodServerEndpointMiddlewares(opt ...endpoint.Middleware) (_ serverServiceInterfaceOption) {
	return func(c *serverServiceInterfaceOpts) { c.testMethodEndpointMiddleware = opt }
}

func ServiceInterfaceCreateServerOptions(opt ...jsonrpc.ServerOption) (_ serverServiceInterfaceOption) {
	return func(c *serverServiceInterfaceOpts) { c.createServerOption = opt }
}

func ServiceInterfaceCreateServerEndpointMiddlewares(opt ...endpoint.Middleware) (_ serverServiceInterfaceOption) {
	return func(c *serverServiceInterfaceOpts) { c.createEndpointMiddleware = opt }
}

func ServiceInterfaceDeleteServerOptions(opt ...jsonrpc.ServerOption) (_ serverServiceInterfaceOption) {
	return func(c *serverServiceInterfaceOpts) { c.deleteServerOption = opt }
}

func ServiceInterfaceDeleteServerEndpointMiddlewares(opt ...endpoint.Middleware) (_ serverServiceInterfaceOption) {
	return func(c *serverServiceInterfaceOpts) { c.deleteEndpointMiddleware = opt }
}

func ServiceInterfaceGetServerOptions(opt ...jsonrpc.ServerOption) (_ serverServiceInterfaceOption) {
	return func(c *serverServiceInterfaceOpts) { c.getServerOption = opt }
}

func ServiceInterfaceGetServerEndpointMiddlewares(opt ...endpoint.Middleware) (_ serverServiceInterfaceOption) {
	return func(c *serverServiceInterfaceOpts) { c.getEndpointMiddleware = opt }
}

func ServiceInterfaceGetAllServerOptions(opt ...jsonrpc.ServerOption) (_ serverServiceInterfaceOption) {
	return func(c *serverServiceInterfaceOpts) { c.getAllServerOption = opt }
}

func ServiceInterfaceGetAllServerEndpointMiddlewares(opt ...endpoint.Middleware) (_ serverServiceInterfaceOption) {
	return func(c *serverServiceInterfaceOpts) { c.getAllEndpointMiddleware = opt }
}

// HTTP JSONRPC Transport
func encodeResponseJSONRPCServiceInterface(_ context.Context, result interface{}) (json.RawMessage, error) {
	b, err := ffjson.Marshal(result)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func MakeHandlerJSONRPCServiceInterface(s service.Interface, logger log.Logger, opts ...serverServiceInterfaceOption) (http.Handler, error) {
	sopt := &serverServiceInterfaceOpts{}
	for _, o := range opts {
		o(sopt)
	}
	s = &loggingMiddlewareServiceInterface{next: s, logger: logger}
	s = &instrumentingMiddlewareServiceInterface{
		next: s,
		requestCount: prometheus2.NewCounterFrom(prometheus.CounterOpts{
			Namespace: "",
			Subsystem: "",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, []string{"method"}),
		requestLatency: prometheus2.NewSummaryFrom(prometheus.SummaryOpts{
			Namespace: "",
			Subsystem: "",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, []string{"method"}),
	}
	r := mux.NewRouter()
	handler := jsonrpc.NewServer(jsonrpc.EndpointCodecMap{
		"delete": jsonrpc.EndpointCodec{
			Endpoint: middlewareChain(append(sopt.genericEndpointMiddleware, sopt.deleteEndpointMiddleware...))(makeDeleteEndpoint(s)),
			Decode: func(_ context.Context, msg json.RawMessage) (interface{}, error) {
				var req deleteRequestServiceInterface
				err := ffjson.Unmarshal(msg, &req)
				if err != nil {
					return nil, fmt.Errorf("couldn't unmarshal body to deleteRequestServiceInterface: %s", err)
				}
				return req, nil
			},
			Encode: encodeResponseJSONRPCServiceInterface,
		},
		"get": jsonrpc.EndpointCodec{
			Endpoint: middlewareChain(append(sopt.genericEndpointMiddleware, sopt.getEndpointMiddleware...))(makeGetEndpoint(s)),
			Decode: func(_ context.Context, msg json.RawMessage) (interface{}, error) {
				var req getRequestServiceInterface
				err := ffjson.Unmarshal(msg, &req)
				if err != nil {
					return nil, fmt.Errorf("couldn't unmarshal body to getRequestServiceInterface: %s", err)
				}
				return req, nil
			},
			Encode: encodeResponseJSONRPCServiceInterface,
		},
		"getAll": jsonrpc.EndpointCodec{
			Endpoint: middlewareChain(append(sopt.genericEndpointMiddleware, sopt.getAllEndpointMiddleware...))(makeGetAllEndpoint(s)),
			Decode: func(_ context.Context, msg json.RawMessage) (interface{}, error) {
				return nil, nil
			},
			Encode: encodeResponseJSONRPCServiceInterface,
		},
		"testMethod": jsonrpc.EndpointCodec{
			Endpoint: middlewareChain(append(sopt.genericEndpointMiddleware, sopt.testMethodEndpointMiddleware...))(makeTestMethodEndpoint(s)),
			Decode: func(_ context.Context, msg json.RawMessage) (interface{}, error) {
				var req testMethodRequestServiceInterface
				err := ffjson.Unmarshal(msg, &req)
				if err != nil {
					return nil, fmt.Errorf("couldn't unmarshal body to testMethodRequestServiceInterface: %s", err)
				}
				return req, nil
			},
			Encode: encodeResponseJSONRPCServiceInterface,
		},
		"create": jsonrpc.EndpointCodec{
			Endpoint: middlewareChain(append(sopt.genericEndpointMiddleware, sopt.createEndpointMiddleware...))(makeCreateEndpoint(s)),
			Decode: func(_ context.Context, msg json.RawMessage) (interface{}, error) {
				var req createRequestServiceInterface
				err := ffjson.Unmarshal(msg, &req)
				if err != nil {
					return nil, fmt.Errorf("couldn't unmarshal body to createRequestServiceInterface: %s", err)
				}
				return req, nil
			},
			Encode: encodeResponseJSONRPCServiceInterface,
		},
	}, sopt.genericServerOption...)
	r.Methods("POST").Path("/rpc/{method}").Handler(handler)
	return r, nil
}
