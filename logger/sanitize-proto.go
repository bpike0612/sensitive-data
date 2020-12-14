package logger

import "google.golang.org/protobuf/reflect/protoreflect"

func sanitizeProtoMessage(m protoreflect.ProtoMessage, parentPrefixes ...string) map[string]interface{} {
	result := make(map[string]interface{})
	reflect := m.ProtoReflect()
	reflect.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if !v.IsValid() {
			return true
		}

		if fd.IsList() || fd.IsMap() {
			return true
		}

		prefixes := parentPrefixes[:]
		logField, shouldLog := extractLogField(fd, len(prefixes) > 0)
		if shouldLog {
			prefixes = append(prefixes, logField)
		}

		// check if current field is a Message itself
		if vi, ok := v.Interface().(protoreflect.Message); ok {
			subRes := sanitizeProtoMessage(vi.Interface(), prefixes...)
			for k, v := range subRes {
				result[k] = v
			}
			return true
		}

		if !shouldLog {
			return true
		}

		result[logFieldName(prefixes)] = v.Interface()

		return true
	})

	return result
}
