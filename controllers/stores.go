package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"storeserv/models"
	"storeserv/utils"
)

type StoresController struct {
	beego.Controller
}

func (s *StoresController) Get() {
	regionID, err := s.GetInt("regionID")

	if err != nil {
		utils.FileLogs.Error("receive parameter regionID is %d error", regionID)
	}

	o := orm.NewOrm()
	qs := o.QueryTable("stores").Filter("Regionid", regionID)

	var stores []*models.Stores
	num, err := qs.All(&stores)

	if err != nil {
		utils.FileLogs.Error("Error: query store by regionID %d is error~", regionID)
	}

	if num == 0 {
		utils.FileLogs.Info("query store by regionID %d is empty", regionID)
	}

	s.Data["json"] = &stores
	s.ServeJSON()
}
