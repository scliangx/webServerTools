package grpc


import (
	"fmt"
	pb "github.com/coderitx/webServerTools/proto" // 引入proto包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	clientAddress = "127.0.0.1:50052"
)

func Client() {
	// 连接
	conn, err := grpc.Dial(clientAddress, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer func() {
		_ = conn.Close()
	}()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		grpclog.Fatalln(err)
	}

	fmt.Println(res.Message)
}
