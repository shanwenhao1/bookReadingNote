package helloProto

type HelloService struct{}

func (p *HelloService) Hello(request *String, reply *String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}
