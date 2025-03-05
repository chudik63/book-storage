package models

type SignUpInput struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Response struct {
	Message string `json:"message"`
}
