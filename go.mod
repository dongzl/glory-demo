module github.com/glory-go/glory-demo

go 1.13

require (
	github.com/apache/dubbo-go-hessian2 v1.9.1
	github.com/glory-go/glory v0.0.0-20210608090500-4164c499f39c
	github.com/golang/protobuf v1.5.2
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/glory-go/glory => ../glory
