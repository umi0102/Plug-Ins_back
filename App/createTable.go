package App

import (
	"Plug-Ins/databases/mysql"
)

type bug struct {
	BugId                  string `gorm:"primarykey;not null;type:varchar(255);comment:bugid"`
	BugProjectName         string `gorm:"not null;type:varchar(30);comment:项目名称"`
	BugProjectPartName     string `gorm:"not null;type:varchar(30);comment:项目bug模块名称"`
	BugName                string `gorm:"not null;type:varchar(30);comment:bug名称"`
	BugFinder              string `gorm:"not null;type:varchar(30);comment:bug发现者"`
	BugFindTime            int64  `gorm:"type:int;comment:bug发现时间"`
	BugAssign              string `gorm:"type:varchar(255);comment:bug指派"`
	BugLevel               string `gorm:"not null;type:varchar(255);comment:bug等级：紧急/常规/历史遗留"`
	BugComplete            int64  `gorm:"not null;type:int;comment:是否完成"`
	BugCompleteTime        int64  `gorm:"type:int;comment:完成时间"`
	BugProjectPartFinisher string `gorm:"type:varchar(50);comment:模块贡献者"`
}

type developer struct {
	DeveloperId           string `gorm:"primarykey;not null;type:varchar(255);comment:DeveloperId"`
	DeveloperProjectName  string `gorm:"not null;type:varchar(255);comment:Userid"`
	DeveloperDeveloper    string `gorm:"not null;type:varchar(255);comment:Userid"`
	DeveloperIdentityType string `gorm:"not null;type:varchar(255);comment:Userid"`
}

type project struct {
	ProjectsId       string `gorm:"primarykey;not null;type:varchar(255);comment:ProjectsId"`
	ProjectsName     string `gorm:"not null;type:varchar(255);comment:Userid"`
	ProjectsCreator  string `gorm:"not null;type:varchar(255);comment:Userid"`
	ProjectsFinished int64  `gorm:"type:int;comment:Userid"`
}

type userinfo struct {
	UserinfoId           string `gorm:"primarykey;not null;type:varchar(255);comment:Userid"`
	UserinfoPhone        string `gorm:"not null;type:varchar(20);comment:手机"`
	UserinfoPassword     string `gorm:"not null;type:varchar(25);comment:密码"`
	UserinfoIdentityType string `gorm:"type:varchar(10);comment:"`
}

type needs struct {
	NeedsId              string `gorm:"primarykey;not null;type:varchar(255);comment:NeedsId"`
	NeedsProjectName     string `gorm:"type:varchar(50);comment:项目名称"`
	NeedsPartName        string `gorm:"type:varchar(50);comment:模块名称"`
	NeedsNeedsAssign     string `gorm:"type:varchar(50);comment:需求指派"`
	NeedsPartPublishTime int64  `gorm:"type:int;comment:需求发布时间"`
	NeedsContent         string `gorm:"type:varchar(255);comment:需求内容"`
	NeedsComplete        int64  `gorm:"type:int;comment:是否完成"`
}

func Create() {
	mysql.CreateTableMysql(&userinfo{}, &project{}, &developer{}, &bug{}, &needs{})
}
