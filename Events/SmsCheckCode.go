/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: SmsSaveCode
 * @Version: 1.0.0
 * @Date: 2022/8/17 11:18
 */

package Events

import (
	"github.com/yqstech/gef-sms/SmsModels"
)

type SmsCheckCode struct {
}

// Do 记录短信验证码
func (that SmsCheckCode) Do(eventName string, data ...interface{}) (error, int) {
	ok := SmsModels.Sms{}.CheckCode(data[0].(string), data[1].(string))
	if ok {
		return nil, 200
	}
	return nil, 0
}
