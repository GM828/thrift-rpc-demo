package service

import (
	"context"
	"fmt"
)

// ExampleServiceImpl 实现 Thrift 生成的 ExampleService 接口
type ExampleServiceImpl struct{}

// SayHello 实现接口中的 SayHello 方法
func (e *ExampleServiceImpl) SayHello(ctx context.Context, name string) error {
	fmt.Printf("Hello, %s!\n", name) // 业务逻辑：打印问候语
	return nil
}
