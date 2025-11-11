package request

type GetOrderFieldReq struct {
	Name     *string `form:"name"`
	OrderID  *int64  `form:"orderID"`
	Status   *int8   `form:"status"`
	Page     *int    `form:"page"`
	PageSize *int    `form:"pageSize"`
}

type CreateOrderFieldReq struct {
	OrderID      *int64  `json:"orderID" binding:"required"`
	Name         *string `json:"name" binding:"required"`
	Key          *string `json:"key" binding:"required"`
	Component    *string `json:"component" binding:"required"`
	Placeholder  *string `json:"placeholder"`
	VerRule      *int8   `json:"verRule" binding:"required"`
	DefaultVal   *string `json:"defaultVal"`
	IsRequired   *int8   `json:"isRequired" binding:"required"`
	IsTitle      *int8   `json:"isTitle"`
	IsEdit       *int8   `json:"isEdit"`
	IsClear      *int8   `json:"isClear"`
	DisplayField *string `json:"displayField"`
	DisplayVal   *string `json:"displayVal"`
	Description  *string `json:"description"`
	Enum         *string `json:"enum"`
	GroupName    *string `json:"groupName"`
	Sort         *int64  `json:"sort" binding:"required"`
	Status       *int8   `json:"status"`
}

type UpdateOrderFieldReq struct {
	ID           *int64  `json:"id" binding:"required"`
	OrderID      *int64  `json:"orderID"`
	Name         *string `json:"name"`
	Key          *string `json:"key"`
	Component    *string `json:"component"`
	Placeholder  *string `json:"placeholder"`
	VerRule      *int8   `json:"verRule"`
	DefaultVal   *string `json:"defaultVal"`
	IsRequired   *int8   `json:"isRequired"`
	IsTitle      *int8   `json:"isTitle"`
	IsEdit       *int8   `json:"isEdit"`
	IsClear      *int8   `json:"isClear"`
	DisplayField *string `json:"displayField"`
	DisplayVal   *string `json:"displayVal"`
	Description  *string `json:"description"`
	Enum         *string `json:"enum"`
	GroupName    *string `json:"groupName"`
	Sort         *int64  `json:"sort"`
	Status       *int8   `json:"status"`
}
