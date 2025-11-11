package demand

import "github.com/yanggelinux/cattle/internal/dto/result"

type ApprovalInfo struct {
	BizGroup     string
	ApproverName string
	Approver     string
}

var approvalMap = map[string]ApprovalInfo{
	"电子社保卡后端": {BizGroup: "电子社保卡后端", ApproverName: "朱荣照", Approver: "zhurongzhao"},
	"移动支付":    {BizGroup: "移动支付", ApproverName: "朱荣照", Approver: "zhurongzhao"},
	"电子社保卡前端": {BizGroup: "电子社保卡前端", ApproverName: "程千浪", Approver: "chengqianlang"},
	"就业在线后端":  {BizGroup: "就业在线后端", ApproverName: "钱晓鹏", Approver: "qianxiaopeng"},
	"就业在线前端":  {BizGroup: "就业在线前端", ApproverName: "钱晓鹏", Approver: "qianxiaopeng"},
	"HR云":     {BizGroup: "HR云", ApproverName: "崔康", Approver: "cuikang"},
	"大数据应用":   {BizGroup: "大数据应用", ApproverName: "冯耀东", Approver: "fengyaodong"},
	"大数据平台":   {BizGroup: "大数据平台", ApproverName: "贾现强", Approver: "jiaxianqiang"},
	"算法开发":    {BizGroup: "算法开发", ApproverName: "刘金超", Approver: "liujinchao"},
	"基础平台":    {BizGroup: "基础平台", ApproverName: "刘峰", Approver: "liufeng"},
	"架构组":     {BizGroup: "架构组", ApproverName: "赵华峰", Approver: "zhaohuafeng"},
	"民生业务测试":  {BizGroup: "民生业务测试", ApproverName: "周璐", Approver: "zhoulu"},
	"创新业务测试":  {BizGroup: "创新业务测试", ApproverName: "张艳华", Approver: "zhangyanhua"},
	"质量":      {BizGroup: "质量", ApproverName: "朱海娇", Approver: "zhuhaijiao"},
	"应用运维":    {BizGroup: "应用运维", ApproverName: "张善慈", Approver: "zhangshanci"},
	"资源管理":    {BizGroup: "资源管理", ApproverName: "张磊", Approver: "zhanglei"},
	"安全运营":    {BizGroup: "安全运营", ApproverName: "邓杰汉", Approver: "dengjiehan"},
	"安全服务":    {BizGroup: "安全服务", ApproverName: "王欣欣", Approver: "wangxinxin"},
	"安全支撑":    {BizGroup: "安全支撑", ApproverName: "李龙吉", Approver: "lilongji"},
	"渠道运营":    {BizGroup: "渠道运营", ApproverName: "光泽", Approver: "guangze"},
	"业务支持":    {BizGroup: "业务支持", ApproverName: "赵悦", Approver: "zhaoyue"},
}

var ReviewProcess []*result.PorcessNode = []*result.PorcessNode{
	&result.PorcessNode{
		Name:          "评审申请",
		ApproverGroup: "",
		Approver:      "",
		ApproverName:  "",
		Role:          "",
		Opt:           "apply",
		Status:        1,
	},
	&result.PorcessNode{
		Name:          "业务组负责人审批",
		ApproverGroup: "业务组",
		Approver:      "",
		ApproverName:  "",
		// 负责审批的角色
		Role:   "teamLeader",
		Opt:    "approve",
		Status: 0,
	},
	//&result.PorcessNode{
	//	Name:          "架构负责人审批",
	//	ApproverGroup: "架构团队",
	//	Approver:      "",
	//	ApproverName:  "",
	//	// 负责审批的角色
	//	Role:   "architect",
	//	Opt:    "approve",
	//	Status: 0,
	//},
	&result.PorcessNode{
		Name:          "运维负责人审批",
		ApproverGroup: "运维团队",
		Approver:      "",
		ApproverName:  "",
		Role:          "ops",
		Opt:           "approve",
		Status:        0,
	},
	&result.PorcessNode{
		Name:          "评审成功",
		ApproverGroup: "",
		Approver:      "",
		ApproverName:  "",
		Role:          "",
		Opt:           "complete",
		Status:        0,
	},
}
