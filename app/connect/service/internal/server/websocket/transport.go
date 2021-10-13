package websocket

import (
	"github.com/go-kratos/kratos/v2/transport"
)

type Transpore struct{}

func (Transpore) Kind() transport.Kind {
	panic("implement me")
}

func (Transpore) Endpoint() string {
	panic("implement me")
}

func (Transpore) Operation() string {
	panic("implement me")
}

func (Transpore) RequestHeader() transport.Header {
	panic("implement me")
}

func (Transpore) ReplyHeader() transport.Header {
	panic("implement me")
}
