/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: AppSmsRecord
 * @Version: 1.0.0
 * @Date: 2022/3/8 9:43 下午
 */

package AdminHandles

import (
	"github.com/wonderivan/logger"
	"github.com/yqstech/gef/Handles/adminHandle"
	"github.com/yqstech/gef/boot/db"
	"github.com/yqstech/gef/builder"
	"github.com/yqstech/gef/util"
)

type AppSmsRecord struct {
	adminHandle.Base
}

// NodeBegin 开始
func (that AppSmsRecord) NodeBegin(pageBuilder *builder.PageBuilder) (error, int) {
	pageBuilder.SetTitle("短信记录")
	pageBuilder.SetPageName("短信记录")
	pageBuilder.SetTbName("tb_app_sms_record")
	return nil, 0
}

// NodeList 初始化列表
func (that AppSmsRecord) NodeList(pageBuilder *builder.PageBuilder) (error, int) {
	uniappTable := pageBuilder.GetHttpParams().ByName("_uniapp_table")
	if uniappTable != "" {
		//新增多应用设置
		uniappList, err := db.New().Table(uniappTable).
			Where("is_delete", 0).
			Where("status", 1).
			Order("index_num asc,id asc").
			Get()
		if err != nil {
			logger.Error(err.Error())
			return err, 500
		}
		if len(uniappList) > 0 {
			//!设置tabs列表和选中项
			validUrl := util.UrlScreenParam(pageBuilder.GetHttpRequest(), []string{}, false, true)
			//!设置系统默认的一项
			pageBuilder.PageTabAdd("系统默认", validUrl+"tab=0&uniapp_id=0")
			//!其他应用往后排
			for index, uniappInfo := range uniappList {
				pageBuilder.PageTabAdd(uniappInfo["app_name"].(string), validUrl+"tab="+util.Int2String(index+1)+"&uniapp_id="+util.Int642String(uniappInfo["id"].(int64)))
			}
			//获取第几页
			tabIndex := that.GetTabIndex(pageBuilder, "tab")
			//设置第几个tab选中
			pageBuilder.SetPageTabSelect(tabIndex)
		}
	}
	pageBuilder.ListColumnClear()
	pageBuilder.SetListOrder("id desc")
	pageBuilder.ListTopBtnsClear()
	pageBuilder.ListRightBtnsClear()
	//获取列表
	upstreamOptions := SmsUpstreamList()
	pageBuilder.ListColumnAdd("template_name", "短信模板", "text", nil)
	pageBuilder.ListColumnAdd("upstream_id", "短信通道", "array", upstreamOptions)
	pageBuilder.ListColumnAdd("tel", "手机号码", "text", nil)
	pageBuilder.ListColumnAdd("ip", "ip地址", "text", nil)
	pageBuilder.ListColumnAdd("content", "短信内容", "text", nil)
	pageBuilder.ListColumnAdd("msg", "短信凭据/错误信息", "text", nil)
	pageBuilder.ListColumnAdd("status", "状态", "array", []map[string]interface{}{
		{"value": "1", "name": "未发送"},
		{"value": "2", "name": "发送中"},
		{"value": "3", "name": "成功"},
		{"value": "4", "name": "失败"},
	})
	pageBuilder.ListColumnAdd("create_time", "发送时间", "text", nil)
	pageBuilder.SetListColumnStyle("msg", "max-width:200px;overflow-x:auto")
	pageBuilder.SetStyle("table td{word-break: keep-all;white-space:nowrap;}")
	return nil, 0
}

// NodeListCondition 修改查询条件
func (that AppSmsRecord) NodeListCondition(pageBuilder *builder.PageBuilder, condition [][]interface{}) ([][]interface{}, error, int) {
	//追加查询条件
	//多开小程序
	uniappId := util.String2Int(pageBuilder.GetHttpParams().ByName("uniapp_id"))
	//多开小程序新增tab多开
	uniappTable := pageBuilder.GetHttpParams().ByName("_uniapp_table")
	if uniappTable != "" {
		tab := util.GetValue(pageBuilder.GetHttpRequest(), "tab")
		getUniappId := util.GetValue(pageBuilder.GetHttpRequest(), "uniapp_id")
		if tab != "" && getUniappId != "" {
			uniappId = util.String2Int(getUniappId)
		}
	}
	if uniappId > 0 {
		condition = append(condition, []interface{}{
			"uniapp_id", "=", uniappId,
		})
	}
	return condition, nil, 0
}
