// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: ids.proto

package ids

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for IDsInterface service

func NewIDsInterfaceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for IDsInterface service

type IDsInterfaceService interface {
	GetID(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseID, error)
	GetServID(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type iDsInterfaceService struct {
	c    client.Client
	name string
}

func NewIDsInterfaceService(name string, c client.Client) IDsInterfaceService {
	return &iDsInterfaceService{
		c:    c,
		name: name,
	}
}

func (c *iDsInterfaceService) GetID(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseID, error) {
	req := c.c.NewRequest(c.name, "IDsInterface.GetID", in)
	out := new(ResponseID)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iDsInterfaceService) GetServID(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IDsInterface.GetServID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IDsInterface service

type IDsInterfaceHandler interface {
	GetID(context.Context, *Request, *ResponseID) error
	GetServID(context.Context, *Request, *Response) error
}

func RegisterIDsInterfaceHandler(s server.Server, hdlr IDsInterfaceHandler, opts ...server.HandlerOption) error {
	type iDsInterface interface {
		GetID(ctx context.Context, in *Request, out *ResponseID) error
		GetServID(ctx context.Context, in *Request, out *Response) error
	}
	type IDsInterface struct {
		iDsInterface
	}
	h := &iDsInterfaceHandler{hdlr}
	return s.Handle(s.NewHandler(&IDsInterface{h}, opts...))
}

type iDsInterfaceHandler struct {
	IDsInterfaceHandler
}

func (h *iDsInterfaceHandler) GetID(ctx context.Context, in *Request, out *ResponseID) error {
	return h.IDsInterfaceHandler.GetID(ctx, in, out)
}

func (h *iDsInterfaceHandler) GetServID(ctx context.Context, in *Request, out *Response) error {
	return h.IDsInterfaceHandler.GetServID(ctx, in, out)
}