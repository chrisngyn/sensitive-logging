package sensitive_logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ExampleRedactSensitiveCore_hide_sensitive_field() {
	logger := zap.NewExample(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return NewRedactSensitiveCore(core, []string{"password"})
	}))
	defer logger.Sync()

	logger.Info("This is a log message", zap.String("password", "secret"))
	// Output:
	// {"level":"info","msg":"This is a log message","password":"*****"}
}

func ExampleRedactSensitiveCore_hide_nested_sensitive_field() {
	logger := zap.NewExample(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return NewRedactSensitiveCore(core, []string{"password"})
	}))
	defer logger.Sync()

	type User struct {
		Password string `json:"password"`
		UserName string `json:"user_name"`
	}
	logger.Info("This is a log message", zap.Any("user", User{Password: "secret", UserName: "user1"}))
	// Output:
	// {"level":"info","msg":"This is a log message","user":{"password":"*****","user_name":"user1"}}
}
