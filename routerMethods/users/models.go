package users

type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type LoginByCodeRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
