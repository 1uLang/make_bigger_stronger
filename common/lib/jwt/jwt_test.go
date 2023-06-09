package jwt

import "testing"

func TestCreateJwt(t *testing.T) {
	jwt := NewJWT(
		"561528c8-2d58-46df-a075-516bef5b7f80",
		86400,
	)
	payload := make(Payload)
	payload["uid"] = "xx-aa-ww"
	payload["zz"] = 123456
	token, err := jwt.Token(payload)
	t.Log(token, err)
}
