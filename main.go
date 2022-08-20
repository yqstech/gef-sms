/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2022/8/16 21:27
 */

package sms

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/yqstech/gef"
	"github.com/yqstech/gef-sms/Events"
	"github.com/yqstech/gef-sms/Registry"
	"net/http"
)

func Init(g *gef.Gef) {
	
	//!数据库自动维护
	dbm := gef.DbManager{}
	//维护表结构
	dbm.AutoTable(tables)
	//维护后台菜单
	dbm.AutoAdminRules(adminRules)
	//维护内置数据
	dbm.AutoInsideData(insideData)
	
	//!注册后台页面
	g.SetAdminPages(Registry.AdminPages)
	//!注册监听事件
	g.SetEvent(Events.Listeners)
	
	//!测试追加前台路由
	g.AddFrontRouters([]gef.FrontRouter{
		{
			Method:             "GET",
			Url:                "/sms_test",
			HandleOrFileSystem: Api(Hello),
		},
	})
}

func Api(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		next(w, r, ps)
	}
}
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "SMS安装成功")
}
