package domain

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Country  string `json:"country"`
}
