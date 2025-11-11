package result

type UserRet struct {
	ID            int64   `json:"id"`
	UserName      string  `json:"userName"`
	Password      string  `json:"password"`
	Email         string  `json:"email"`
	DisplayName   string  `json:"displayName"`
	DeptName      string  `json:"deptName"`
	RoleNames     string  `json:"roleNames"`
	RoleIDs       []int64 `json:"roleIDs"`
	Origin        int8    `json:"origin"`
	LastLoginTime string  `json:"lastLoginTime"`
	UpdatedTime   string  `json:"updatedTime"`
	CreatedTime   string  `json:"createdTime"`
}

type UserRoleRet struct {
	UserID      int64  `json:"userID"`
	RoleID      int64  `json:"roleID"`
	RoleName    string `json:"roleName"`
	DisplayName string `json:"displayName"`
	IsSuper     int8   `json:"isSuper"`
}

type UserResult struct {
	Total   int64      `json:"total"`
	RetList []*UserRet `json:"retList"`
}
