/**
 * @Author: 云起时
 * @Email: limingxiang@yqstech.com
 * @Description:
 * @File: database-tables
 * @Version: 1.0.0
 * @Date: 2022/8/16 22:27
 */

package sms

import "github.com/yqstech/gef"

// tables 需要维护的所有的表结构体
var tables = []interface{}{
	&TbSmsUpstream{},
	&TbSmsUpstreamParams{},
	&TbSmsTemplate{},
	&TbAppSmsUpstream{},
	&TbAppSmsTemplate{},
	&TbAppSmsRecord{},
	&TbAppSmsHoldBack{},
	&TbAppSmsBlackWhite{},
}

//自动生成的表补充参数 将下边的#去掉
//正则替换，comment参数
//NOT NULL#"(.*)`(.*)// (.*)
//NOT NULL;comment:$3"$1`$2// $3

//正则替换，补充default:''
//char(.*)\);#NOT
//char$1);default:'';NOT

// TbSmsUpstream 短信通道表
type TbSmsUpstream struct {
	gef.ID
	UpstreamName string `gorm:"column:upstream_name;type:varchar(20);default:'';NOT NULL;comment:通道名称" json:"upstream_name"` // 通道名称
	EventName    string `gorm:"column:event_name;type:varchar(50);default:'';NOT NULL;comment:关联事件" json:"event_name"`       // 关联事件
	IndexNum     int    `gorm:"column:index_num;type:int(11);default:200;NOT NULL;comment:优先级" json:"index_num"`             // 优先级
	gef.CUSD
}

func (m *TbSmsUpstream) TableName() string {
	return "tb_sms_upstream"
}

// TbSmsUpstreamParams 短信通道配置项
type TbSmsUpstreamParams struct {
	gef.ID
	UpstreamID int    `gorm:"column:upstream_id;type:int(11);default:0;NOT NULL;comment:通道ID" json:"upstream_id"`      // 通道ID
	ParamName  string `gorm:"column:param_name;type:varchar(30);default:'';NOT NULL;comment:参数名称" json:"param_name"`   // 参数名称
	ParamTitle string `gorm:"column:param_title;type:varchar(30);default:'';NOT NULL;comment:参数标题" json:"param_title"` // 参数标题
	gef.CUSD
}

func (m *TbSmsUpstreamParams) TableName() string {
	return "tb_sms_upstream_params"
}

// TbSmsTemplate 短信模板表
type TbSmsTemplate struct {
	gef.ID
	TemplateName   string `gorm:"column:template_name;type:varchar(30);default:'';NOT NULL;comment:模板标识符" json:"template_name"`       // 模板标识符
	TemplateTitle  string `gorm:"column:template_title;type:varchar(30);default:'';NOT NULL;comment:模板名称" json:"template_title"`      // 模板名称
	TemplateVars   string `gorm:"column:template_vars;type:varchar(200);default:'';NOT NULL;comment:模板支持变量" json:"template_vars"`     // 模板支持变量
	DefaultContent string `gorm:"column:default_content;type:varchar(200);default:'';NOT NULL;comment:默认模板内容" json:"default_content"` // 默认模板内容
	gef.CUSD
}

func (m *TbSmsTemplate) TableName() string {
	return "tb_sms_template"
}

// TbAppSmsUpstream 应用短信通道列表
type TbAppSmsUpstream struct {
	gef.ID
	UpstreamID int    `gorm:"column:upstream_id;type:int(11);default:0;NOT NULL" json:"upstream_id"`
	Configs    string `gorm:"column:configs;type:varchar(5000);default:{};NOT NULL;comment:应用拓展配置项" json:"configs"` // 应用拓展配置项
	IndexNum   int    `gorm:"column:index_num;type:int(11);default:200;NOT NULL;comment:排序" json:"index_num"`       // 排序
	gef.CUSD
}

func (m *TbAppSmsUpstream) TableName() string {
	return "tb_app_sms_upstream"
}

