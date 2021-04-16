package services

import "github.com/bee-well/auth/mq"

// SignIn ensures that the user credentials are correct and that
// an authorization token is generated and sent back to the caller
func SignIn(email, password string) (string, error) {
	if email == "admin" && password == "password" {
		m := mq.NewMq()
		m.Publish("users", []byte("user authenticated"))
		return "TOKEN", nil
	}
	return "", nil
}
