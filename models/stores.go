package models

import "github.com/astaxie/beego/orm"

type Stores struct {
	Id int
	Name string
	Desc string
	RegionID int `orm:"column(regionID)"`
}


func init() {
	orm.RegisterModel(new(Stores))
}
