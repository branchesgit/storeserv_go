package models

import "github.com/astaxie/beego/orm"

//type User struct {
//	Id int
//	SystemID int
//	UserName string `json: "userName"`
//	Password string	`json: "password"`
//}

type LoginUser struct {
	StoreID  int    `json: "storeID"`
	UserName string `json: "userName"`
	Password string `json: "password"`
}

func init() {
	orm.RegisterModel(new(Employees))
}

type Employees struct {
	Id        int
	SystemID  string `orm: "column(systemId)"`
	Psd       string
	Name      string
	Gender    int
	Cellphone string
	StoreID   int `orm: "column(storeId)"`
}

func (e *Employees) TableName() string {
	return "employees"
}
