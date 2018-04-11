//package controllers
//Author: Sriram Kaushik
//Date: 04/10/2018

package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kaushiksriram100/ss-testing-framework/models"
)

type ProcessController struct {
	beego.Controller
}

func (this *ProcessController) Get() {
	this.Layout = "layout.tpl"
	this.TplName = "home.tpl"
	this.Render()
}

func (this *ProcessController) Post() {

	inputs := this.Input()

	//convert the inputtext to a byte array

	inputdata := &models.Data{}
	resultdata := &models.Data{}

	(*inputdata).RawData = []byte(inputs["inputtext"][0])

	//fmt.Println((*inputdata).RawData)

	//add a counting semaphore here as stream splitter can process only one file at a time. So process one by one. use a empty struct.

	err := inputdata.RunBySS(resultdata)

	if err != nil {
		this.Data["ERROR"] = err
	}

	this.Data["OUTPUT"] = string((*resultdata).RawData)

	this.Layout = "layout.tpl"
	this.TplName = "home.tpl"
	this.Render()
}
