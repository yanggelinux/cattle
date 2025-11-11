package request

type GetRoleReq struct {
	RoleName *string `form:"roleName"`
	Page     *int    `form:"page"`
	PageSize *int    `form:"pageSize"`
}

type CreateRoleReq struct {
	RoleName    *string `json:"roleName" binding:"required"`
	DisplayName *string `json:"displayName" binding:"required"`
	IsSuper     *int8   `json:"isSuper" binding:"required"`
}

type UpdateRoleReq struct {
	ID          *int64  `json:"id" binding:"required"`
	RoleName    *string `json:"roleName"`
	DisplayName *string `json:"displayName"`
	IsSuper     *int8   `json:"isSuper"`
}
