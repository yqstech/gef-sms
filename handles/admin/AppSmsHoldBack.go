/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: AppSmsHoldBack
 * @Version: 1.0.0
 * @Date: 2022/3/8 9:43 下午
 */

package admin

import (
	"github.com/gohouse/gorose/v2"
	"github.com/wonderivan/logger"
	"github.com/yqstech/gef-sms/models"
	"github.com/yqstech/gef/Handles/adminHandle"
	"github.com/yqstech/gef/Models"
	"github.com/yqstech/gef/boot/db"
	"github.com/yqstech/gef/builder"
	"github.com/yqstech/gef/util"
)

type AppSmsHoldBack struct {
	adminHandle.Base
}

// NodeBegin 开始
func (that AppSmsHoldBack) NodeBegin(pageBuilder *builder.PageBuilder) (error, int) {
	pageBuilder.SetTitle("短信防火墙规则")
	pageBuilder.SetPageName("规则")
	pageBuilder.SetTbName("tb_app_sms_hold_back")
	return nil, 0
}

// NodeList 初始化列表
func (that AppSmsHoldBack) NodeList(pageBuilder *builder.PageBuilder) (error, int) {
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
	pageBuilder.ListColumnAdd("rule_type", "规则类型", "array", models.SmsHoldBackRuleTypes)
	pageBuilder.ListColumnAdd("range_second", "几秒钟内", "text", nil)
	pageBuilder.ListColumnAdd("sms_max", "短信超过几条", "text", nil)
	pageBuilder.ListColumnAdd("action", "执行操作", "array", models.SmsHoldBackActions)
	pageBuilder.ListColumnAdd("frozen_second", "暂停秒数", "text", nil)
	pageBuilder.ListColumnAdd("note", "备注", "text", nil)
	pageBuilder.ListColumnAdd("status", "状态", "array", Models.OptionModels{}.ByKey("status", true))
	return nil, 0
}

// NodeListCondition 修改查询条件
func (that AppSmsHoldBack) NodeListCondition(pageBuilder *builder.PageBuilder, condition [][]interface{}) ([][]interface{}, error, int) {
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
	condition = append(condition, []interface{}{
		"uniapp_id", "=", uniappId,
	})
	return condition, nil, 0
}

// NodeListData 重写列表数据
func (that AppSmsHoldBack) NodeListData(pageBuilder *builder.PageBuilder, data []gorose.Data) ([]gorose.Data, error, int) {
	for k, v := range data {
		if v["frozen_second"].(int64) < 0 {
			data[k]["frozen_second"] = "#"
		}
	}
	return data, nil, 0
}

// NodeForm 初始化表单
func (that AppSmsHoldBack) NodeForm(pageBuilder *builder.PageBuilder, id int64) (error, int) {
	uniappTable := pageBuilder.GetHttpParams().ByName("_uniapp_table")
	if uniappTable != "" {
		//后台管理多个小程序，在这里新增一个多选小程序
		uniappOptions, err, code := Models.Model{}.SelectOptionsData(uniappTable, map[string]string{
			"id":       "value",
			"app_name": "name",
		}, "0", "系统默认", "status=1", "index_num asc,id asc")
		if err != nil {
			return err, code
		}
		pageBuilder.FormFieldsAdd("uniapp_id", "select-sm", "选择应用", "请选择所属应用", "", true, uniappOptions, "", nil)
	}
	pageBuilder.FormFieldsAdd("rule_type", "radio", "规则类型", "", "0", true, models.SmsHoldBackRuleTypes, "", nil)
	pageBuilder.FormFieldsAdd("range_second", "text", "几秒钟内", "填写大于0的整数", "", false, nil, "", nil)
	pageBuilder.FormFieldsAdd("sms_max", "text", "短信超过几条", "填写大于0的整数", "", false, nil, "", nil)
	pageBuilder.FormFieldsAdd("action", "radio", "执行操作", "", "1", false, models.SmsHoldBackActions, "", map[string]interface{}{
		"if": "formFields.rule_type<2",
	})
	pageBuilder.FormFieldsAdd("action2", "radio", "执行操作", "", "1", false, []map[string]interface{}{{"name": "暂停短信功能(白名单除外)", "value": "1"}}, "", map[string]interface{}{
		"if": "formFields.rule_type==2",
	})
	pageBuilder.FormFieldsAdd("frozen_second", "text", "暂停秒数", "填写大于0的整数", "", false, nil, "", map[string]interface{}{
		"if": "!(formFields.rule_type<2 && formFields.action==2)",
	})
	pageBuilder.FormFieldsAdd("note", "text", "备注", "", "", false, nil, "", nil)
	return nil, 0
}

// NodeSaveData 表单保存数据前使用
func (that AppSmsHoldBack) NodeSaveData(pageBuilder *builder.PageBuilder, oldData gorose.Data, postData map[string]interface{}) (map[string]interface{}, error, int) {
	action := util.PostValue(pageBuilder.GetHttpRequest(), "action")
	if action != "fastUpdate" {
		if postData["rule_type"] == "2" {
			postData["action"] = postData["action2"]
		} else {
			if postData["action"] == "2" {
				postData["frozen_second"] = -1
			}
		}
		delete(postData, "action2")
	}
	return postData, nil, 0
}

// NodeAutoData 新增或修改调用，新增自动补充小程序ID参数
func (that AppSmsHoldBack) NodeAutoData(pageBuilder *builder.PageBuilder, postData map[string]interface{}, action string) (map[string]interface{}, error, int) {
	//多开小程序新增tab多开
	uniappTable := pageBuilder.GetHttpParams().ByName("_uniapp_table")
	if uniappTable != "" {
		//后台管理多个小程序，这里是表单选择无需重复添加
		return postData, nil, 0
	}
	//多开小程序-小程序后台-自动添加属性
	uniappId := util.String2Int(pageBuilder.HttpParams.ByName("uniapp_id"))
	if action == "add" {
		postData["uniapp_id"] = uniappId
	}
	return postData, nil, 0
}

// NodeAutoCondition 新增、修改、删除权限限定
func (that AppSmsHoldBack) NodeAutoCondition(pageBuilder *builder.PageBuilder, condition [][]interface{}) ([][]interface{}, error, int) {
	condition, err, code := that.Base.NodeAutoCondition(pageBuilder, condition)
	if err != nil {
		return condition, err, code
	}
	uniappId := util.String2Int(pageBuilder.HttpParams.ByName("uniapp_id"))
	uniappIds := []any{uniappId}
	//多开小程序新增tab多开
	uniappTable := pageBuilder.GetHttpParams().ByName("_uniapp_table")
	if uniappTable != "" {
		uniappList, err := db.New().Table(uniappTable).
			Where("is_delete", 0).
			Where("status", 1).
			Order("index_num asc,id asc").
			Get()
		if err != nil {
			logger.Error(err.Error())
			return condition, err, 500
		}
		ids := db.Data2Ids(uniappList, "id")
		uniappIds = append(uniappIds, ids...)
	}
	condition = append(condition, []interface{}{"uniapp_id", "in", uniappIds})
	return condition, nil, 0
}
