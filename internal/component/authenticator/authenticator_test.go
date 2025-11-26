package authenticator_test

import (
	"fmt"
	"testing"
	"tron_robot/internal/component/authenticator"
)

func TestNewInstance(t *testing.T) {
	ins := authenticator.NewInstance(&authenticator.Config{
		Issuer:       "backstage",
		QrcodeWidth:  200,
		QrcodeHeight: 200,
	})

	salt := "abc"

	generator := ins.Generate("admin", salt)

	url, err := generator.Url()
	if err != nil {
		t.Fatal(err)
	}

	secret, err := generator.Secret()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(url)
	fmt.Println(secret)
}
