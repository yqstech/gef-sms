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
	"github.com/yqstech/gef/Handles/adminHandle"
	"github.com/yqstech/gef/builder"
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
	pageBuilder.ListColumnClear()
	pageBuilder.SetListOrder("id desc")
	pageBuilder.ListTopBtnsClear()
	pageBuilder.ListRightBtnsClear()
	//获取列表
	upstreamOptions := that.SmsUpstreamList()
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
