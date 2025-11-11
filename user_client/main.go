package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"thrift-rpc-demo/gen-go/user"
)

func main() {
	// 1. 创建传输层（连接服务端 Thrift 端口，如 9091）
	transport, err := thrift.NewTSocket("localhost:9090")
	if err != nil {
		panic(err)
	}
	defer transport.Close()
	transport.Open()

	// 2. 配置协议（与服务端一致，如二进制协议）
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	// 3. 创建客户端实例
	client := user.NewUserServiceClientFactory(transport, protocolFactory)

	// 4. 调用服务端登录接口
	req := &user.UserLoginRequest{
		UserName: "lqf",
		Password: "768170",
	}
	resp, err := client.Login(context.Background(), req)
	if err != nil {
		// 处理异常
		if ex, ok := err.(*user.LoginException); ok {
			fmt.Printf("登录失败: %s\n", ex.Message)
		} else {
			fmt.Printf("调用错误: %v\n", err)
		}
		return
	}

	// 5. 处理响应
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Printf("登录成功，用户信息( JSON 格式):\n%s\n", string(jsonData))

}
