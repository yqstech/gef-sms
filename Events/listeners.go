/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: Listeners
 * @Version: 1.0.0
 * @Date: 2021/11/24 3:11 下午
 */

package Events

import "github.com/yqstech/gef/Event"

// Listeners 事件监听列表
var Listeners = map[string][]Event.Listener{
	//短信记录验证码
	"SmsSaveCode": {
		SmsSaveCode{},
	},
	//短信校验验证码
	"SmsCheckCode": {
		SmsCheckCode{},
	},
	//短信渲染并发送
	"SmsDisplayAndSend": []Event.Listener{},
	//发送短信
	"SmsSend": []Event.Listener{
		//#Map tel(string) ip(string) content(string) template_out_id(string)
		SmsSend{}, //发送短信
	},
	//发送短信通道
	"SmsAli": []Event.Listener{
		SmsAli{}, //阿里短信
	},
	"SmsAm": []Event.Listener{
		SmsAm{}, //云市场
	},
	"SmsJdcx": []Event.Listener{
		SmsJdcx{}, //京东万象
	},
	"SmsMock": []Event.Listener{
		SmsMock{}, //模拟短信
	},
}
