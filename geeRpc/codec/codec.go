package codec

import "io"

type Header struct {
	ServiceMethod string //format "Service,Method" 服务名和方法名，通常与 Go 语言中的结构体和方法相映射。
	Seq           uint64 //sequence number chosen by client/请求的序号，也可以认为是某个请求的 ID，用来区分不同的请求。
	Error         string //错误信息
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(closer io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implemented
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
