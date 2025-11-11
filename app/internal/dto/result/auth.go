package result

type LoginResult struct {
	UserID        int64  `json:"userID"`
	UserName      string `json:"userName"`
	Email         string `json:"email"`
	DisplayName   string `json:"displayName"`
	DeptName      string `json:"deptName"`
	Token         string `json:"token"`
	Authorization string `json:"authorization"`
	Project       string `json:"project"`
	*RolePermResult
}

type RolePermResult struct {
	IsSuper          int8     `json:"isSuper"`
	RoleNames        string   `json:"roleNames"`
	RoleDisplayNames string   `json:"roleDisplayNames"`
	Menus            []string `json:"menus"`
	Uris             []string `json:"uris"`
}

type TokenResult struct {
	Token string `json:"token"`
}

type AuthorizationResult struct {
	Authorization string `json:"authorization"`
}
