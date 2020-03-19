// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package go_micro_service_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserService interface {
	Signup(ctx context.Context, in *ReqSignup, opts ...client.CallOption) (*Resp, error)
	Signin(ctx context.Context, in *ReqSignin, opts ...client.CallOption) (*Resp, error)
	UserInfo(ctx context.Context, in *ReqUserInfo, opts ...client.CallOption) (*Resp, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Signup(ctx context.Context, in *ReqSignup, opts ...client.CallOption) (*Resp, error) {
	req := c.c.NewRequest(c.name, "UserService.Signup", in)
	out := new(Resp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Signin(ctx context.Context, in *ReqSignin, opts ...client.CallOption) (*Resp, error) {
	req := c.c.NewRequest(c.name, "UserService.Signin", in)
	out := new(Resp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserInfo(ctx context.Context, in *ReqUserInfo, opts ...client.CallOption) (*Resp, error) {
	req := c.c.NewRequest(c.name, "UserService.UserInfo", in)
	out := new(Resp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Signup(context.Context, *ReqSignup, *Resp) error
	Signin(context.Context, *ReqSignin, *Resp) error
	UserInfo(context.Context, *ReqUserInfo, *Resp) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Signup(ctx context.Context, in *ReqSignup, out *Resp) error
		Signin(ctx context.Context, in *ReqSignin, out *Resp) error
		UserInfo(ctx context.Context, in *ReqUserInfo, out *Resp) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Signup(ctx context.Context, in *ReqSignup, out *Resp) error {
	return h.UserServiceHandler.Signup(ctx, in, out)
}

func (h *userServiceHandler) Signin(ctx context.Context, in *ReqSignin, out *Resp) error {
	return h.UserServiceHandler.Signin(ctx, in, out)
}

func (h *userServiceHandler) UserInfo(ctx context.Context, in *ReqUserInfo, out *Resp) error {
	return h.UserServiceHandler.UserInfo(ctx, in, out)
}
