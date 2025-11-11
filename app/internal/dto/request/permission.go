package request

type GetPermissionReq struct {
	Name     *string `form:"name"`
	Code     *string `form:"code"`
	Project  *string `form:"project"`
	Page     *int    `form:"page"`
	PageSize *int    `form:"pageSize"`
}

type CreatePermissionReq struct {
	ParentID  *int64  `json:"parentID" binding:"required"`
	Name      *string `json:"name" binding:"required"`
	Code      *string `json:"code" binding:"required"`
	Uri       *string `json:"uri" binding:"required"`
	Method    *string `json:"method" binding:"required"`
	Project   *string `json:"project" binding:"required"`
	PermType  *int8   `json:"permType" binding:"required"`
	IsEnabled *int8   `json:"isEnabled" binding:"required"`
	Sort      *int64  `json:"sort" binding:"required"`
}

type UpdatePermissionReq struct {
	ID        *int64  `json:"id" binding:"required"`
	ParentID  *int64  `json:"parentID"`
	Name      *string `json:"name"`
	Code      *string `json:"code"`
	Uri       *string `json:"uri"`
	Method    *string `json:"method"`
	Project   *string `json:"project"`
	PermType  *int8   `json:"permType"`
	IsEnabled *int8   `json:"isEnabled"`
	Sort      *int64  `json:"sort"`
}

type GetRolePermReq struct {
	RoleID  *int64  `form:"roleID" binding:"required"`
	Project *string `form:"project"  binding:"required"`
	IsSuper *int8   `form:"isSuper"`
}

type UpdateRolePermReq struct {
	RoleID     *int64  `form:"roleID" binding:"required"`
	Project    *string `form:"project"  binding:"required"`
	PermIDList []int64 `form:"permIDList" binding:"required"`
}
