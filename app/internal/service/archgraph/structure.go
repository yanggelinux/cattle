package archgraph

type Text struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Value string  `json:"value"`
}

type LinkGraph struct {
	ID        int64  `json:"id"`
	GraphName string `json:"graphName"`
}
type Properties struct {
	Width     float64    `json:"width"`
	Height    float64    `json:"height"`
	LinkGraph *LinkGraph `json:"linkGraph"`
}
type Node struct {
	X          float64    `json:"x"`
	Y          float64    `json:"y"`
	ZIndex     float64    `json:"zIndex"`
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Text       Text       `json:"text"`
	Properties Properties `json:"properties"`
}

type Edge struct {
	SourceNodeId   string     `json:"sourceNodeId"`
	TargetNodeId   string     `json:"targetNodeId"`
	SourceAnchorId string     `json:"sourceAnchorId"`
	TargetAnchorId string     `json:"targetAnchorId"`
	X              int64      `json:"x"`
	Y              int64      `json:"y"`
	ZIndex         int64      `json:"zIndex"`
	ID             string     `json:"id"`
	Type           string     `json:"type"`
	Text           Text       `json:"text"`
	Properties     Properties `json:"properties"`
}
