package main

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	// 引入 Thrift 生成的接口和本地的服务实现
	"thrift-rpc-demo/gen-go/com/example"
	"thrift-rpc-demo/server/service" // 导入 service 包
)

func main() {
	multiplexedProcessor := thrift.NewTMultiplexedProcessor()
	// 1. 创建 TCP 传输层（监听 9090 端口）
	transport, err := thrift.NewTServerSocket(":9090")
	if err != nil {
		panic(err)
	}

	// 2. 定义协议（二进制协议，需与客户端一致）
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	// 3. 实例化服务实现（从 service 包中获取）
	handler := &service.ExampleServiceImpl{}

	// 4. 注册服务处理器（将实现与 Thrift 框架绑定）
	processor := example.NewExampleServiceProcessor(handler)
	multiplexedProcessor.RegisterProcessor("ExampleService", processor)

	// 5. 创建并启动服务器
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		thrift.NewTTransportFactory(),
		protocolFactory,
	)

	fmt.Println("Thrift server starting on :9090...")
	if err := server.Serve(); err != nil {
		panic(err)
	}
}
