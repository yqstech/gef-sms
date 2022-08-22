/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: SmsUpstreamParams
 * @Version: 1.0.0
 * @Date: 2021/10/28 10:16 下午
 */

package AdminHandles

import (
	"github.com/yqstech/gef/Handles/adminHandle"
	"github.com/yqstech/gef/builder"
	"github.com/yqstech/gef/util"
)

type SmsUpstreamParams struct {
	adminHandle.Base
}

// NodeBegin 开始
func (that SmsUpstreamParams) NodeBegin(pageBuilder *builder.PageBuilder) (error, int) {
	pageBuilder.SetTitle("短信通道配置项")
	pageBuilder.SetPageName("通道配置项")
	pageBuilder.SetTbName("tb_sms_upstream_params")
	return nil, 0
}

// NodeList 初始化列表
func (that SmsUpstreamParams) NodeList(pageBuilder *builder.PageBuilder) (error, int) {
	pageBuilder.SetListOrder("id asc")
	//重设新增按钮增加
	ActionUrl := "add"
	upstreamId := util.GetValue(pageBuilder.GetHttpRequest(), "id")
	if upstreamId != "" {
		ActionUrl = ActionUrl + "?upstream_id=" + upstreamId
		pageBuilder.SetButton("add", builder.Button{
			ButtonName: "新增通道配置项",
			Action:     "add",
			ActionType: 2,
			ActionUrl:  ActionUrl,
			Class:      "def",
			Icon:       "ri-play-list-add-fill",
			Display:    "",
			Expand: map[string]string{
				"w": "98%",
				"h": "98%",
			},
		})
	}

	//获取列表
	upstreamList := that.SmsUpstreamList()
	pageBuilder.ListColumnAdd("upstream_id", "短信通道名称", "array", upstreamList)
	pageBuilder.ListColumnAdd("param_name", "配置项", "text", nil)
	pageBuilder.ListColumnAdd("param_title", "配置项名称", "text", nil)
	return nil, 0
}

// NodeListCondition 修改查询条件
func (that SmsUpstreamParams) NodeListCondition(pageBuilder *builder.PageBuilder, condition [][]interface{}) ([][]interface{}, error, int) {
	upstreamID := 0
	upstreamId := util.GetValue(pageBuilder.GetHttpRequest(), "id")
	if upstreamId != "" {
		upstreamID = util.String2Int(upstreamId)
		//追加查询条件
		condition = append(condition, []interface{}{
			"upstream_id", "=", upstreamID,
		})
	}

	return condition, nil, 0
}

// NodeForm 初始化表单
func (that SmsUpstreamParams) NodeForm(pageBuilder *builder.PageBuilder, id int64) (error, int) {
	upstreamId := util.GetValue(pageBuilder.GetHttpRequest(), "upstream_id")
	if upstreamId != "" {
		pageBuilder.FormFieldsAdd("upstream_id", "hidden", "配置项", "", upstreamId, true, nil, "", nil)
	} else {
		//获取列表
		upstreamList := that.SmsUpstreamList()
		pageBuilder.FormFieldsAdd("upstream_id", "select", "所属通道", "", "", true, upstreamList, "", nil)
	}

	pageBuilder.FormFieldsAdd("param_name", "text", "配置项", "", "", true, nil, "", nil)
	pageBuilder.FormFieldsAdd("param_title", "text", "配置项名称", "", "", true, nil, "", nil)
	return nil, 0
}
