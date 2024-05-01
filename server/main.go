package main

import (
	"log"
	"net"

	"github.com/example/grpc_sample" // 修正されたインポートパス
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Sample struct{}

func (s *Sample) GetData(ctx context.Context, req *grpc_sample.GetDataRequest) (*grpc_sample.GetDataResponse, error) {
	log.Print("Received request: ", req.NumType)
	userDatas := []*grpc_sample.UserData{
		{UserId: "1", UserName: "ユーザー1"},
		{UserId: "2", UserName: "ユーザー2"},
	}
	numMax := int32(len(userDatas))
	response := &grpc_sample.GetDataResponse{
		UserDatas: userDatas,
		NumMax:    numMax,
	}
	return response, nil
}

func main() {
	log.Print("Server is starting...")
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	grpc_sample.RegisterSampleServiceServer(srv, &Sample{})
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
