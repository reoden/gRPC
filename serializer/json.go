package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// 假设 message 是 protoiface.MessageV1 类型的变量

func ProtobufToJSON(message proto.Message) (string, error) {
	var protoMessage proto.Message = message.(proto.Message)
	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}
	return marshaler.MarshalToString(protoMessage)
}
