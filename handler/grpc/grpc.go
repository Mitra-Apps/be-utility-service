package grpc

import (
	"context"

	pb "github.com/Mitra-Apps/be-utility-service/domain/proto/image"
	"github.com/Mitra-Apps/be-utility-service/service"
)

type GrpcRoute struct {
	service service.ServiceInterface
	pb.UnimplementedImageServiceServer
}

func New(service service.ServiceInterface) pb.ImageServiceServer {
	return &GrpcRoute{
		service: service,
	}
}

func (g *GrpcRoute) GetImages(ctx context.Context, req *pb.GetImagesRequest) (*pb.GetImagesResponse, error) {
	images, err := g.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	protoImages := []*pb.Image{}
	for _, img := range images {
		protoImages = append(protoImages, img.ToProto())
	}

	return &pb.GetImagesResponse{
		Images: protoImages,
	}, nil
}
