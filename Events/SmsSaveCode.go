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
	"time"
)

type SmsSaveCode struct {
}

// Do 记录短信验证码
func (that SmsSaveCode) Do(eventName string, data ...interface{}) (error, int) {
	SmsModels.Sms{}.SaveCode(data[0].(string), data[1].(string), data[2].(time.Duration))
	return nil, 200
}
