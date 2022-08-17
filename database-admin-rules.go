/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: database-admin-rules
 * @Version: 1.0.0
 * @Date: 2022/8/16 22:19
 */

package sms

var adminRules = []map[string]interface{}{
	{
		"route": "#config",
		"children": []map[string]interface{}{
			{
				"name":      "短信设置",
				"type":      1,
				"is_compel": 0,
				"icon":      "icon-message-3-fill",
				"route":     "#sms_config",
				"children": []map[string]interface{}{
					{
						"name":      "短信通道",
						"type":      1,
						"is_compel": 0,
						"icon":      "",
						"route":     "/app_sms_upstream/index",
						"children":  []map[string]interface{}{},
					},
					{
						"name":      "短信模板",
						"type":      1,
						"is_compel": 0,
						"icon":      "",
						"route":     "/app_sms_template/index",
						"children":  []map[string]interface{}{},
					},
					{
						"name":      "短信黑白名单",
						"type":      1,
						"is_compel": 0,
						"icon":      "",
						"route":     "/app_sms_black_white/index",
						"children":  []map[string]interface{}{},
					},
					{
						"name":      "短信防火墙",
						"type":      1,
						"is_compel": 0,
						"icon":      "",
						"route":     "/app_sms_hold_back/index",
						"children":  []map[string]interface{}{},
					}, {
						"name":      "短信发送记录",
						"type":      1,
						"is_compel": 0,
						"icon":      "",
						"route":     "/app_sms_record/index",
						"children":  []map[string]interface{}{},
					},
				},
			},
		},
	},
	{
		"route": "#dev",
		"children": []map[string]interface{}{
			{
				"name":      "短信模块",
				"type":      1,
				"is_compel": 0,
				"icon":      "layui-icon-reply-fill",
				"route":     "#sms",
				"children": []map[string]interface{}{
					{
						"name":      "短信通道配置",
						"type":      1,
						"is_compel": 0,
						"icon":      "",
						"route":     "/sms_upstream/index",
						"children": []map[string]interface{}{
							{
								"name":      "短信通道配置参数",
								"type":      1,
								"is_compel": 0,
								"icon":      "",
								"route":     "/sms_upstream_params/index",
								"children":  []map[string]interface{}{},
							},
						},
					},
					{
						"name":      "短信模板管理",
						"type":      1,
						"is_compel": 0,
						"icon":      "",
						"route":     "/sms_template/index",
						"children":  []map[string]interface{}{},
					},
				},
			},
		},
	},
}
