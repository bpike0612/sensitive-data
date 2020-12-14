package logger

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"sensitive-data/protocol"
)

func extractLogField(fd protoreflect.FieldDescriptor, hasPrefix bool) (string, bool) {
	// check if there are any options available for the field
	opts, ok := fd.Options().(*descriptorpb.FieldOptions)
	if !ok || opts == nil {
		return "", false
	}
	// check if protocol.FieldOptions specified for the field
	// basically extracting anything from [(protocol.options).logField = "something_here"]
	field, ok := proto.GetExtension(opts, protocol.E_Options).(*protocol.FieldOptions)
	if !ok || field == nil {
		return "", false
	}

	if field.LogField == "" && !hasPrefix {
		return "", false
	}

	return field.LogField, true
}
