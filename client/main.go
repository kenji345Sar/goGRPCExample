package main

import (
	"log"

	"github.com/example/grpc_sample" // 修正されたインポートパス
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	log.Print("Client is starting...")

	// サーバーに接続
	conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// gRPCクライアントを作成
	client := grpc_sample.NewSampleServiceClient(conn)

	// GetData RPCメソッドを呼び出す
	req := &grpc_sample.GetDataRequest{NumType: "1"} // 適切なリクエストパラメータを設定
	res, err := client.GetData(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call GetData: %v", err)
	}

	log.Printf("Received response: %v", res)
}
