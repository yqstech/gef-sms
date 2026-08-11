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
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/wonderivan/logger"
	"github.com/yqstech/gef/util"
	"io"
	"net/http"
	"strconv"
	"time"
)

type SmsZj struct {
}

func (that SmsZj) Do(eventName string, data ...interface{}) (error, int) {
	programs := data[0].(map[string]interface{})
	Content := programs["content"].(string) //短信内容
	tel := programs["tel"].(string)         //手机号码
	//配置参数
	sign := programs["sign"].(string)         //短信签名
	account := programs["account"].(string)   //发短信账号
	password := programs["password"].(string) //发短信密码
	url := programs["url"].(string)           //发短信接口地址
	// 获取当前时间戳，精确到毫秒
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	// 定义请求数据
	postData := map[string]interface{}{
		"userName":  account,                                    //发短信账号
		"content":   "【" + sign + "】" + Content,                 //短信内容
		"phoneList": []string{tel},                              //手机号列表
		"timestamp": timestamp,                                  //时间戳
		"sign":      that.getSign(timestamp, account, password), //接口签名
	}
	// 发送 HTTP 请求
	resultJson := that.httpRequest(url, postData)
	result := map[string]any{}
	util.JsonDecode(resultJson, &result)
	if code, ok := result["code"]; ok {
		if util.Any2Int64(code) == int64(0) {
			//成功
			return nil, 200
		} else {
			//失败
			return errors.New(result["message"].(string) + "；code=" + util.Interface2String(result["code"])), 201
		}
	}
	//异常
	logger.Error("短信接口返回数据异常:" + resultJson)
	logger.Error("请求地址:" + url)
	logger.Error("请求数据:" + util.JsonEncode(postData))
	return errors.New("短信返回数据异常，请稍后再试！"), 201
}

// 计算 sign 参数值
func (that SmsZj) getSign(timestamp int64, account, password string) string {
	data := account + strconv.FormatInt(timestamp, 10) + that.getMD5Hash(password)
	return that.getMD5Hash(data)
}

// 计算 MD5 值
func (that SmsZj) getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// 发送 HTTP 请求
func (that SmsZj) httpRequest(url string, data map[string]interface{}) string {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer func() { _ = resp.Body.Close() }()

	body, _ := io.ReadAll(resp.Body)
	return string(body)
}
