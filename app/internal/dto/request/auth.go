package request

type LoginReq struct {
	UserName *string `json:"userName" binding:"required"`
	Password *string `json:"password" binding:"required"`
}

type LoginByAuthorizeReq struct {
	UserName      *string `json:"userName" binding:"required"`
	Authorization *string `json:"authorization" binding:"required"`
	Project       *string `json:"project"`
}

type GetAuthorizationReq struct {
	UserName *string `json:"userName" binding:"required"`
}

type GetTokenReq struct {
	UserName *string `json:"userName" binding:"required"`
	Password *string `json:"password"`
}

type GetUserPermReq struct {
	UserID *int64 `form:"userID" binding:"required"`
}
