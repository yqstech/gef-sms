/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: AppSmsBlackWhite
 * @Version: 1.0.0
 * @Date: 2022/3/8 9:42 下午
 */

package AdminHandles

import (
	"github.com/yqstech/gef-sms/SmsModels"
	"github.com/yqstech/gef/Handles/adminHandle"
	"github.com/yqstech/gef/Models"
	"github.com/yqstech/gef/builder"
	"github.com/yqstech/gef/util"
)

type AppSmsBlackWhite struct {
	adminHandle.Base
}

// NodeBegin 开始
func (that AppSmsBlackWhite) NodeBegin(pageBuilder *builder.PageBuilder) (error, int) {
	pageBuilder.SetTitle("短信黑名单和白名单")
	pageBuilder.SetPageName("黑白名单")
	pageBuilder.SetTbName("tb_app_sms_black_white")
	return nil, 0
}

// NodeList 初始化列表
func (that AppSmsBlackWhite) NodeList(pageBuilder *builder.PageBuilder) (error, int) {
	pageBuilder.ListColumnClear()
	pageBuilder.ListColumnAdd("type", "类型", "array", SmsModels.SmsBlackWhiteTypes)
	pageBuilder.ListColumnAdd("rule_type", "匹配类型", "array", SmsModels.SmsBlackWhiteRuleTypes)
	pageBuilder.ListColumnAdd("rule", "手机号或ip地址", "text", nil)
	pageBuilder.ListColumnAdd("note", "备注", "text", nil)
	pageBuilder.ListColumnAdd("status", "状态", "array", Models.OptionModels{}.ByKey("status", true))

	return nil, 0
}

// NodeListCondition 修改查询条件
func (that AppSmsBlackWhite) NodeListCondition(pageBuilder *builder.PageBuilder, condition [][]interface{}) ([][]interface{}, error, int) {
	//追加查询条件
	//多开小程序
	uniappId := util.String2Int(pageBuilder.GetHttpParams().ByName("uniapp_id"))
	condition = append(condition, []interface{}{
		"uniapp_id", "=", uniappId,
	})
	return condition, nil, 0
}

// NodeForm 初始化表单
func (that AppSmsBlackWhite) NodeForm(pageBuilder *builder.PageBuilder, id int64) (error, int) {
	pageBuilder.FormFieldsAdd("type", "radio", "类型", "", "0", true, SmsModels.SmsBlackWhiteTypes, "", nil)
	pageBuilder.FormFieldsAdd("rule_type", "radio", "匹配类型", "", "0", true, SmsModels.SmsBlackWhiteRuleTypes, "", nil)
	pageBuilder.FormFieldsAdd("rule", "text", "匹配内容", "手机号或ip地址", "", true, nil, "", nil)
	pageBuilder.FormFieldsAdd("note", "text", "备注", "", "", false, nil, "", nil)
	return nil, 0
}

// NodeAutoData 新增或修改调用，新增自动补充小程序ID参数
func (that AppSmsBlackWhite) NodeAutoData(pageBuilder *builder.PageBuilder, postData map[string]interface{}, action string) (map[string]interface{}, error, int) {
	uniappId := util.String2Int(pageBuilder.HttpParams.ByName("uniapp_id"))
	if action == "add" {
		postData["uniapp_id"] = uniappId
	}
	return postData, nil, 0
}
