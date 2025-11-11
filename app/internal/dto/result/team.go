package result

type TeamRet struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Leader      string `json:"leader"`
	Director    string `json:"director"`
	UpdatedTime string `json:"updatedTime"`
	CreatedTime string `json:"createdTime"`
}

type TeamResult struct {
	Total   int64      `json:"total"`
	RetList []*TeamRet `json:"retList"`
}
