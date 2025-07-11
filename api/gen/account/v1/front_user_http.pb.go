// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             (unknown)
// source: account/v1/front_user.proto

package accountv1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationFrontLoginLogin = "/account.v1.FrontLogin/Login"
const OperationFrontLoginLogout = "/account.v1.FrontLogin/Logout"
const OperationFrontLoginRegister = "/account.v1.FrontLogin/Register"

type FrontLoginHTTPServer interface {
	Login(context.Context, *LoginRequest) (*LoginLoginReply, error)
	Logout(context.Context, *RegisterRequest) (*LoginLoginReply, error)
	Register(context.Context, *RegisterRequest) (*LoginLoginReply, error)
}

func RegisterFrontLoginHTTPServer(s *http.Server, srv FrontLoginHTTPServer) {
	r := s.Route("/")
	r.POST("/front/account/login/login", _FrontLogin_Login0_HTTP_Handler(srv))
	r.POST("/front/account/login/register", _FrontLogin_Register0_HTTP_Handler(srv))
	r.GET("/front/account/login/logout", _FrontLogin_Logout0_HTTP_Handler(srv))
}

func _FrontLogin_Login0_HTTP_Handler(srv FrontLoginHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFrontLoginLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginLoginReply)
		return ctx.Result(200, reply)
	}
}

func _FrontLogin_Register0_HTTP_Handler(srv FrontLoginHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFrontLoginRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginLoginReply)
		return ctx.Result(200, reply)
	}
}

func _FrontLogin_Logout0_HTTP_Handler(srv FrontLoginHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFrontLoginLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginLoginReply)
		return ctx.Result(200, reply)
	}
}

type FrontLoginHTTPClient interface {
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginLoginReply, err error)
	Logout(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *LoginLoginReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *LoginLoginReply, err error)
}

type FrontLoginHTTPClientImpl struct {
	cc *http.Client
}

func NewFrontLoginHTTPClient(client *http.Client) FrontLoginHTTPClient {
	return &FrontLoginHTTPClientImpl{client}
}

func (c *FrontLoginHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginLoginReply, error) {
	var out LoginLoginReply
	pattern := "/front/account/login/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFrontLoginLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FrontLoginHTTPClientImpl) Logout(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*LoginLoginReply, error) {
	var out LoginLoginReply
	pattern := "/front/account/login/logout"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationFrontLoginLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FrontLoginHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*LoginLoginReply, error) {
	var out LoginLoginReply
	pattern := "/front/account/login/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFrontLoginRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

const OperationUserBindMobile = "/account.v1.User/BindMobile"
const OperationUserCertificate = "/account.v1.User/Certificate"
const OperationUserChangePassword = "/account.v1.User/ChangePassword"
const OperationUserGetUser = "/account.v1.User/GetUser"
const OperationUserResetPassword = "/account.v1.User/ResetPassword"
const OperationUserUnBindMobile = "/account.v1.User/UnBindMobile"
const OperationUserUpdateUser = "/account.v1.User/UpdateUser"

type UserHTTPServer interface {
	// BindMobile 绑定手机号
	BindMobile(context.Context, *BindMobileRequest) (*LoginLoginReply, error)
	// Certificate 实名认证
	Certificate(context.Context, *CertificateRequest) (*emptypb.Empty, error)
	// ChangePassword 修改密码
	ChangePassword(context.Context, *ResetPasswordRequest) (*emptypb.Empty, error)
	GetUser(context.Context, *emptypb.Empty) (*GetUserReply, error)
	// ResetPassword 重设密码接口
	ResetPassword(context.Context, *ResetPasswordRequest) (*emptypb.Empty, error)
	// UnBindMobile 解绑手机号
	UnBindMobile(context.Context, *BindMobileRequest) (*LoginLoginReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*GetUserReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.GET("/front/account/user/info", _User_GetUser0_HTTP_Handler(srv))
	r.POST("/front/account/user/bindMobile", _User_UpdateUser0_HTTP_Handler(srv))
	r.POST("/front/account/user/bindMobile", _User_BindMobile0_HTTP_Handler(srv))
	r.POST("/front/account/user/unBindMobile", _User_UnBindMobile0_HTTP_Handler(srv))
	r.POST("/front/account/user/resetPassword", _User_ResetPassword0_HTTP_Handler(srv))
	r.POST("/front/account/user/changePassword", _User_ChangePassword0_HTTP_Handler(srv))
	r.POST("/front/account/user/saveCertificate", _User_Certificate0_HTTP_Handler(srv))
}

func _User_GetUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_BindMobile0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BindMobileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserBindMobile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BindMobile(ctx, req.(*BindMobileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginLoginReply)
		return ctx.Result(200, reply)
	}
}

func _User_UnBindMobile0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BindMobileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUnBindMobile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UnBindMobile(ctx, req.(*BindMobileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginLoginReply)
		return ctx.Result(200, reply)
	}
}

