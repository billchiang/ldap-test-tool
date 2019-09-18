package http

import (
	"github.com/astaxie/beego"
	"github.com/billchiang/ldap-test-tool/g"
)

func Start() {
	beego.BConfig.CopyRequestBody = true
	ConfigRouters()
	beego.Run(g.Config().Http.Listen)
}
