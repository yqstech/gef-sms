/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: database-inside-data
 * @Version: 1.0.0
 * @Date: 2022/8/16 22:20
 */

package sms

import (
	"github.com/yqstech/gef/dbManager"
)

var insideData = []dbManager.InsideData{
	{TableName: "tb_sms_upstream", Condition: [][]interface{}{{"id", "1"}}, Data: map[string]interface{}{
		"id":            1,
		"upstream_name": "阿里云短信",
		"event_name":    "SmsAli",
		"index_num":     "1",
	}}, {TableName: "tb_sms_upstream", Condition: [][]interface{}{{"id", "2"}}, Data: map[string]interface{}{
		"id":            2,
		"upstream_name": "万象&创信",
		"event_name":    "SmsJdcx",
		"index_num":     "2",
		"status":        0,
	}}, {TableName: "tb_sms_upstream", Condition: [][]interface{}{{"id", "3"}}, Data: map[string]interface{}{
		"id":            3,
		"upstream_name": "云市场短信",
		"event_name":    "SmsAm",
		"index_num":     "3",
		"status":        0,
	}}, {TableName: "tb_sms_upstream", Condition: [][]interface{}{{"id", "4"}}, Data: map[string]interface{}{
		"id":            4,
		"upstream_name": "模拟发短信",
		"event_name":    "SmsMock",
		"index_num":     "4",
	}}, {TableName: "tb_sms_upstream", Condition: [][]interface{}{{"id", "5"}}, Data: map[string]interface{}{
		"id":            5,
		"upstream_name": "创信",
		"event_name":    "SmsCdcx",
		"index_num":     "5",
	}}, {TableName: "tb_sms_upstream", Condition: [][]interface{}{{"id", "6"}}, Data: map[string]interface{}{
		"id":            6,
		"upstream_name": "掌骏",
		"event_name":    "SmsZj",
		"index_num":     "6",
	}}, {TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "1"}}, Data: map[string]interface{}{
		"id":          1,
		"upstream_id": 1,
		"param_name":  "accessKeyId",
		"param_title": "accessKeyId",
	}}, {TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "2"}}, Data: map[string]interface{}{
		"id":          2,
		"upstream_id": 1,
		"param_name":  "accessKeySecret",
		"param_title": "accessKeySecret",
	}}, {TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "3"}}, Data: map[string]interface{}{
		"id":          3,
		"upstream_id": 1,
		"param_name":  "SignName",
		"param_title": "短信签名",
	}}, {TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "4"}}, Data: map[string]interface{}{
		"id":          4,
		"upstream_id": 2,
		"param_name":  "appkey",
		"param_title": "万象接口凭据",
	}}, {TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "5"}}, Data: map[string]interface{}{
		"id":          5,
		"upstream_id": 2,
		"param_name":  "sign",
		"param_title": "短信签名",
	}}, {TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "6"}}, Data: map[string]interface{}{
		"id":          6,
		"upstream_id": 3,
		"param_name":  "sms_url",
		"param_title": "短信接口地址",
	}}, {TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "7"}}, Data: map[string]interface{}{
		"id":          7,
		"upstream_id": 3,
		"param_name":  "api_token",
		"param_title": "接口凭据(token)",
	}}, {TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "8"}}, Data: map[string]interface{}{
		"id":          8,
		"upstream_id": 3,
		"param_name":  "sign",
		"param_title": "短信签名",
	}}, {TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "9"}}, Data: map[string]interface{}{
		"id":          9,
		"upstream_id": 5,
		"param_name":  "sign",
		"param_title": "短信签名",
	}}, {TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "10"}}, Data: map[string]interface{}{
		"id":          10,
		"upstream_id": 5,
		"param_name":  "account",
		"param_title": "创信账户",
	}}, {TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "11"}}, Data: map[string]interface{}{
		"id":          11,
		"upstream_id": 5,
		"param_name":  "password",
		"param_title": "账户密码",
	}}, {TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "12"}}, Data: map[string]interface{}{
		"id":          12,
		"upstream_id": 5,
		"param_name":  "extno",
		"param_title": "SP服务号",
	}}, {TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "13"}}, Data: map[string]interface{}{
		"id":          13,
		"upstream_id": 5,
		"param_name":  "url",
		"param_title": "接口地址",
	}},
	{TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "14"}}, Data: map[string]interface{}{
		"id":          14,
		"upstream_id": 6,
		"param_name":  "sign",
		"param_title": "短信签名",
	}}, {TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "15"}}, Data: map[string]interface{}{
		"id":          15,
		"upstream_id": 6,
		"param_name":  "account",
		"param_title": "账户名称",
	}}, {TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "16"}}, Data: map[string]interface{}{
		"id":          16,
		"upstream_id": 6,
		"param_name":  "password",
		"param_title": "账户密码",
	}}, {TableName: "tb_sms_upstream_params", Condition: [][]interface{}{{"id", "17"}}, Data: map[string]interface{}{
		"id":          17,
		"upstream_id": 6,
		"param_name":  "url",
		"param_title": "接口地址",
	}},
	//短信模板
	{TableName: "tb_sms_template", Condition: [][]interface{}{{"id", "1"}}, Data: map[string]interface{}{
		"id":              1,
		"template_name":   "default",
		"template_title":  "默认模板",
		"template_vars":   "code",
		"default_content": "您的验证码是{{code}}，请不要告诉任何人！",
	}}, {TableName: "tb_app_sms_template", Condition: [][]interface{}{{"id", "1"}}, Data: map[string]interface{}{
		"id":               1,
		"template_name":    "default",
		"template_out_id":  "SMS_123456",
		"template_content": "您的验证码是{{code}}，请不要告诉任何人！",
	}},
}
