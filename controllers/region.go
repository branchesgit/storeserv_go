package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"storeserv/models"
)

type RegionController struct {
	beego.Controller
}

// 获取配置的地区信息；
// /region/get
func (r *RegionController) Get() {
	o := orm.NewOrm()
	qs := o.QueryTable("region").OrderBy("Pid")

	var regions []*models.Region
	num, err := qs.All(&regions)

	if err != nil {
		fmt.Printf("error")
	}

	trs := new(models.Region).ChangeRegionsToTreeRegions(regions, int(num))
	r.Data["json"] = trs
	r.ServeJSON()
}