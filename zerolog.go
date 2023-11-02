package sensitive_logging

import (
	"encoding/json"
)

type RedactSensitiveMarshaller struct {
	sensitiveFields []string
}

func NewRedactSensitiveMarshaller(sensitiveFields []string) RedactSensitiveMarshaller {
	return RedactSensitiveMarshaller{
		sensitiveFields: sensitiveFields,
	}
}

func (m RedactSensitiveMarshaller) Marshal(v any) ([]byte, error) {
	v = redactedSensitiveData(m.sensitiveFields, v)
	return json.Marshal(v)
}
