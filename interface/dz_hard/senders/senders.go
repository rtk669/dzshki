package senders

type SenderMethod interface {
	Send()
}

type SendModule struct {
	SenderMethod SenderMethod
	sendsSlice   []SenderMethod
}

func NewSendModule(senderMethod SenderMethod) *SendModule {
	return &SendModule{
		SenderMethod: senderMethod,
		sendsSlice:   make([]SenderMethod, 0),
	}
}

func (s *SendModule) Collect(x string, y string) {
	SenderMethod 
	s.sendsSlice = append(s.sendsSlice, )
}

func (s *SendModule) Send() {

	for _, y := range s.sendsSlice {
		s.SenderMethod.Send()
	}
}
