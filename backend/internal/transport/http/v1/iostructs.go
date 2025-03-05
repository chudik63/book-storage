package v1

type signUpInput struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type response struct {
	Message string `json:"message"`
}
