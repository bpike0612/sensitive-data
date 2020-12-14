package logger

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

func sanitize(args ...interface{}) []interface{} {
	res := make([]interface{}, 0, len(args))

	for _, v := range args {
		switch a := v.(type) {
		case protoreflect.ProtoMessage:
			res = append(res, sanitizeProtoMessage(a))
		default:
			res = append(res, a)
		}
	}

	return res
}
