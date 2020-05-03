package routers

import (
	"github.com/astaxie/beego"
	"movie/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
    beego.Include(&controllers.UserController{})
}
