package users

type LoginRequest struct {
	Phone        string `json:"phone" binding:"max=14"`
	Password     string `json:"password" binding:"max=20"`
	Name         string `json:"name" binding:"max=50"`
	Code         string `json:"code" binding:"max=10"`
	IdentityType string `json:"identity_type" binding:"max=12" `
}
