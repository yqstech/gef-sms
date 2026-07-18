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
	"github.com/wonderivan/logger"
	"github.com/yqstech/gef-sms/SmsModels"
	"github.com/yqstech/gef/Handles/adminHandle"
	"github.com/yqstech/gef/Models"
	"github.com/yqstech/gef/Utils/db"
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

// NodeForm 初始化表单
func (that AppSmsBlackWhite) NodeForm(pageBuilder *builder.PageBuilder, id int64) (error, int) {
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
	pageBuilder.FormFieldsAdd("type", "radio", "类型", "", "0", true, SmsModels.SmsBlackWhiteTypes, "", nil)
	pageBuilder.FormFieldsAdd("rule_type", "radio", "匹配类型", "", "0", true, SmsModels.SmsBlackWhiteRuleTypes, "", nil)
	pageBuilder.FormFieldsAdd("rule", "text", "匹配内容", "手机号或ip地址", "", true, nil, "", nil)
	pageBuilder.FormFieldsAdd("note", "text", "备注", "", "", false, nil, "", nil)
	return nil, 0
}

// NodeAutoData 新增或修改调用，新增自动补充小程序ID参数
func (that AppSmsBlackWhite) NodeAutoData(pageBuilder *builder.PageBuilder, postData map[string]interface{}, action string) (map[string]interface{}, error, int) {
	//多开小程序新增tab多开
	uniappTable := pageBuilder.GetHttpParams().ByName("_uniapp_table")
	if uniappTable != "" {
		//后台管理多个小程序，这里是表单选择无需重复添加
		//logger.Info("postData", util.JsonEncode(postData))
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
func (that AppSmsBlackWhite) NodeAutoCondition(pageBuilder *builder.PageBuilder, condition [][]interface{}) ([][]interface{}, error, int) {
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
