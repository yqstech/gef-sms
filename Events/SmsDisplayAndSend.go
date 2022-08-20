/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: SmsDisplayAndSend
 * @Version: 1.0.0
 * @Date: 2022/8/17 13:17
 */

package Events

import (
	"bytes"
	"errors"
	"github.com/yqstech/gef/GoEasy/Event"
	"github.com/yqstech/gef/GoEasy/Utils/db"
	"github.com/wonderivan/logger"
	"strings"
	"text/template"
)

type SmsDisplayAndSend struct {
}

// Do 短信模板解析并发送
func (that SmsDisplayAndSend) Do(eventName string, data ...interface{}) (error, int) {
	tel := data[0].(string)
	ip := data[1].(string)
	templateName := data[2].(string)
	templateParams := data[3].(map[string]interface{})
	
	//获取外部模板ID和短信内容
	//查询应用短信模板信息
	SmsTemplate, err := db.New().Table("tb_app_sms_template").
		Where("template_name", templateName).
		Where("is_delete", 0).
		Where("status", 1).
		First()
	if err != nil {
		logger.Error(err.Error())
		return errors.New("系统运行错误"), 500
	}
	if SmsTemplate == nil {
		return errors.New("未设置短信模板"), 202
	}
	//短信模板生成短信内容
	templateContent := SmsTemplate["template_content"].(string)
	templateOutId := SmsTemplate["template_out_id"].(string)
	if templateContent == "" {
		return errors.New("未设置短信模板"), 203
	}
	//短信模板转为go模板
	templateContent = strings.Replace(templateContent, "{{", "{{.", -1)
	//go模板
	tmpl, err := template.New("sms_template").Parse(templateContent)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("短信模板解析错误！"), 500
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, templateParams)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("短信模板设置错误！"), 500
	}
	
	//短信发送参数
	ps := map[string]interface{}{
		"template_name":   templateName, //模板名称
		"content":         buf.String(), //短信内容，接口如果需要上传短信内容，则使用这个
		"template_out_id": templateOutId, //短信外模板ID
		"params":          templateParams,//短信参数，远程模板，需要在短信服务商那里渲染模板，
	}
	
	err, code := Event.Trigger("SmsSend", tel, ip, ps)
	return err, code
}
