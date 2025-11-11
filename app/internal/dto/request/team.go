package request

type GetTeamReq struct {
	Name     *string `form:"name"`
	Page     *int    `form:"page"`
	PageSize *int    `form:"pageSize"`
}

type CreateTeamReq struct {
	Name     *string `json:"name" binding:"required"`
	Leader   *string `json:"leader" binding:"required"`
	Director *string `json:"director" binding:"required"`
}

type UpdateTeamReq struct {
	ID       *int64  `json:"id" binding:"required"`
	Name     *string `json:"name"`
	Leader   *string `json:"leader"`
	Director *string `json:"director"`
}
