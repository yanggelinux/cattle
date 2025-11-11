package request

type GetArchGroupReq struct {
	ParentID  *int64  `form:"parentID"`
	GroupName *string `form:"groupName"`
}

type CreateArchGroupReq struct {
	ParentID  *int64  `json:"parentID" binding:"required"`
	GroupName *string `json:"groupName" binding:"required"`
}

type UpdateArchGroupReq struct {
	ID        *int64  `json:"id" binding:"required"`
	ParentID  *int64  `json:"parentID"`
	GroupName *string `json:"groupName"`
}
