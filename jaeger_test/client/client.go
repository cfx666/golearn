package main

import (
	"context"
	"fmt"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"learngo/grpc_interceptor_test/proto"
	"learngo/jaeger_test/otgrpc"
)

func main() {

	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{ // 采样器，对于大流量的服务，如果每个请求都采集，压力较大。可以设置采样率，减少数据量
			Type:  jaeger.SamplerTypeConst, // 采样类型，Const：固定采样率，Probabilistic：概率采样，RateLimiting：限速采样
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "192.168.137.129:6831",
		},
		ServiceName: "jaeger_test",
	}

	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		return
	}
	defer closer.Close()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)))
	if err != nil {
		panic("连接失败")
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "世界",
	})
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(r.Message)
}
