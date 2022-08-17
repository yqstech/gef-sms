/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: database-inside-data
 * @Version: 1.0.0
 * @Date: 2022/8/16 22:20
 */

package sms

import "github.com/gef"

var insideData = []gef.InsideData{
	{TableName: "tb_admin", Condition: [][]interface{}{{"id", "1"}}, Data: map[string]interface{}{
		"group_id": 0,
		"name":     "系统管理员",
		"account":  "root",
		"password": "99b74a36ec072636d319ff7b49917d95",
	}},
	{TableName: "tb_configs_group", Condition: [][]interface{}{{"id", "1"}}, Data: map[string]interface{}{
		"id":         1,
		"group_name": "系统配置",
		"note":       "系统全局配置项",
	}},
	{TableName: "tb_configs", Condition: [][]interface{}{{"name", "upload_extension"}}, Data: map[string]interface{}{
		"group_id":   1,
		"name":       "upload_extension",
		"value":      "jpg,png,gif,jpeg,JPG,PNG,GIF,JPEG,xml,XML",
		"title":      "可用附件类型",
		"notice":     "允许上传的附件拓展名，多个拓展名用英文逗号分割，不用加点",
		"field_type": "text",
		"index_num":  1,
	}},
	{TableName: "tb_app_configs", Condition: [][]interface{}{{"name", "upload_extension"}}, Data: map[string]interface{}{
		"group_id": 1,
		"name":     "upload_extension",
		"value":    "jpg,png,gif,jpeg,JPG,PNG,GIF,JPEG,xml,XML",
	}},
	{TableName: "tb_option_models", Condition: [][]interface{}{{"id", 1}}, Data: map[string]interface{}{
		"id":          1,
		"name":        "是|否",
		"static_data": "[{\"name\":\"是\",\"value\":\"1\",\"color\":\"#66CC00\",\"icon\":\"ri-checkbox-circle-fill\"},{\"name\":\"否\",\"value\":\"0\",\"color\":\"#FF6666\",\"icon\":\"ri-forbid-fill\"}]",
	}},
	{TableName: "tb_option_models", Condition: [][]interface{}{{"id", 2}}, Data: map[string]interface{}{
		"id":          2,
		"name":        "启用|禁用",
		"static_data": "[{\"name\":\"启用\",\"value\":\"1\",\"color\":\"#66CC00\",\"icon\":\"ri-checkbox-circle-fill\"},{\"name\":\"禁用\",\"value\":\"0\",\"color\":\"#FF6666\",\"icon\":\"ri-forbid-fill\"}]",
	}},
	{TableName: "tb_option_models", Condition: [][]interface{}{{"id", 3}}, Data: map[string]interface{}{
		"id":          3,
		"name":        "显示|隐藏",
		"static_data": "[{\"name\":\"显示\",\"value\":\"1\"},{\"name\":\"隐藏\",\"value\":\"0\"}]",
	}},
	{TableName: "tb_option_models", Condition: [][]interface{}{{"id", 4}}, Data: map[string]interface{}{
		"id":          4,
		"name":        "启用|禁用",
		"static_data": "[{\"name\":\"开\",\"value\":\"1\",\"color\":\"#66CC00\",\"icon\":\"ri-checkbox-circle-fill\"},{\"name\":\"关\",\"value\":\"0\",\"color\":\"#FF6666\",\"icon\":\"ri-forbid-fill\"}]",
	}},
	{TableName: "tb_option_models", Condition: [][]interface{}{{"id", 5}}, Data: map[string]interface{}{
		"id":          5,
		"name":        "启用|禁用",
		"static_data": "[{\"name\":\"菜单\",\"value\":\"1\",\"icon\":\"ri-menu-fill\",\"color\":\"#0099CC\"},{\"name\":\"操作\",\"value\":\"2\",\"icon\":\"ri-cursor-line\",\"color\":\"#CC6699\"}]",
	}},
}
