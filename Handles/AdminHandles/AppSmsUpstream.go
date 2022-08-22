/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: ManagerSmsUpstream
 * @Version: 1.0.0
 * @Date: 2021/10/29 11:19 上午
 */

package AdminHandles

import (
	"errors"
	"github.com/gohouse/gorose/v2"
	"github.com/wonderivan/logger"
	"github.com/yqstech/gef/Handles/adminHandle"
	"github.com/yqstech/gef/Utils/db"
	"github.com/yqstech/gef/builder"
	"github.com/yqstech/gef/util"
)

type AppSmsUpstream struct {
	adminHandle.Base
}

//有效通道ID列表
var upstreamIds []int64

// NodeBegin 开始
func (that AppSmsUpstream) NodeBegin(pageBuilder *builder.PageBuilder) (error, int) {
	//同步通道信息
	upstreamList, err := db.New().Table("tb_sms_upstream").
		Where("is_delete", 0).
		Where("status", "1").Get()
	if err != nil {
		logger.Error(err.Error())
		return errors.New("查询信息出错！"), 0
	}
	for _, upstreamInfo := range upstreamList {
		upstreamIds = append(upstreamIds, upstreamInfo["id"].(int64))
		ManagerUpstream, err := db.New().Table("tb_app_sms_upstream").
			Where("upstream_id", upstreamInfo["id"]).
			Where("is_delete", 0).
			First()
		if err != nil {
			logger.Error(err.Error())
			return errors.New("查询信息出错！"), 0
		}
		if ManagerUpstream == nil {
			db.New().Table("tb_app_sms_upstream").Insert(map[string]interface{}{
				"upstream_id": upstreamInfo["id"],
				"configs":     "{}",
				"create_time": util.TimeNow(),
				"update_time": util.TimeNow(),
				"status":      0,
			})
		}
	}

	//
	pageBuilder.SetTitle("短信通道管理")
	pageBuilder.SetPageName("短信通道")
	pageBuilder.SetTbName("tb_app_sms_upstream")
	pageBuilder.SetPageNotice("配置并开启多个通道，发送短信时会按照优先级(值小的优先级高)选取第一个通道发送短信，失败会启用下一个通道重新发送，直到发送成功或无可用通道！")
	return nil, 0
}

// NodeList 初始化列表
func (that AppSmsUpstream) NodeList(pageBuilder *builder.PageBuilder) (error, int) {
	pageBuilder.ListColumnClear()
	//清除列表顶部和右侧按钮
	pageBuilder.ListRightBtnsClear()
	pageBuilder.ListTopBtnsClear()
	//重置右侧按钮
	pageBuilder.SetListRightBtns("edit")
	//排序
	pageBuilder.SetListOrder("index_num,id asc")

	//获取列表
	upstreamOptions := that.SmsUpstreamList()

	pageBuilder.SetButton("edit", builder.Button{
		ButtonName: "通道设置",
		Action:     "edit",
		ActionType: 2,
		ActionUrl:  "edit",
		Class:      "layui-btn-normal",
		Icon:       "ri-settings-4-line",
		Display:    "!item.btn_edit || item.btn_edit!='hide'",
		Expand: map[string]string{
			"w": "98%",
			"h": "76%",
		},
	})

	pageBuilder.ListColumnAdd("upstream_id", "短信通道", "array", upstreamOptions)
	pageBuilder.ListColumnAdd("index_num", "优先级", "text", nil)
	pageBuilder.ListColumnAdd("status", "状态", "switch", nil)
	return nil, 0
}

// NodeListCondition 修改查询条件
func (that AppSmsUpstream) NodeListCondition(pageBuilder *builder.PageBuilder, condition [][]interface{}) ([][]interface{}, error, int) {
	//追加查询条件
	condition = append(condition, []interface{}{
		"upstream_id", "in", upstreamIds,
	})
	return condition, nil, 0
}

// NodeForm 初始化表单
func (that AppSmsUpstream) NodeForm(pageBuilder *builder.PageBuilder, id int64) (error, int) {
	//查询通道
	if id <= 0 {
		return errors.New("获取通道ID失败！"), 0
	}
	pageBuilder.FormFieldsAdd("index_num", "text", "通道优先级", "值越小越优先", "200", true, nil, "", nil)
	//查询原通道信息
	managerUpstream, err := db.New().Table("tb_app_sms_upstream").
		Where("id", id).
		Where("is_delete", 0).
		First()
	if err != nil {
		return err, 0
	}
	if managerUpstream == nil {
		return errors.New("获取通道信息失败！"), 0
	}
	// 查询通道配置项
	UpstreamParams, err := db.New().Table("tb_sms_upstream_params").
		Where("upstream_id", managerUpstream["upstream_id"]).
		Where("is_delete", 0).
		Where("status", 1).Get()
	if err != nil {
		return err, 0
	}
	configs := map[string]interface{}{}
	util.JsonDecode(managerUpstream["configs"].(string), &configs)

	for _, param := range UpstreamParams {
		value := ""
		if v, ok := configs[param["param_name"].(string)]; ok {
			value = v.(string)
		}
		pageBuilder.FormFieldsAdd(param["param_name"].(string), "text", param["param_title"].(string), "", value, true, nil, "",
			nil)
	}
	return nil, 0
}

// NodeSaveData 表单保存数据前使用
func (that AppSmsUpstream) NodeSaveData(pageBuilder *builder.PageBuilder, oldData gorose.Data, postData map[string]interface{}) (map[string]interface{}, error, int) {
	indexNum := postData["index_num"]
	delete(postData, "index_num")
	RealData := map[string]interface{}{
		"index_num": indexNum,
		"configs":   util.JsonEncode(postData),
	}
	return RealData, nil, 0
}
