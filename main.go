//Main. Program to process input stream and start SS and then return output stream.
//Author: Sriram Kaushik
//Date: 04/10/2018

package main

import (
	"github.com/astaxie/beego"
	"github.com/kaushiksriram100/ss-testing-framework/controllers"
)

func main() {

	//beego.Router("/", &controllers.ProcessController{})

	//Get the semaphore from controller and initialize it.

	controllers.Semaphore = make(chan struct{}, 1)

	beego.Router("/process", &controllers.ProcessController{})
	beego.Run()
}
