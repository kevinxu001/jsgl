package routers

import (
	"github.com/kevinxu001/jsgl/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
