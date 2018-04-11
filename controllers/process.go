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

	err := inputdata.RunBySS(resultdata)

	if err != nil {
		this.Data["ERROR"] = err
	}

	this.Data["OUTPUT"] = string((*resultdata).RawData)

	this.Layout = "layout.tpl"
	this.TplName = "home.tpl"
	this.Render()
}
