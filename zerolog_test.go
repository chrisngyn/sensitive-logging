package sensitive_logging

import (
	"os"

	"github.com/rs/zerolog"
)

func ExampleRedactSensitiveMarshaller_Marshal() {
	m := NewRedactSensitiveMarshaller([]string{"password"})

	zerolog.InterfaceMarshalFunc = m.Marshal
	logger := zerolog.New(os.Stdout)

	type User struct {
		Password string `json:"password"`
		UserName string `json:"user_name"`
	}

	logger.Info().Any("user", User{
		Password: "123456",
		UserName: "test-user",
	}).Msg("This is a log message")
	// Output:
	// {"level":"info","user":{"password":"*****","user_name":"test-user"},"message":"This is a log message"}
}
