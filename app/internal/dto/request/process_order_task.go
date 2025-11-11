package request

type UpdateFirewallOrWhitelistReq struct {
	OrderID     *int64  `json:"orderID" binding:"required"`
	CheckResult *string `json:"checkResult"`
	ExecResult  *string `json:"execResult"`
}
