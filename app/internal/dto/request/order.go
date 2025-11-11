package request

type GetOrderReq struct {
	Name      *string `form:"name"`
	OrderType *int8   `form:"orderType"`
	Status    *int8   `form:"status"`
	Page      *int    `form:"page"`
	PageSize  *int    `form:"pageSize"`
}

type CreateOrderReq struct {
	Name       *string  `json:"name" binding:"required"`
	GroupID    *int64   `json:"groupID" binding:"required"`
	ProcessID  *int64   `json:"processID" binding:"required"`
	OrderType  *int8    `json:"orderType" binding:"required"`
	NodeType   []string `json:"nodeType"`
	Label      *string  `json:"label"`
	Layout     *int8    `json:"layout"`
	TaskUrl    *string  `json:"taskUrl"`
	TaskMethod *string  `json:"taskMethod"`
	Sort       *int64   `json:"sort" binding:"required"`
	Status     *int8    `json:"status" binding:"required"`
}

type UpdateOrderReq struct {
	ID         *int64   `json:"id" binding:"required"`
	Name       *string  `json:"name"`
	GroupID    *int64   `json:"groupID"`
	ProcessID  *int64   `json:"processID"`
	OrderType  *int8    `json:"orderType"`
	NodeType   []string `json:"nodeType"`
	Label      *string  `json:"label"`
	Layout     *int8    `json:"layout"`
	TaskUrl    *string  `json:"taskUrl"`
	TaskMethod *string  `json:"taskMethod"`
	Sort       *int64   `json:"sort"`
	Status     *int8    `json:"status"`
}
