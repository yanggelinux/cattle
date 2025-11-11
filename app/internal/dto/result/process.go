package result

import "gorm.io/datatypes"

type ProcessRet struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	ProcInfo    datatypes.JSON `json:"procInfo"`
	Status      int8           `json:"status"`
	UpdatedTime string         `json:"updatedTime"`
	CreatedTime string         `json:"createdTime"`
}

type ProcessDetailRet struct {
	ProcessRet
	NodeData datatypes.JSON `json:"nodeData"`
	EdgeData datatypes.JSON `json:"edgeData"`
}

type ProcessResult struct {
	Total   int64         `json:"total"`
	RetList []*ProcessRet `json:"retList"`
}

type ProcessOptResult struct {
	ID int64 `json:"id"`
}
