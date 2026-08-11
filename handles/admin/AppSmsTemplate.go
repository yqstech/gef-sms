/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: AppSmsTemplate
 * @Version: 1.0.0
 * @Date: 2021/10/29 10:05 下午
 */

package admin

import (
	"errors"
	"github.com/wonderivan/logger"
	"github.com/yqstech/gef/Handles/adminHandle"
	"github.com/yqstech/gef/Models"
	"github.com/yqstech/gef/boot/db"
	"github.com/yqstech/gef/builder"
	"github.com/yqstech/gef/util"
	"strings"
)

type AppSmsTemplate struct {
	adminHandle.Base
}

// NodeBegin 开始
func (that AppSmsTemplate) NodeBegin(pageBuilder *builder.PageBuilder) (error, int) {
	uniappId := util.String2Int(pageBuilder.GetHttpParams().ByName("uniapp_id"))
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
			return err, 500
		}
		ids := db.Data2Ids(uniappList, "id")
		uniappIds = append(uniappIds, ids...)
	}
	//同步模板信息
	templateList, err := db.New().Table("tb_sms_template").
		Where("is_delete", 0).
		Where("status", "1").Get()
	if err != nil {
		logger.Error(err.Error())
		return errors.New("查询信息出错！"), 0
	}
	for _, templateInfo := range templateList {
		for _, uniId := range uniappIds {
			ManagerTemplate, err := db.New().Table("tb_app_sms_template").
				Where("is_delete", 0).
				Where("uniapp_id", uniId).
				Where("template_name", templateInfo["template_name"]).
				First()
			if err != nil {
				logger.Error(err.Error())
				return errors.New("查询信息出错！"), 0
			}
			if ManagerTemplate == nil {
				db.New().Table("tb_app_sms_template").Insert(map[string]interface{}{
					"uniapp_id":        uniId,
					"template_name":    templateInfo["template_name"],
					"template_content": templateInfo["default_content"],
					"create_time":      util.TimeNow(),
					"update_time":      util.TimeNow(),
					"status":           0,
				})
			}
		}
	}
	pageBuilder.SetTitle("短信模板管理")
	pageBuilder.SetPageName("短信模板")
	pageBuilder.SetTbName("tb_app_sms_template")
	return nil, 0
}

// NodeList 初始化列表
func (that AppSmsTemplate) NodeList(pageBuilder *builder.PageBuilder) (error, int) {
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
	//清除列表顶部和右侧按钮
	pageBuilder.ListRightBtnsClear()
	pageBuilder.ListTopBtnsClear()
	//重置右侧按钮
	pageBuilder.SetListRightBtns("edit", "disable", "enable")
	//获取列表
	templateOptions, err, code := Models.Model{}.SelectOptionsData("tb_sms_template", map[string]string{
		"template_name":  "value",
		"template_title": "name",
	}, "", "", "", "")
	if err != nil {
		return err, code
	}
	pageBuilder.ListColumnAdd("template_name", "模板标识", "text", nil)
	pageBuilder.ListColumnAdd("template_name", "模板名称", "array", templateOptions)
	pageBuilder.ListColumnAdd("template_out_id", "外部模板ID", "text", nil)
	pageBuilder.ListColumnAdd("template_content", "模板内容", "text", nil)
	pageBuilder.ListColumnAdd("status", "状态", "array", Models.OptionModels{}.ByKey("status", true))
	return nil, 0
}

// NodeListCondition 修改查询条件
func (that AppSmsTemplate) NodeListCondition(pageBuilder *builder.PageBuilder, condition [][]interface{}) ([][]interface{}, error, int) {
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
func (that AppSmsTemplate) NodeForm(pageBuilder *builder.PageBuilder, id int64) (error, int) {
	//查询通道
	if id <= 0 {
		return errors.New("获取通道ID失败！"), 0
	}
	uniappId := util.String2Int(pageBuilder.GetHttpParams().ByName("uniapp_id"))
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
			return err, 500
		}
		ids := db.Data2Ids(uniappList, "id")
		uniappIds = append(uniappIds, ids...)
	}
	//查询应用短信模板
	managerTemplate, err := db.New().Table("tb_app_sms_template").
		WhereIn("uniapp_id", uniappIds).
		Where("id", id).
		Where("is_delete", 0).
		First()
	if err != nil {
		return err, 0
	}
	if managerTemplate == nil {
		return errors.New("获取通道信息失败！"), 0
	}
	// 查询短信模板
	templateInfo, err := db.New().Table("tb_sms_template").
		Where("template_name", managerTemplate["template_name"]).
		Where("is_delete", 0).
		Where("status", 1).First()
	if err != nil {
		return err, 0
	}
	vars := strings.Split(templateInfo["template_vars"].(string), ",")
	contentNotices := "当前模板支持的变量有："
	for _, v := range vars {
		contentNotices = contentNotices + "{{" + v + "}}"
	}

	pageBuilder.FormFieldsAdd("template_out_id", "text", "外部模板ID", "例如阿里云模板ID SMS_123456", "", false, nil, "", nil)
	pageBuilder.FormFieldsAdd("template_content", "textarea", "模板内容", "变量格式为变量+双花括号，例如验证码：{{code}}", "", false, nil, "", nil)
	pageBuilder.FormFieldsAdd("", "notice", "", contentNotices, "", false, nil, "", nil)
	return nil, 0
}

// NodeAutoData 新增或修改调用，新增自动补充小程序ID参数
func (that AppSmsTemplate) NodeAutoData(pageBuilder *builder.PageBuilder, postData map[string]interface{}, action string) (map[string]interface{}, error, int) {
	//多开小程序新增tab多开
	uniappTable := pageBuilder.GetHttpParams().ByName("_uniapp_table")
	if uniappTable != "" {
		//后台管理多个小程序，这里是表单选择无需重复添加
		return postData, nil, 0
	}
	uniappId := util.String2Int(pageBuilder.HttpParams.ByName("uniapp_id"))
	if action == "add" {
		postData["uniapp_id"] = uniappId
	}
	return postData, nil, 0
}
