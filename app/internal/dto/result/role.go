package result

type RoleRet struct {
	ID          int64  `json:"id"`
	RoleName    string `json:"roleName"`
	DisplayName string `json:"displayName"`
	IsSuper     int8   `json:"isSuper"`
	UpdatedTime string `json:"updatedTime"`
	CreatedTime string `json:"createdTime"`
}

type RoleResult struct {
	Total   int64      `json:"total"`
	RetList []*RoleRet `json:"retList"`
}
