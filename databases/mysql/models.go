package mysql

type Userinfo struct {
	Username string
	Password string
}
type err error

type ProjectName struct {
	Name string `json:"projectname"`
}
type DeveloperArr []ProjectName

type DevelopBack struct {
	State bool
	Data  DeveloperArr
}
type projectlist struct {
	Nameid   int
	Name     string
	Creator  string
	Finished int
}

type userArr []projectlist

type projectBack struct {
	State bool
	Data  userArr
}

type User struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}
