package handler

const HelloServiceName = "handler.HelloService"

func (hs *NewHelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

type NewHelloService struct {
}
