package users

type LoginRequest struct {
	Phone    string `json:"phone" binding:"required,min=10,max=14"`
	Password string `json:"password" binding:"max=20"`
	Name     string `json:"name" binding:"max=50"`
	Code     string `json:"code" binding:"max=10"`
}

type LoginByCodeRequest struct {
	Phone string `json:"phone" binding:"required,min=10,max=14"`
	Code  string `json:"code" binding:"required,min=1,max=10"`
}