func _User_ResetPassword0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ResetPasswordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserResetPassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ResetPassword(ctx, req.(*ResetPasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _User_ChangePassword0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ResetPasswordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserChangePassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ChangePassword(ctx, req.(*ResetPasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _User_Certificate0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CertificateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCertificate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Certificate(ctx, req.(*CertificateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	BindMobile(ctx context.Context, req *BindMobileRequest, opts ...http.CallOption) (rsp *LoginLoginReply, err error)
	Certificate(ctx context.Context, req *CertificateRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	ChangePassword(ctx context.Context, req *ResetPasswordRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetUser(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GetUserReply, err error)
	ResetPassword(ctx context.Context, req *ResetPasswordRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	UnBindMobile(ctx context.Context, req *BindMobileRequest, opts ...http.CallOption) (rsp *LoginLoginReply, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *GetUserReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) BindMobile(ctx context.Context, in *BindMobileRequest, opts ...http.CallOption) (*LoginLoginReply, error) {
	var out LoginLoginReply
	pattern := "/front/account/user/bindMobile"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserBindMobile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) Certificate(ctx context.Context, in *CertificateRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/front/account/user/saveCertificate"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCertificate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) ChangePassword(ctx context.Context, in *ResetPasswordRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/front/account/user/changePassword"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserChangePassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) GetUser(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/front/account/user/info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/front/account/user/resetPassword"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserResetPassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) UnBindMobile(ctx context.Context, in *BindMobileRequest, opts ...http.CallOption) (*LoginLoginReply, error) {
	var out LoginLoginReply
	pattern := "/front/account/user/unBindMobile"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUnBindMobile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/front/account/user/bindMobile"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

const OperationUserDeliveryAddressCreateUserDeliveryAddress = "/account.v1.UserDeliveryAddress/CreateUserDeliveryAddress"
const OperationUserDeliveryAddressDeleteUserDeliveryAddress = "/account.v1.UserDeliveryAddress/DeleteUserDeliveryAddress"
const OperationUserDeliveryAddressGetUserDeliveryAddress = "/account.v1.UserDeliveryAddress/GetUserDeliveryAddress"
const OperationUserDeliveryAddressUpdateUserDeliveryAddress = "/account.v1.UserDeliveryAddress/UpdateUserDeliveryAddress"

type UserDeliveryAddressHTTPServer interface {
	// CreateUserDeliveryAddress 创建用户地址接口
	CreateUserDeliveryAddress(context.Context, *CreateUserDeliveryAddressReq) (*CreateUserDeliveryAddressRes, error)
	// DeleteUserDeliveryAddress 删除用户地址接口
	DeleteUserDeliveryAddress(context.Context, *DeleteUserDeliveryAddressReq) (*emptypb.Empty, error)
	// GetUserDeliveryAddress 获取用户地址接口
	GetUserDeliveryAddress(context.Context, *GetUserDeliveryAddressReq) (*CreateUserDeliveryAddressRes, error)
	// UpdateUserDeliveryAddress 更新用户地址接口
	UpdateUserDeliveryAddress(context.Context, *UpdateUserDeliveryAddressReq) (*UpdateUserDeliveryAddressRes, error)
}

func RegisterUserDeliveryAddressHTTPServer(s *http.Server, srv UserDeliveryAddressHTTPServer) {
	r := s.Route("/")
	r.POST("/front/account/userDeliveryAddress/add", _UserDeliveryAddress_CreateUserDeliveryAddress0_HTTP_Handler(srv))
	r.POST("/front/account/userDeliveryAddress/save", _UserDeliveryAddress_UpdateUserDeliveryAddress0_HTTP_Handler(srv))
	r.GET("/front/account/userDeliveryAddress/get", _UserDeliveryAddress_GetUserDeliveryAddress0_HTTP_Handler(srv))
	r.POST("/front/account/userDeliveryAddress/remove", _UserDeliveryAddress_DeleteUserDeliveryAddress0_HTTP_Handler(srv))
}

func _UserDeliveryAddress_CreateUserDeliveryAddress0_HTTP_Handler(srv UserDeliveryAddressHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserDeliveryAddressReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDeliveryAddressCreateUserDeliveryAddress)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUserDeliveryAddress(ctx, req.(*CreateUserDeliveryAddressReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserDeliveryAddressRes)
		return ctx.Result(200, reply)
	}
}

func _UserDeliveryAddress_UpdateUserDeliveryAddress0_HTTP_Handler(srv UserDeliveryAddressHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserDeliveryAddressReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDeliveryAddressUpdateUserDeliveryAddress)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserDeliveryAddress(ctx, req.(*UpdateUserDeliveryAddressReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserDeliveryAddressRes)
		return ctx.Result(200, reply)
	}
}

func _UserDeliveryAddress_GetUserDeliveryAddress0_HTTP_Handler(srv UserDeliveryAddressHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserDeliveryAddressReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDeliveryAddressGetUserDeliveryAddress)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserDeliveryAddress(ctx, req.(*GetUserDeliveryAddressReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserDeliveryAddressRes)
		return ctx.Result(200, reply)
	}
}

func _UserDeliveryAddress_DeleteUserDeliveryAddress0_HTTP_Handler(srv UserDeliveryAddressHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteUserDeliveryAddressReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDeliveryAddressDeleteUserDeliveryAddress)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUserDeliveryAddress(ctx, req.(*DeleteUserDeliveryAddressReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type UserDeliveryAddressHTTPClient interface {
	CreateUserDeliveryAddress(ctx context.Context, req *CreateUserDeliveryAddressReq, opts ...http.CallOption) (rsp *CreateUserDeliveryAddressRes, err error)
	DeleteUserDeliveryAddress(ctx context.Context, req *DeleteUserDeliveryAddressReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetUserDeliveryAddress(ctx context.Context, req *GetUserDeliveryAddressReq, opts ...http.CallOption) (rsp *CreateUserDeliveryAddressRes, err error)
	UpdateUserDeliveryAddress(ctx context.Context, req *UpdateUserDeliveryAddressReq, opts ...http.CallOption) (rsp *UpdateUserDeliveryAddressRes, err error)
}

type UserDeliveryAddressHTTPClientImpl struct {
	cc *http.Client
}

func NewUserDeliveryAddressHTTPClient(client *http.Client) UserDeliveryAddressHTTPClient {
	return &UserDeliveryAddressHTTPClientImpl{client}
}

func (c *UserDeliveryAddressHTTPClientImpl) CreateUserDeliveryAddress(ctx context.Context, in *CreateUserDeliveryAddressReq, opts ...http.CallOption) (*CreateUserDeliveryAddressRes, error) {
	var out CreateUserDeliveryAddressRes
	pattern := "/front/account/userDeliveryAddress/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserDeliveryAddressCreateUserDeliveryAddress))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserDeliveryAddressHTTPClientImpl) DeleteUserDeliveryAddress(ctx context.Context, in *DeleteUserDeliveryAddressReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/front/account/userDeliveryAddress/remove"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserDeliveryAddressDeleteUserDeliveryAddress))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserDeliveryAddressHTTPClientImpl) GetUserDeliveryAddress(ctx context.Context, in *GetUserDeliveryAddressReq, opts ...http.CallOption) (*CreateUserDeliveryAddressRes, error) {
	var out CreateUserDeliveryAddressRes
	pattern := "/front/account/userDeliveryAddress/get"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserDeliveryAddressGetUserDeliveryAddress))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserDeliveryAddressHTTPClientImpl) UpdateUserDeliveryAddress(ctx context.Context, in *UpdateUserDeliveryAddressReq, opts ...http.CallOption) (*UpdateUserDeliveryAddressRes, error) {
	var out UpdateUserDeliveryAddressRes
	pattern := "/front/account/userDeliveryAddress/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserDeliveryAddressUpdateUserDeliveryAddress))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
