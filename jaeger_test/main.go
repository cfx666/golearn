package main

import (
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"time"
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

	//opentracing.SetGlobalTracer(tracer)

	span := tracer.StartSpan("main")
	time.Sleep(time.Second)
	span.Finish()

	/*	span = opentracing.StartSpan("FuncA", opentracing.ChildOf(span.Context()))
		time.Sleep(time.Millisecond * 500)
		span.Finish()*/

}
