// Code generated by protoc-gen-lava. DO NOT EDIT.
// versions:
// - protoc-gen-lava v0.1.0
// - protoc         v3.17.3
// source: proto/hello/helloworld.proto

package hello

import (
	gin "github.com/gin-gonic/gin"
	grpcc "github.com/pubgo/lava/clients/grpcc"
	binding "github.com/pubgo/lava/pkg/binding"
	xgen "github.com/pubgo/lava/xgen"
	xerror "github.com/pubgo/xerror"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

func GetGreeterClient(srv string, opts ...func(cfg *grpcc.Cfg)) GreeterClient {
	return &greeterClient{grpcc.GetClient(srv, opts...)}
}
func init() {
	var mthList []xgen.GrpcRestHandler
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &HelloRequest{},
		Output:       &HelloReply{},
		Service:      "hello.Greeter",
		Name:         "SayHello",
		Method:       "GET",
		Path:         "/say/{name}",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})
	xgen.Add(RegisterGreeterServer, mthList)
	xgen.Add(RegisterGreeterHandler, nil)
	xgen.Add(RegisterGreeterGinServer, nil)
}
func RegisterGreeterGinServer(r gin.IRouter, server GreeterServer) {
	xerror.Assert(r == nil || server == nil, "router or server is nil")
	r.Handle("GET", "/say/{name}", func(ctx *gin.Context) {
		var req = new(HelloRequest)
		xerror.Panic(binding.MapFormByTag(req, ctx.Request.URL.Query(), "json"))
		var resp, err = server.SayHello(ctx, req)
		xerror.Panic(err)
		ctx.JSON(200, resp)
	})
}
