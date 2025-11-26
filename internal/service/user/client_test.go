package user_test

import (
	"context"
	"testing"
	usersvc "tron_robot/internal/service/user"
	userpb "tron_robot/internal/service/user/pb"
	"tron_robot/internal/utils/xcrypt"
	"xbase/log"
	"xbase/registry/consul"
	"xbase/transport/rpcx"

	"google.golang.org/api/idtoken"
)

var transporter = rpcx.NewTransporter(
	rpcx.WithClientDiscovery(consul.NewRegistry()),
)

func TestClient_Login(t *testing.T) {
	client, err := usersvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.Login(context.Background(), &userpb.LoginArgs{
		Email:    "account",
		Password: "password",
		ClientIP: "127.0.0.1",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("UID: %d", reply.UID)
}

func TestClient_ModifyNickname(t *testing.T) {
	client, err := usersvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.ModifyNickname(context.Background(), &userpb.ModifyNicknameArgs{
		UID:      1,
		Nickname: "nickname",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestClient_ModifyAvatar(t *testing.T) {
	client, err := usersvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.ModifyAvatar(context.Background(), &userpb.ModifyAvatarArgs{
		UID:    1,
		Avatar: "/upload/user/avatar.jpg",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestClient_ModifyPassword(t *testing.T) {
	client, err := usersvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.ModifyPassword(context.Background(), &userpb.ModifyPasswordArgs{
		UID:         1,
		OldPassword: "123456",
		NewPassword: "123456789",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestClient_GoogleLogin(t *testing.T) {
	idToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjU1YzE4OGE4MzU0NmZjMTg4ZTUxNTc2YmE3MjgzNmUwNjAwZThiNzMiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJuYmYiOjE3MDg0OTgyMzAsImF1ZCI6Ijc1NjExODg1ODc4NC0zaGhlZWNlcTc4NnZwcTlsYTIyZWY3cHJxcWc0dDViZi5hcHBzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsInN1YiI6IjEwNzE3MjkyODczNzAyNzIyOTYzNSIsImVtYWlsIjoieXVlYmFuZnV4aWFvQGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJhenAiOiI3NTYxMTg4NTg3ODQtM2hoZWVjZXE3ODZ2cHE5bGEyMmVmN3BycXFnNHQ1YmYuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJuYW1lIjoiZnV4aWFvIHl1ZWJhbiIsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vYS9BQ2c4b2NJeWh5d1RkdXFOWnpwRUJLSW1WNk1UMlh3eEFKZVVKeXA5Y0RVNTZNRjVlUT1zOTYtYyIsImdpdmVuX25hbWUiOiJmdXhpYW8iLCJmYW1pbHlfbmFtZSI6Inl1ZWJhbiIsImlhdCI6MTcwODQ5ODUzMCwiZXhwIjoxNzA4NTAyMTMwLCJqdGkiOiJjODdhNWQ3N2Y1MGU4MWM2MGI4MGU1YTY5OWMwZDEwZjk3ZGVmYjgyIn0.MCDPMTXww2U7lzWvO8E2ynm96TjfH5X_0-gMZGq0tcLBDFGD1zzzMyF0Z9eJukwhHj29tgJN8j4k45WsGzGpWguGiDe4QGEtlCc7GeU_4pvnJvoa1eQ2GAEDw1yD6M3V17d-1YqrEiawbDFETcpoYf1q0XpqClupDB0w7dG88NADsjokYHIbyS5xCXTsGMzsdidZmYE-fbJzYep1fp4f17PpPKY7O7JUHUQErhpVb33X19hAuIFchaiAuIaiPK5RKNc9WnA6GbSQj0XNBHBB-dY9GCXYUg6evhmWOEGw4JmSvYYOZGDJdYCQyg1TCs9aEEm8ONyc4ozVlbhx3E_t1A"
	audience := "756118858784-3hheeceq786vpq9la22ef7prqqg4t5bf.apps.googleusercontent.com"

	payload, err := idtoken.Validate(context.Background(), idToken, audience)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(payload)
}

func TestClient_ComparePasswd(t *testing.T) {
	ok, err := xcrypt.Compare("$2a$10$0mKbNvWb51uqeUucZN5cJeVEEW0OdqAvC9gGTAQljb..xvSZAgOnW",
		"123456789", "NBwAoVcr")
	if err != nil {
		t.Log(err)
	}
	log.Warnf("%v", ok)
}
