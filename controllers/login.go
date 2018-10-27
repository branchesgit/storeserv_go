package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"storeserv/models"
	"storeserv/utils"
)

type LoginController struct {
	beego.Controller
}

// 处理用户登录，并验证登录逻辑的；
func (l *LoginController) Post() {
	u := models.LoginUser{}
	var err error

	if err = json.Unmarshal(l.Ctx.Input.RequestBody, &u); err == nil {
		utils.FileLogs.Error("receive user information: %s, %s", u.UserName, u.Password);
	}

	var employees []*models.Employees;
	o := orm.NewOrm();
	num, err := o.QueryTable(new(models.Employees)).All(&employees);

	if err != nil {
		utils.FileLogs.Error("can't find any employee with storeID %d", u.StoreID);
	}

	h := md5.New();
	flag := false;
	for i := 0; i < int(num); i++ {
		var em *models.Employees = employees[i]

		if (em.SystemID == u.UserName) {
			if (em.Psd == u.Password) {
				flag = true;
				utils.FileLogs.Info("login success %s", u.UserName);
			} else {
				// 先将数据库中的字符进行md5加码， 再比对；
				h.Write([]byte(em.Psd));
				md5Psd := fmt.Sprintf("%x", h.Sum(nil))

				if (md5Psd == u.Password) {
					flag = true
					utils.FileLogs.Info("login success %s", u.UserName);
				}
			}

			if (flag) {
				break
			}
		}
	}

	if (flag) {
		l.Data["json"] = "{\"status\": \"success\"}"
	} else {
		l.Data["json"] = "{\"status\": \"fail\"}"
	}

	l.ServeJSON()
}
