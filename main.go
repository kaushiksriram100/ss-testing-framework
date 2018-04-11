//Main entry. Program to process input stream and start SS and then return output stream.

package main

import (
	"github.com/astaxie/beego"
	"github.com/kaushiksriram100/ss-testing-framework/controllers"
)

func main() {

	//beego.Router("/", &controllers.ProcessController{})
	beego.Router("/process", &controllers.ProcessController{})
	beego.Run()
}
