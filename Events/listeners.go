/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: Listeners
 * @Version: 1.0.0
 * @Date: 2021/11/24 3:11 下午
 */

package Events

import "github.com/yqstech/gef/event"

// Listeners 事件监听列表
var Listeners = map[string][]event.Listener{
	//短信防火墙
	"SmsHoldBack": {
		SmsHoldBack{},
	},
	//短信记录验证码
	"SmsSaveCode": {
		SmsSaveCode{},
	},
	//短信校验验证码
	"SmsCheckCode": {
		SmsCheckCode{},
	},
	//短信渲染并发送
	"SmsDisplayAndSend": []event.Listener{
		SmsDisplayAndSend{},
	},
	//发送短信
	"SmsSend": []event.Listener{
		//#Map tel(string) ip(string) content(string) template_out_id(string)
		SmsSend{}, //发送短信
	},
	//发送短信通道
	"SmsAli": []event.Listener{
		SmsAli{}, //阿里短信
	},
	"SmsAm": []event.Listener{
		SmsAm{}, //云市场
	},
	"SmsJdcx": []event.Listener{
		SmsJdcx{}, //京东万象
	},
	"SmsMock": []event.Listener{
		SmsMock{}, //模拟短信
	},
	"SmsCdcx": []event.Listener{
		SmsCdcx{}, //成都创信
	},
	"SmsZj": []event.Listener{
		SmsZj{}, //广州掌骏
	},
}
