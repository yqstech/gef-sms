/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: AppSmsTemplate
 * @Version: 1.0.0
 * @Date: 2021/10/29 10:05 下午
 */

package AdminHandles

import (
	"errors"
	"github.com/wonderivan/logger"
	"github.com/yqstech/gef/Handles/adminHandle"
	"github.com/yqstech/gef/Models"
	"github.com/yqstech/gef/Utils/db"
	"github.com/yqstech/gef/builder"
	"github.com/yqstech/gef/util"
	"strings"
)

type AppSmsTemplate struct {
	adminHandle.Base
}

// NodeBegin 开始
func (that AppSmsTemplate) NodeBegin(pageBuilder *builder.PageBuilder) (error, int) {
	//同步模板信息
	templateList, err := db.New().Table("tb_sms_template").
		Where("is_delete", 0).
		Where("status", "1").Get()
	if err != nil {
		logger.Error(err.Error())
		return errors.New("查询信息出错！"), 0
	}
	for _, templateInfo := range templateList {
		ManagerTemplate, err := db.New().Table("tb_app_sms_template").
			Where("template_name", templateInfo["template_name"]).
			Where("is_delete", 0).
			First()
		if err != nil {
			logger.Error(err.Error())
			return errors.New("查询信息出错！"), 0
		}
		if ManagerTemplate == nil {
			db.New().Table("tb_app_sms_template").Insert(map[string]interface{}{
				"template_name":    templateInfo["template_name"],
				"template_content": templateInfo["default_content"],
				"create_time":      util.TimeNow(),
				"update_time":      util.TimeNow(),
				"status":           0,
			})
		}
	}
	pageBuilder.SetTitle("短信模板管理")
	pageBuilder.SetPageName("短信模板")
	pageBuilder.SetTbName("tb_app_sms_template")
	return nil, 0
}

// NodeList 初始化列表
func (that AppSmsTemplate) NodeList(pageBuilder *builder.PageBuilder) (error, int) {
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
	pageBuilder.ListColumnAdd("status", "状态", "array", Models.OptionModels{}.ById(2, true))
	return nil, 0
}

// NodeForm 初始化表单
func (that AppSmsTemplate) NodeForm(pageBuilder *builder.PageBuilder, id int64) (error, int) {
	//查询通道
	if id <= 0 {
		return errors.New("获取通道ID失败！"), 0
	}
	//查询应用短信模板
	managerTemplate, err := db.New().Table("tb_app_sms_template").
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
