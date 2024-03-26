package handler

const HelloServiceName = "grpc_validate_test.proto.HelloService"

func (hs *NewHelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

type NewHelloService struct {
}
