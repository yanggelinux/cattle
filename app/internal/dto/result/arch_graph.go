package result

import "time"

type ArchGraphRet struct {
	ArchGraphCommon
	UpdatedTime string `json:"updatedTime"`
	CreatedTime string `json:"createdTime"`
}

type ArchGraphCommon struct {
	ID         int64  `json:"id"`
	GroupID    int64  `json:"groupID"`
	GroupName  string `json:"groupName"`
	GraphKey   string `json:"graphKey"`
	GraphLabel string `json:"graphLabel"`
	GraphName  string `json:"graphName"`
	ImageData  string `json:"imageData"`
	Owner      string `json:"owner"`
	Status     int8   `json:"status"`
	IsShared   int8   `json:"isShared"`
}
type ArchGraphRecord struct {
	ArchGraphCommon
	UpdatedTime time.Time `json:"updatedTime"`
	CreatedTime time.Time `json:"createdTime"`
}

type ArchGraphResult struct {
	Total   int64           `json:"total"`
	RetList []*ArchGraphRet `json:"retList"`
}
type ArchGraphRecordRet struct {
	ID          int64  `json:"id"`
	GraphID     int64  `json:"graphID"`
	ImageData   string `json:"imageData"`
	ImageHash   string `json:"imageHash"`
	CreatedTime string `json:"createdTime"`
}
type ArchGraphRecordResult struct {
	RetList []*ArchGraphRecordRet `json:"retList"`
}

type ArchGraphDetailResult struct {
	ArchGraphRet
	NodeData  interface{} `json:"nodeData"`
	EdgeData  interface{} `json:"edgeData"`
	ImageHash string      `json:"imageHash"`
}

type ArchGraphOptResult struct {
	ID int64 `json:"id"`
}

type ArchGraphReviewRet struct {
	ID          int64    `json:"id"`
	GraphID     int64    `json:"graphID"`
	GraphKey    string   `json:"graphKey"`
	Reviewer    string   `json:"reviewer"`
	Content     string   `json:"content"`
	NotifyParty []string `json:"notifyParty"`
	CreatedTime string   `json:"createdTime"`
}
type ArchGraphReviewResult struct {
	Total   int64                 `json:"total"`
	RetList []*ArchGraphReviewRet `json:"retList"`
}

type NodeInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
type EdgeInfo struct {
	ID       string `json:"id"`
	SourceID string `json:"sourceID"`
	TargetID string `json:"targetID"`
}

type ArchGraphData struct {
	NodeInfos []*NodeInfo `json:"nodeInfos"`
	EdgeInfos []*EdgeInfo `json:"edgeInfos"`
}
