package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		Indent:          " ",
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}
	b, err := marshaler.Marshal(message)
	return string(b), err
}

func JSONToProtobuf(in []byte, message proto.Message) error {
	err := protojson.Unmarshal(in, message)
	return err
}
