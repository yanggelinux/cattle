package processorder

import "github.com/yanggelinux/cattle/internal/store/model"

var GraphOrderPorcess []*model.PorcessNode = []*model.PorcessNode{
	&model.PorcessNode{
		Name:         "工单申请",
		Type:         "procStart",
		ApprovalType: 0,
		ApprovalInfo: []*model.ApprovalInfo{
			&model.ApprovalInfo{
				Approver:     "",
				ApproverName: "",
				Role:         "",
				RoleName:     "",
				Status:       1,
			},
		},
		Status: 1,
	},
	&model.PorcessNode{
		Name:         "架构负责人审批",
		Type:         "procApproval",
		ApprovalType: 0,
		ApprovalInfo: []*model.ApprovalInfo{
			&model.ApprovalInfo{
				Approver:     "",
				ApproverName: "",
				Role:         "architect",
				RoleName:     "架构师",
				Status:       0,
			},
		},
		Status: 0,
	},
	&model.PorcessNode{
		Name:         "安全负责人审批",
		Type:         "procApproval",
		ApprovalType: 0,
		ApprovalInfo: []*model.ApprovalInfo{
			&model.ApprovalInfo{
				Approver:     "",
				ApproverName: "",
				Role:         "security",
				RoleName:     "安全管理员",
				Status:       0,
			},
		},
		Status: 0,
	},
	&model.PorcessNode{
		Name:         "审批成功",
		Type:         "procEnd",
		ApprovalType: 0,
		ApprovalInfo: []*model.ApprovalInfo{
			&model.ApprovalInfo{
				Approver:     "",
				ApproverName: "",
				Role:         "",
				RoleName:     "",
				Status:       0,
			},
		},
		Status: 0,
	},
}
