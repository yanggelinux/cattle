package request

type GetOrderGroupReq struct {
	Name     *string `form:"name"`
	Status   *int8   `form:"status"`
	Page     *int    `form:"page"`
	PageSize *int    `form:"pageSize"`
}

type CreateOrderGroupReq struct {
	Name   *string `json:"name" binding:"required"`
	Sort   *int64  `json:"sort" binding:"required"`
	Status *int8   `json:"status" binding:"required"`
}

type UpdateOrderGroupReq struct {
	ID     *int64  `json:"id" binding:"required"`
	Name   *string `json:"name"`
	Sort   *int64  `json:"sort"`
	Status *int8   `json:"status"`
}
