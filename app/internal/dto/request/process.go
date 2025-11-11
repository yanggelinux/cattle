package request

import "gorm.io/datatypes"

type GetProcessReq struct {
	Name     *string `form:"name"`
	Status   *int8   `form:"status"`
	Page     *int    `form:"page"`
	PageSize *int    `form:"pageSize"`
}

type CreateProcessReq struct {
	Name     *string        `json:"name" binding:"required"`
	NodeData datatypes.JSON `json:"nodeData"`
	EdgeData datatypes.JSON `json:"edgeData"`
	Status   *int8          `json:"status"`
}

type UpdateProcessReq struct {
	ID       *int64         `json:"id" binding:"required"`
	Name     *string        `json:"name"`
	NodeData datatypes.JSON `json:"nodeData"`
	EdgeData datatypes.JSON `json:"edgeData"`
	Status   *int8          `json:"status"`
}
