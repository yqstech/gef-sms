/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description: 短信通道管理
 * @File: SmsUpstream
 * @Version: 1.0.0
 * @Date: 2021/10/28 10:16 下午
 */

package AdminHandles

import (
	"github.com/yqstech/gef/Handles/adminHandle"
	"github.com/yqstech/gef/Models"
	"github.com/yqstech/gef/builder"
	"github.com/yqstech/gef/config"
)

type SmsUpstream struct {
	adminHandle.Base
}

var SmsEventList = []map[string]interface{}{
	{"name": "SmsAli", "value": "SmsAli"},
	{"name": "SmsJdcx", "value": "SmsJdcx"},
	{"name": "SmsAm", "value": "SmsAm"},
	{"name": "SmsMock", "value": "SmsMock"},
}

// NodeBegin 开始
func (that SmsUpstream) NodeBegin(pageBuilder *builder.PageBuilder) (error, int) {
	pageBuilder.SetTitle("短信通道管理")
	pageBuilder.SetPageName("短信通道")
	pageBuilder.SetTbName("tb_sms_upstream")
	return nil, 0
}

// NodeList 初始化列表
func (that SmsUpstream) NodeList(pageBuilder *builder.PageBuilder) (error, int) {
	pageBuilder.SetListOrder("id asc")
	//新增右侧参数设置按钮
	pageBuilder.SetButton("params", builder.Button{
		ButtonName: "配置项",
		Action:     "/sms_upstream_params/index",
		ActionType: 2,
		ActionUrl:  config.AdminPath + "/sms_upstream_params/index",
		Class:      "",
		Icon:       "ri-list-settings-fill",
		Display:    "(!item.btn_params || item.btn_params!='hide')",
		Expand: map[string]string{
			"w": "98%",
			"h": "98%",
		},
	})
	//!重置右侧按钮
	pageBuilder.SetListRightBtns("edit", "disable", "enable", "params", "delete")

	pageBuilder.ListColumnAdd("upstream_name", "通道名称", "text", nil)
	pageBuilder.ListColumnAdd("event_name", "关联事件", "text", nil)
	pageBuilder.ListColumnAdd("status", "状态", "array", Models.OptionModels{}.ById(2, true))
	return nil, 0
}

// NodeForm 初始化表单
func (that SmsUpstream) NodeForm(pageBuilder *builder.PageBuilder, id int64) (error, int) {
	pageBuilder.FormFieldsAdd("upstream_name", "text", "通道名称", "", "", true, nil, "", nil)
	pageBuilder.FormFieldsAdd("event_name", "select", "关联事件", "", "", true, SmsEventList, "", nil)
	return nil, 0
}
