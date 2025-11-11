package result

type OrderGroupRet struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Status      int8   `json:"status"`
	Sort        int64  `json:"sort"`
	UpdatedTime string `json:"updatedTime"`
	CreatedTime string `json:"createdTime"`
}

type OrderGroupResult struct {
	Total   int64            `json:"total"`
	RetList []*OrderGroupRet `json:"retList"`
}
