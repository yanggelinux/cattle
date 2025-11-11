package result

type DemandRet struct {
	ID               int64          `json:"id"`
	Name             string         `json:"name"`
	DemandType       int8           `json:"demandType"`
	OrderNo          string         `json:"orderNo"`
	Biz              string         `json:"biz"`
	Owner            string         `json:"owner"`
	Description      string         `json:"description"`
	Opinion          string         `json:"opinion"`
	ReviewProcess    []*PorcessNode `json:"reviewProcess"`
	CurReviewNode    *PorcessNode   `json:"curReviewNode"`
	Evaluation       *Evaluation    `json:"evaluation"`
	EvaluationRes    string         `json:"evaluationRes"`
	EvaluationReason string         `json:"evaluationReason"`
	IsEvaluate       int8           `json:"isEvaluate"`
	HasReview        int8           `json:"hasReview"`
	ActiveIndex      int8           `json:"activeIndex"`
	Status           int8           `json:"status"`
	UpdatedTime      string         `json:"updatedTime"`
	CreatedTime      string         `json:"createdTime"`
}

type DemandResult struct {
	Total   int64        `json:"total"`
	RetList []*DemandRet `json:"retList"`
}

type Evaluation struct {
	OpsEvaluation string `json:"opsEvaluation"`
	OpsReason     string `json:"opsReason"`
	ResEvaluation string `json:"resEvaluation"`
	ResReason     string `json:"resReason"`
	NetEvaluation string `json:"netEvaluation"`
	NetReason     string `json:"netReason"`
}
