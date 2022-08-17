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
	"github.com/gef"
	"github.com/gef/sms/Events"
	"github.com/gef/sms/Registry"
)

func Init(g *gef.Gef) {

	//!数据库自动维护
	dbm := gef.DbManager{}
	//维护表结构
	dbm.AutoTable(tables)
	//维护后台菜单
	dbm.AutoAdminRules(adminRules)

	//!注册后台页面
	g.SetAdminPages(Registry.AdminPages)
	//!注册监听事件
	g.SetEvent(Events.Listeners)
}
