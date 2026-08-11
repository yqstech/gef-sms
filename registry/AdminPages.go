/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: AdminPages
 * @Version: 1.0.0
 * @Date: 2022/8/16 23:41
 */

package registry

import (
	"github.com/yqstech/gef-sms/handles/admin"
	"github.com/yqstech/gef/builder"
)

var AdminPages = map[string]builder.NodePager{
	"sms_upstream":        &admin.SmsUpstream{},       //短信通道管理
	"sms_upstream_params": &admin.SmsUpstreamParams{}, //短信通道配置项
	"sms_template":        &admin.SmsTemplate{},       //短信模板管理

	"app_sms_upstream":    &admin.AppSmsUpstream{},   //应用短信通道
	"app_sms_template":    &admin.AppSmsTemplate{},   //应用短信模板
	"app_sms_black_white": &admin.AppSmsBlackWhite{}, //应用短信黑白名单
	"app_sms_record":      &admin.AppSmsRecord{},     //应用短信记录
	"app_sms_hold_back":   &admin.AppSmsHoldBack{},   //应用短信防火墙
}
