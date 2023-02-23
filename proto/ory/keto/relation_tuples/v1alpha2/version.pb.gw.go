// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: ory/keto/relation_tuples/v1alpha2/version.proto

/*
Package rts is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package rts

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_VersionService_GetVersion_0(ctx context.Context, marshaler runtime.Marshaler, client VersionServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetVersionRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetVersion(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_VersionService_GetVersion_0(ctx context.Context, marshaler runtime.Marshaler, server VersionServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetVersionRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetVersion(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterVersionServiceHandlerServer registers the http handlers for service VersionService to "mux".
// UnaryRPC     :call VersionServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterVersionServiceHandlerFromEndpoint instead.
func RegisterVersionServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server VersionServiceServer) error {

	mux.Handle("GET", pattern_VersionService_GetVersion_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/ory.keto.relation_tuples.v1alpha2.VersionService/GetVersion", runtime.WithHTTPPathPattern("/version"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_VersionService_GetVersion_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VersionService_GetVersion_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterVersionServiceHandlerFromEndpoint is same as RegisterVersionServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterVersionServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterVersionServiceHandler(ctx, mux, conn)
}

// RegisterVersionServiceHandler registers the http handlers for service VersionService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterVersionServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterVersionServiceHandlerClient(ctx, mux, NewVersionServiceClient(conn))
}

// RegisterVersionServiceHandlerClient registers the http handlers for service VersionService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "VersionServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "VersionServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "VersionServiceClient" to call the correct interceptors.
func RegisterVersionServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client VersionServiceClient) error {

	mux.Handle("GET", pattern_VersionService_GetVersion_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/ory.keto.relation_tuples.v1alpha2.VersionService/GetVersion", runtime.WithHTTPPathPattern("/version"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_VersionService_GetVersion_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VersionService_GetVersion_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_VersionService_GetVersion_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"version"}, ""))
)

var (
	forward_VersionService_GetVersion_0 = runtime.ForwardResponseMessage
)