// TbAppSmsTemplate 应用短信模板
type TbAppSmsTemplate struct {
	gef.ID
	TemplateName    string `gorm:"column:template_name;type:varchar(30);default:'';NOT NULL;comment:模板标识" json:"template_name"`          // 模板标识
	TemplateOutID   string `gorm:"column:template_out_id;type:varchar(30);default:'';NOT NULL;comment:外部模板ID" json:"template_out_id"`    // 外部模板ID
	TemplateContent string `gorm:"column:template_content;type:varchar(200);default:'';NOT NULL;comment:短信模板内容" json:"template_content"` // 短信模板内容
	gef.CUSD
}

func (m *TbAppSmsTemplate) TableName() string {
	return "tb_app_sms_template"
}

// TbAppSmsRecord 应用短信记录
type TbAppSmsRecord struct {
	gef.ID
	TemplateName  string `gorm:"column:template_name;type:varchar(30);default:'';NOT NULL;comment:模板标识" json:"template_name"`       // 模板标识
	UpstreamID    int    `gorm:"column:upstream_id;type:int(11);default:0;NOT NULL;comment:短信通道" json:"upstream_id"`                // 短信通道
	Tel           string `gorm:"column:tel;type:varchar(20);default:'';NOT NULL;comment:手机号" json:"tel"`                            // 手机号
	IP            string `gorm:"column:ip;type:varchar(20);default:'';NOT NULL;comment:ip地址" json:"ip"`                             // ip地址
	TemplateOutID string `gorm:"column:template_out_id;type:varchar(30);default:'';NOT NULL;comment:外置模板ID" json:"template_out_id"` // 外置模板ID
	Content       string `gorm:"column:content;type:varchar(1024);default:'';NOT NULL;comment:短信内容" json:"content"`                 // 短信内容
	Msg           string `gorm:"column:msg;type:varchar(50);default:'';NOT NULL;comment:消息" json:"msg"`                             // 消息
	gef.CUSD
}

func (m *TbAppSmsRecord) TableName() string {
	return "tb_app_sms_record"
}

// TbAppSmsHoldBack 应用短信防火墙
type TbAppSmsHoldBack struct {
	gef.ID
	RuleType     int    `gorm:"column:rule_type;type:tinyint(4);default:0;NOT NULL;comment:0单个手机号 1单个IP地址 2短信频率" json:"rule_type"` // 0单个手机号 1单个IP地址 2短信频率
	RangeSecond  int    `gorm:"column:range_second;type:int(11);default:0;NOT NULL;comment:多少秒内" json:"range_second"`              // 多少秒内
	SmsMax       int    `gorm:"column:sms_max;type:int(11);default:0;NOT NULL;comment:短信次数" json:"sms_max"`                        // 短信次数
	Action       int    `gorm:"column:action;type:int(11);default:0;NOT NULL;comment:执行操作，1临时冻结，2拉黑" json:"action"`                // 执行操作，1临时冻结，2拉黑
	FrozenSecond int    `gorm:"column:frozen_second;type:int(11);default:0;NOT NULL;comment:冻结时间" json:"frozen_second"`            // 冻结时间
	Note         string `gorm:"column:note;type:varchar(30);default:'';NOT NULL;comment:备注" json:"note"`                           // 备注
	gef.CUSD
}

func (m *TbAppSmsHoldBack) TableName() string {
	return "tb_app_sms_hold_back"
}

// TbAppSmsBlackWhite 短信黑白名单
type TbAppSmsBlackWhite struct {
	gef.ID
	Type     int    `gorm:"column:type;type:tinyint(4);default:0;NOT NULL;comment:0黑名单 1白名单" json:"type"`            // 0黑名单 1白名单
	RuleType int    `gorm:"column:rule_type;type:tinyint(4);default:0;NOT NULL;comment:0手机号 1IP地址" json:"rule_type"` // 0手机号 1IP地址
	Rule     string `gorm:"column:rule;type:varchar(30);default:'';NOT NULL;comment:手机号或ip" json:"rule"`             // 手机号或ip
	Note     string `gorm:"column:note;type:varchar(30);default:'';NOT NULL;comment:备注" json:"note"`                 // 备注
	gef.CUSD
}

func (m *TbAppSmsBlackWhite) TableName() string {
	return "tb_app_sms_black_white"
}
