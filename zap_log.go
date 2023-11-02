package sensitive_logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// RedactSensitiveCore wraps an existing Core and redacts sensitive data.
type RedactSensitiveCore struct {
	zapcore.Core
	sensitiveFields []string
}

func NewRedactSensitiveCore(core zapcore.Core, sensitiveFields []string) zapcore.Core {
	return &RedactSensitiveCore{
		Core:            core,
		sensitiveFields: sensitiveFields,
	}
}

func (c *RedactSensitiveCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *RedactSensitiveCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	// Redact sensitive data before writing the log entry
	redactedFields := c.redactSensitiveData(fields)
	return c.Core.Write(entry, redactedFields)
}

func (c *RedactSensitiveCore) redactSensitiveData(fields []zapcore.Field) []zapcore.Field {
	var redactedFields []zapcore.Field

	for _, field := range fields {
		if contains(c.sensitiveFields, field.Key) {
			// Redact the sensitive data
			redactedFields = append(redactedFields, zap.String(field.Key, redactedValue))
			continue
		}

		if field.Type == zapcore.ReflectType {
			field.Interface = redactedSensitiveData(c.sensitiveFields, field.Interface)
		}

		redactedFields = append(redactedFields, field)
	}

	return redactedFields
}
