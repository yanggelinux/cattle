package model

import "gorm.io/datatypes"

func NewProcess() *Process {
	return &Process{Model: &Model{}}
}

type Process struct {
	ID       int64          `json:"id" gorm:"column:id"`
	Name     string         `json:"name" gorm:"column:name" `
	ProcInfo datatypes.JSON `json:"ProcInfo" gorm:"column:proc_info" type:"json"`
	NodeData datatypes.JSON `json:"nodeData" gorm:"column:node_data" type:"json"`
	EdgeData datatypes.JSON `json:"edgeData" gorm:"column:edge_data" type:"json"`
	Status   int8           `json:"status" gorm:"column:status"`
	*Model
}

func (m *Process) TableName() string {
	return "process"
}

func (m *Process) IDField() string {
	return "id"
}
func (m *Process) NameField() string {
	return "name"
}
func (m *Process) ProcInfoField() string {
	return "proc_info"
}
func (m *Process) NodeDataField() string {
	return "node_data"
}
func (m *Process) EdgeDataField() string {
	return "edge_data"
}
func (m *Process) StatusField() string {
	return "status"
}

type PorcessNode struct {
	Name     string `json:"name"`
	DeptName string `json:"deptName"`
	// 节点类型 procStart procEnd procApproval
	Type string `json:"type"`
	// 审批人，默认时候为空，只有在指定审批人的时候在生效
	Approver string `json:"approver"`
	// 0 一个子节点通过就全通过，1 所有的通过才通过
	ApprovalType int8 `json:"approvalType"`
	// 0 审批时不可编辑，1 审批时编辑
	ApprovalEdit int8 `json:"approvalEdit"`
	// 节点的审批信息
	ApprovalInfo []*ApprovalInfo `json:"approvalInfo"`
	// 总的状态
	Status int8 `json:"status"`
}
type ApprovalInfo struct {
	Approver     string `json:"approver"`
	ApproverName string `json:"approverName"`
	Role         string `json:"role"`
	RoleName     string `json:"roleName"`
	// 审批节点状态
	Status int8 `json:"status"`
}
