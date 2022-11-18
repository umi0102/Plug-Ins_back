package users

type LoginRequest struct {
	Phone    int    `json:"phone"`
	Password string `json:"password"`
}

type LoginByCodeRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
