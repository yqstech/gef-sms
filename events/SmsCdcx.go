/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: SmsAli
 * @Version: 1.0.0
 * @Date: 2022/7/29 23:03
 */

package events

import (
	"errors"
	"github.com/wonderivan/logger"
	"github.com/yqstech/gef/util"
)

type SmsCdcx struct {
}

func (that SmsCdcx) Do(eventName string, data ...interface{}) (error, int) {
	programs := data[0].(map[string]interface{})
	account := programs["account"].(string)
	password := programs["password"].(string)
	extno := programs["extno"].(string)
	url := programs["url"].(string)
	Content := programs["content"].(string)
	tel := programs["tel"].(string)
	sign := programs["sign"].(string)
	url = url + "?mobile=" + tel + "&content=【" + sign + "】" + Content +
		"&account=" + account +
		"&password=" + password +
		"&extno=" + extno + "&action=send&rt=json"
	content, err := util.FastHttpGet(url)
	if err != nil {
		logger.Error(err.Error())
		return err, 500
	}
	resultInfo := map[string]interface{}{}
	util.JsonDecode(content, &resultInfo)

	if util.Interface2String(resultInfo["status"]) != "0" {
		return errors.New("status=" + util.Interface2String(resultInfo["status"])), 201
	}
	list := resultInfo["list"].([]map[string]interface{})
	if len(list) == 0 {
		return errors.New("反馈失败"), 201
	}
	if util.Interface2String(list[0]["result"]) == "0" {
		return nil, 200
	} else {
		return errors.New("result=" + util.Interface2String(list[0]["result"])), 202
	}
}
