package main

import (
	"log"
	"net"

	"github.com/example/grpc_sample"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	log.Print("main start")

	// 9000番ポートでクライアントからのリクエストを受け付けるようにする
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Sample構造体のアドレスを渡すことで、クライアントからGetDataリクエストされると
	// GetDataメソッドが呼ばれるようになる
	grpc_sample.RegisterSampleServiceServer(grpcServer, &Sample{})

	// 以下でリッスンし続ける
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	log.Print("main end")
}

type Sample struct {
	name string
}

func (s *Sample) GetData(
	ctx context.Context,
	req *grpc_sample.GetDataRequest, // リクエストの型を更新します。
) (*grpc_sample.GetDataResponse, error) { // レスポンスの型を更新します。
	log.Print("Received num_type: ", req.NumType)

	// ここでuser_datasとnum_maxを設定します。
	userDatas := []*grpc_sample.UserData{
		// ここにユーザーデータを追加します。
		{UserId: "1", UserName: "ユーザー1"},
		{UserId: "2", UserName: "ユーザー2"},
		// ...
	}
	numMax := int32(len(userDatas)) // 例として、userDatasの長さをnum_maxとしています。

	// レスポンスを作成して返します。
	response := &grpc_sample.GetDataResponse{
		UserDatas: userDatas,
		NumMax:    numMax,
	}
	return response, nil
}
