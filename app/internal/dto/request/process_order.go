package request

import "gorm.io/datatypes"

type GetProcessOrderReq struct {
	GraphName  *string `form:"graphName"`
	DemandName *string `form:"demandName"`
	Approver   *string `form:"approver"`
	Owner      *string `form:"owner"`
	Env        *string `form:"env"`
	Title      *string `form:"title"`
	OrderType  *int8   `form:"orderType"`
	Status     *int8   `form:"status"`
	StartTime  *string `form:"startTime"`
	EndTime    *string `form:"endTime"`
	Label      *string `form:"label"`
	Page       *int    `form:"page"`
	PageSize   *int    `form:"pageSize"`
}

type CreateProcessOrderReq struct {
	OrderID          *int64         `json:"orderID"`
	GraphID          *int64         `json:"graphID"`
	Title            *string        `json:"title" binding:"required"`
	Env              *string        `json:"env"`
	GraphName        *string        `json:"graphName"`
	OrderName        *string        `json:"orderName"`
	OrderLabel       *string        `json:"orderLabel"`
	DemandName       *string        `json:"demandName"`
	Owner            *string        `json:"owner" binding:"required"`
	OrderType        *int8          `json:"orderType" binding:"required"`
	ImageData        *string        `json:"imageData"`
	EnabledImageData *string        `json:"enabledImageData"`
	OrderInfo        datatypes.JSON `json:"orderInfo"`
	OrderField       datatypes.JSON `json:"orderField"`
	Description      *string        `json:"description"`
}

type ApplyProcessOrderReq struct {
	IDs              []int64        `json:"ids"`
	OrderID          *int64         `json:"orderID"`
	GraphID          *int64         `json:"graphID"`
	Title            *string        `json:"title"`
	Env              *string        `json:"env"`
	GraphName        *string        `json:"graphName"`
	OrderName        *string        `json:"orderName"`
	Owner            *string        `json:"owner"`
	ImageData        *string        `json:"imageData"`
	EnabledImageData *string        `json:"enabledImageData"`
	DemandName       *string        `json:"demandName"`
	OrderType        *int8          `json:"orderType"`
	OrderInfo        datatypes.JSON `json:"orderInfo"`
	OrderField       datatypes.JSON `json:"orderField"`
	Description      *string        `json:"description"`
}

type ReApplyProcessOrderReq struct {
	// 资源申请工单相应的id
	ID *int64 `json:"id" binding:"required"`
}

type UpdateProcessOrderReq struct {
	ID         *int64         `json:"id" binding:"required"`
	Env        *string        `json:"env"`
	DemandName *string        `json:"demandName"`
	OrderLabel *string        `json:"orderLabel"`
	IsApproval *int8          `json:"isApproval"`
	OrderType  *int8          `json:"orderType" binding:"required"`
	OrderInfo  datatypes.JSON `json:"orderInfo"`
}

type ApproveProcessOrderReq struct {
	ID           *int64  `json:"id" binding:"required"`
	GraphID      *int64  `json:"graphID" binding:"required"`
	Action       *string `json:"action" binding:"required"`
	ProcNodeName *string `json:"procNodeName"`
	OrderLabel   *string `json:"orderLabel"`
	Opinion      *string `json:"opinion"`
	Approver     *string `json:"approver" binding:"required"`
	ApproverName *string `json:"approverName" binding:"required"`
}

type AssignApproverReq struct {
	ID       *int64  `json:"id" binding:"required"`
	Approver *string `json:"approver" binding:"required"`
}

type GetProcessOrderFieldReq struct {
	OrderType *int8 `form:"orderType" binding:"required"`
}

type GetUnapprovedOrderReq struct {
	GraphID *int64  `form:"graphID"`
	Title   *string `form:"title"`
}
