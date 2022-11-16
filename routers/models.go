package routers

// Userinfo 用户信息

type Userinfo struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

type ProjectList struct {
	Name         string `json:"name"`
	Creator      string `json:"creator"`
	IdentityType string `json:"identityType"`
}

type ProjectName struct {
	Name string `json:"projectname"`
}

type JoinDeveloper struct {
	Projectname  string `json:"projectname"`
	Developer    string `json:"developer"`
	IdentityType string `json:"identityType"`
}
type LeaveDeveloper struct {
	Projectname string `json:"projectname"`
	Developer   string `json:"developer"`
}
