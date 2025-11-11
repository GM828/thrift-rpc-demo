package main

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	// 引入生成的代码包（路径根据实际模块名调整）
	"thrift-rpc-demo/gen-go/com/example"
)

func main() {
	// 1. 创建传输层（连接服务端）
	transport, err := thrift.NewTSocket("localhost:9090")
	if err != nil {
		panic(err)
	}

	// 2. 包装传输层（带缓冲的传输）
	transportFactory := thrift.NewTTransportFactory()
	trans, err := transportFactory.GetTransport(transport)
	if err != nil {
		panic(err)
	}
	defer trans.Close() // 确保关闭连接

	// 3. 打开连接
	if err := trans.Open(); err != nil {
		panic(err)
	}

	// 4. 定义协议（需与服务端一致，此处为二进制协议）
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := example.NewExampleServiceClientFactory(trans, protocolFactory)

	// 5. 调用服务方法
	err = client.SayHello(context.Background(), "Thrift")
	if err != nil {
		panic(err)
	}
	fmt.Println("Client: 调用 sayHello 成功")
}
