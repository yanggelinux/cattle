package result

type DashboardGraphResult struct {
	TotalCount      int64 `json:"totalCount"`
	UnapprovedCount int64 `json:"unapprovedCount"`
	ApprovingCount  int64 `json:"approvingCount"`
	SuccessCount    int64 `json:"successCount"`
	FailedCount     int64 `json:"failedCount"`
}

type DashboardOrderResult struct {
	TotalCount         int64           `json:"totalCount"`
	UnapprovedCount    int64           `json:"unapprovedCount"`
	ApprovingCount     int64           `json:"approvingCount"`
	SuccessCount       int64           `json:"successCount"`
	FailedCount        int64           `json:"failedCount"`
	GraphApplyDist     []*OrderDistRet `json:"graphApplyDist"`
	GraphChangeDist    []*OrderDistRet `json:"graphChangeDist"`
	ResApplyChangeDist []*OrderDistRet `json:"resApplyChangeDist"`
}

type DashboardDemandResult struct {
	TotalCount      int64 `json:"totalCount"`
	UnapprovedCount int64 `json:"unapprovedCount"`
	ApprovingCount  int64 `json:"approvingCount"`
	SuccessCount    int64 `json:"successCount"`
	FailedCount     int64 `json:"failedCount"`
}

type OrderDistRet struct {
	Count int64  `json:"count"`
	DT    string `json:"dt"`
}
