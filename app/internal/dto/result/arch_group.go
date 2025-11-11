package result

type ArchGroupRet struct {
	ID          int64  `json:"id"`
	ParentID    int64  `json:"parentID"`
	GroupName   string `json:"groupName"`
	ItemCount   int64  `json:"itemCount"`
	UpdatedTime string `json:"updatedTime"`
	CreatedTime string `json:"createdTime"`
}

type ArchGroupResult struct {
	RetList []*ArchGroupRet `json:"retList"`
}
