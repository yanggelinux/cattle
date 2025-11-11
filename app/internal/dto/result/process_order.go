package result

import "github.com/yanggelinux/cattle/internal/store/model"

type ProcessOrderRet struct {
	ID           int64                `json:"id"`
	OrderID      int64                `json:"orderID"`
	GraphID      int64                `json:"graphID"`
	Title        string               `json:"title"`
	GraphName    string               `json:"graphName"`
	OrderName    string               `json:"orderName"`
	DemandName   string               `json:"demandName"`
	Owner        string               `json:"owner"`
	Env          string               `json:"env"`
	OrderType    int8                 `json:"orderType"`
	OrderProcess []*model.PorcessNode `json:"orderProcess"`
	CurOrderNode *model.PorcessNode   `json:"curOrderNode"`
	HasApproval  int8                 `json:"hasApproval"`
	TaskStatus   int8                 `json:"taskStatus"`
	Status       int8                 `json:"status"`
	ActiveIndex  int8                 `json:"activeIndex"`
	UpdatedTime  string               `json:"updatedTime"`
	CreatedTime  string               `json:"createdTime"`
}

type ProcessOrderUnapprovedRet struct {
	ID          int64       `json:"id"`
	OrderID     int64       `json:"orderID"`
	GraphID     int64       `json:"graphID"`
	Title       string      `json:"title"`
	Env         string      `json:"env"`
	GraphName   string      `json:"graphName"`
	OrderName   string      `json:"orderName"`
	DemandName  string      `json:"demandName"`
	Owner       string      `json:"owner"`
	OrderType   int8        `json:"orderType"`
	Status      int8        `json:"status"`
	OrderInfo   interface{} `json:"orderInfo"`
	UpdatedTime string      `json:"updatedTime"`
	CreatedTime string      `json:"createdTime"`
}

type ProcessOrderResult struct {
	Total   int64              `json:"total"`
	RetList []*ProcessOrderRet `json:"retList"`
}

type ProcessOrderUnapprovedResult struct {
	Total   int64                        `json:"total"`
	RetList []*ProcessOrderUnapprovedRet `json:"retList"`
}

type ProcessOrderDetailResult struct {
	ProcessOrderRet
	OrderInfo        interface{} `json:"orderInfo"`
	OrderField       interface{} `json:"orderField"`
	Opinion          string      `json:"opinion"`
	ImageHash        string      `json:"imageHash"`
	EnabledImageHash string      `json:"enabledImageHash"`
	ImageData        string      `json:"imageData"`
	EnabledImageData string      `json:"enabledImageData"`
	Description      string      `json:"description"`
	OrderLabel       string      `json:"orderLabel"`
	TaskResult       interface{} `json:"taskResult"`
	OrderLayout      int8        `json:"orderLayout"`
}

type ProcessOrderFieldResult struct {
	FieldData interface{} `json:"fieldData"`
}

// 审批节点
type PorcessNode struct {
	Name          string `json:"name"`
	DeptName      string `json:"deptName"`
	ApproverGroup string `json:"approverGroup"`
	Approver      string `json:"approver"`
	ApproverName  string `json:"approverName"`
	Role          string `json:"role"`
	Opt           string `json:"opt"`
	Status        int8   `json:"status"`
}
