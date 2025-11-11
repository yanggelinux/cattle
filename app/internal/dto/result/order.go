package result

type OrderRet struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	GroupID     int64    `json:"groupID"`
	ProcessID   int64    `json:"processID"`
	GroupName   string   `json:"groupName"`
	ProcessName string   `json:"processName"`
	OrderType   int8     `json:"orderType"`
	NodeType    []string `json:"nodeType"`
	Label       string   `json:"label"`
	Layout      int8     `json:"layout"`
	IsTask      int8     `json:"isTask"`
	TaskUrl     string   `json:"taskUrl"`
	TaskMethod  string   `json:"taskMethod"`
	Sort        int64    `json:"sort"`
	Status      int8     `json:"status"`
	UpdatedTime string   `json:"updatedTime"`
	CreatedTime string   `json:"createdTime"`
}

type OrderResult struct {
	Total   int64       `json:"total"`
	RetList []*OrderRet `json:"retList"`
}

type OrderNodeRet struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}
