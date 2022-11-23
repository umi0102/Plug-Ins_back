package projects

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
type Userinfo struct {
	Usericon     string `json:"usericon"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	IdentityType string `json:"identity_type"`
}
