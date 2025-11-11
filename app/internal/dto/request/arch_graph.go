package request

import "gorm.io/datatypes"

type GetArchGraphReq struct {
	GroupID   *int64  `form:"groupID"`
	GraphName *string `form:"graphName"`
	Status    *int8   `form:"status"`
	Page      *int    `form:"page"`
	PageSize  *int    `form:"pageSize"`
}

type GetArchGraphRecordReq struct {
	GraphID *int64 `form:"graphID"`
}

type CreateArchGraphReq struct {
	GroupID    *int64  `json:"groupID" binding:"required"`
	GraphName  *string `json:"graphName" binding:"required"`
	GraphLabel *string `json:"graphLabel"`
	Owner      *string `json:"owner" binding:"required"`
	Status     *int8   `json:"status"`
}

type SelectArchGraphReq struct {
	ID      *int64 `json:"id" binding:"required"`
	GraphID *int64 `json:"graphID" binding:"required"`
}

type UpdateArchGraphReq struct {
	ID         *int64         `json:"id" binding:"required"`
	GroupID    *int64         `json:"groupID"`
	GraphName  *string        `json:"graphName"`
	GraphLabel *string        `json:"graphLabel"`
	NodeData   datatypes.JSON `json:"nodeData"`
	EdgeData   datatypes.JSON `json:"edgeData"`
	ImageData  *string        `json:"imageData"`
	Owner      *string        `json:"owner"`
	IsShared   *int8          `json:"isShared"`
	Status     *int8          `json:"status"`
	Action     *string        `json:"action"`
}

type GetArchGraphReviewReq struct {
	GraphID  *int64  `form:"graphID"`
	GraphKey *string `form:"graphKey"`
	Page     *int    `form:"page"`
	PageSize *int    `form:"pageSize"`
}
type CreateArchGraphReviewReq struct {
	GraphID     *int64   `json:"graphID" binding:"required"`
	GraphKey    *string  `json:"graphKey" binding:"required"`
	Content     *string  `json:"content" binding:"required"`
	GraphName   *string  `json:"graphName"`
	NotifyParty []string `json:"notifyParty"`
}

type GetArchGraphDataReq struct {
	GraphLabel *string `form:"graphLabel"`
}
