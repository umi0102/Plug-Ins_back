package users

type LoginRequest struct {
	Phone    int    `json:"phone"`
	Password string `json:"password"`
}

type LoginByCodeRequest struct {
	Phone int `json:"phone"`
	Code  int `json:"code"`
}
