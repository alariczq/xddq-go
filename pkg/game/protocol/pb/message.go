package pb

import protoreflect "google.golang.org/protobuf/reflect/protoreflect"

type Message interface {
	protoreflect.ProtoMessage
	MsgId() int32
}
