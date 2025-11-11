package result

type OrderFieldRet struct {
	ID           int64  `json:"id"`
	OrderID      int64  `json:"orderID"`
	OrderName    string `json:"orderName"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	Component    string `json:"component"`
	Placeholder  string `json:"placeholder"`
	VerRule      int8   `json:"verRule"`
	DefaultVal   string `json:"defaultVal"`
	IsRequired   int8   `json:"isRequired"`
	IsTitle      int8   `json:"isTitle"`
	IsEdit       int8   `json:"isEdit"`
	IsClear      int8   `json:"isClear"`
	DisplayField string `json:"displayField"`
	DisplayVal   string `json:"displayVal"`
	Description  string `json:"description"`
	Enum         string `json:"enum"`
	GroupName    string `json:"groupName"`
	Status       int8   `json:"status"`
	Sort         int64  `json:"sort"`
	UpdatedTime  string `json:"updatedTime"`
	CreatedTime  string `json:"createdTime"`
}

type OrderFieldResult struct {
	Total   int64            `json:"total"`
	RetList []*OrderFieldRet `json:"retList"`
}
