package main
import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pt "../protocol"
)
const (
	post  = "127.0.0.1:18887"
)
func main() {
	// 客户端连接服务器
	conn,err:=grpc.Dial(post,grpc.WithInsecure())
	if err!=nil {
		fmt.Println("连接服务器失败",err)
	}

	defer conn.Close()

	//获得grpc句柄
	c:=pt.NewHelloServerClient(conn)

	//远程调用 SayHello接口
	r1, err := c.SayHello(context.Background(), &pt.HelloRequest{Name: "Scott"})
	if err != nil {
		fmt.Println("cloud not get Hello server ..", err)
		return
	}
	fmt.Println("HelloServer resp: ", r1.Message)
	//远程调用 GetHelloMsg接口
	r2, err := c.GetHelloMsg(context.Background(), &pt.HelloRequest{Name: "Scott"})
	if err != nil {
		fmt.Println("cloud not get hello msg ..", err)
		return
	}
	fmt.Println("HelloServer resp: ", r2.Msg)
}