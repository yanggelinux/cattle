package result

type PlatformRet struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Project     string  `json:"project" `
	Url         string  `json:"url"`
	PlatGroup   string  `json:"platGroup"`
	LoginType   int8    `json:"loginType"`
	Sort        int64   `json:"sort"`
	Description string  `json:"description"`
	IconSrc     string  `json:"iconSrc"`
	RoleNames   string  `json:"roleNames"`
	RoleIDs     []int64 `json:"roleIDs"`
	UpdatedTime string  `json:"updatedTime"`
	CreatedTime string  `json:"createdTime"`
}

type PlatformResult struct {
	Total   int64          `json:"total"`
	RetList []*PlatformRet `json:"retList"`
}
