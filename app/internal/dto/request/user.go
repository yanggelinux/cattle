package request

type GetUserReq struct {
	UserName *string `form:"userName"`
	Email    *string `form:"email"`
	Page     *int    `form:"page"`
	PageSize *int    `form:"pageSize"`
}

type CreateUserReq struct {
	UserName    *string `json:"userName" binding:"required"`
	Password    *string `json:"password" binding:"required"`
	DisplayName *string `json:"displayName" binding:"required"`
	Email       *string `json:"email" binding:"required"`
	DeptName    *string `json:"deptName" binding:"required"`
	RoleIDs     []int64 `json:"roleIDs" binding:"required"`
}

type UpdateUserReq struct {
	ID          *int64  `json:"id" binding:"required"`
	UserName    *string `json:"userName"`
	Password    *string `json:"password"`
	DisplayName *string `json:"displayName"`
	Email       *string `json:"email"`
	DeptName    *string `json:"deptName"`
	RoleIDs     []int64 `json:"roleIDs"`
	Origin      *int8   `json:"origin"`
}
