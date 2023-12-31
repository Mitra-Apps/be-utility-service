// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/image/image.proto

package imageconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	image "github.com/Mitra-Apps/be-utility-service/domain/proto/image"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// ImageServiceName is the fully-qualified name of the ImageService service.
	ImageServiceName = "proto.ImageService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ImageServiceGetImagesProcedure is the fully-qualified name of the ImageService's GetImages RPC.
	ImageServiceGetImagesProcedure = "/proto.ImageService/GetImages"
)

// ImageServiceClient is a client for the proto.ImageService service.
type ImageServiceClient interface {
	GetImages(context.Context, *connect.Request[image.GetImagesRequest]) (*connect.Response[image.GetImagesResponse], error)
}

// NewImageServiceClient constructs a client for the proto.ImageService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewImageServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ImageServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &imageServiceClient{
		getImages: connect.NewClient[image.GetImagesRequest, image.GetImagesResponse](
			httpClient,
			baseURL+ImageServiceGetImagesProcedure,
			opts...,
		),
	}
}

// imageServiceClient implements ImageServiceClient.
type imageServiceClient struct {
	getImages *connect.Client[image.GetImagesRequest, image.GetImagesResponse]
}

// GetImages calls proto.ImageService.GetImages.
func (c *imageServiceClient) GetImages(ctx context.Context, req *connect.Request[image.GetImagesRequest]) (*connect.Response[image.GetImagesResponse], error) {
	return c.getImages.CallUnary(ctx, req)
}

// ImageServiceHandler is an implementation of the proto.ImageService service.
type ImageServiceHandler interface {
	GetImages(context.Context, *connect.Request[image.GetImagesRequest]) (*connect.Response[image.GetImagesResponse], error)
}

// NewImageServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewImageServiceHandler(svc ImageServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	imageServiceGetImagesHandler := connect.NewUnaryHandler(
		ImageServiceGetImagesProcedure,
		svc.GetImages,
		opts...,
	)
	return "/proto.ImageService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ImageServiceGetImagesProcedure:
			imageServiceGetImagesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedImageServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedImageServiceHandler struct{}

func (UnimplementedImageServiceHandler) GetImages(context.Context, *connect.Request[image.GetImagesRequest]) (*connect.Response[image.GetImagesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.ImageService.GetImages is not implemented"))
}
