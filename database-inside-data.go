/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: database-inside-data
 * @Version: 1.0.0
 * @Date: 2022/8/16 22:20
 */

package sms

import "github.com/yqstech/gef"

var insideData = []gef.InsideData{
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
	}}, {TableName: "tb_sms_upstream", Condition: [][]interface{}{{"id", "3"}}, Data: map[string]interface{}{
		"id":            3,
		"upstream_name": "云市场短信",
		"event_name":    "SmsAm",
		"index_num":     "3",
	}}, {TableName: "tb_sms_upstream", Condition: [][]interface{}{{"id", "4"}}, Data: map[string]interface{}{
		"id":            4,
		"upstream_name": "模拟发短信",
		"event_name":    "SmsMock",
		"index_num":     "4",
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
	}},
}
