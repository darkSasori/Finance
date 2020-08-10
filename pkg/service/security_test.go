package service

import "testing"

func TestToken(t *testing.T) {
	token, err := encode("username")
	if err != nil {
		t.Error(err)
	}

	username, err := decode(token)
	if err != nil {
		t.Error(err)
	}

	if username != "username" {
		t.Error("Expected same username")
	}
}
