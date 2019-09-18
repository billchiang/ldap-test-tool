package controllers

import (
	"github.com/astaxie/beego"
	"github.com/billchiang/ldap-test-tool/g"
	"github.com/billchiang/ldap-test-tool/models"
)

type SearchFilterController struct {
	beego.Controller
}

type SearchResult struct {
	Results []models.LDAP_RESULT `json:"results"`
	Success bool                 `json:"success"`
}

func (this *SearchFilterController) Get() {
	searchFilter := this.Ctx.Input.Param(":filter")
	results, err := models.Single_Search(g.Config().Ldap, searchFilter)
	if err != nil {
		var failedResult MsgResult
		failedResult.Msg = err.Error()
		this.Data["json"] = failedResult
	} else {
		var successResult SearchResult
		successResult.Success = true
		successResult.Results = results
		this.Data["json"] = successResult
	}
	this.ServeJSON()
}
