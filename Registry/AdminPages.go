/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: AdminPages
 * @Version: 1.0.0
 * @Date: 2022/8/16 23:41
 */

package Registry

import (
	"github.com/yqstech/gef/GoEasy/EasyApp"
	"github.com/yqstech/gef-sms/Handles/AdminHandles"
)

var AdminPages = map[string]EasyApp.AppPage{
	"sms_upstream":        AdminHandles.SmsUpstream{},       //短信通道管理
	"sms_upstream_params": AdminHandles.SmsUpstreamParams{}, //短信通道配置项
	"sms_template":        AdminHandles.SmsTemplate{},       //短信模板管理
	"app_sms_upstream":    AdminHandles.AppSmsUpstream{},    //应用短信通道
	"app_sms_template":    AdminHandles.AppSmsTemplate{},    //应用短信模板
	"app_sms_black_white": AdminHandles.AppSmsBlackWhite{},  //应用短信黑白名单
	"app_sms_record":      AdminHandles.AppSmsRecord{},      //应用短信记录
	"app_sms_hold_back":   AdminHandles.AppSmsHoldBack{},    //应用短信防火墙
}
