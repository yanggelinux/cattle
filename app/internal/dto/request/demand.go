package request

type GetDemandReq struct {
	Status     *int8   `form:"status"`
	Name       *string `form:"name"`
	DemandType *int8   `form:"demandType"`
	Page       *int    `form:"page"`
	PageSize   *int    `form:"pageSize"`
}
type GetDemandDetailReq struct {
	ID   *int64  `form:"id"`
	Name *string `form:"name"`
}

type CreateDemandReq struct {
	Name        *string `json:"name" binding:"required"`
	DemandType  *int8   `json:"demandType" binding:"required"`
	Biz         *string `json:"biz" binding:"required"`
	Owner       *string `json:"owner" binding:"required"`
	Description *string `json:"description" binding:"required"`
	Status      *int8   `json:"status"`
	OrderNo     *string `json:"orderNo"`
}

type UpdateDemandReq struct {
	ID          *int64  `json:"id" binding:"required"`
	Name        *string `json:"name"`
	DemandType  *int8   `json:"demandType"`
	Biz         *string `json:"biz"`
	Owner       *string `json:"owner"`
	Description *string `json:"description"`
	Opinion     *string `json:"opinion"`
	Status      *int8   `json:"status"`
	OrderNo     *string `json:"orderNo"`
}
type ApproveDemandReq struct {
	ID           *int64  `json:"id" binding:"required"`
	Action       *string `json:"action" binding:"required"`
	Opinion      *string `json:"opinion"`
	Approver     *string `json:"approver" binding:"required"`
	ApproverName *string `json:"approverName" binding:"required"`
}

type EvaluateDemandReq struct {
	ID            *int64  `json:"id" binding:"required"`
	OpsEvaluation *string `json:"opsEvaluation"`
	OpsReason     *string `json:"opsReason"`
	ResEvaluation *string `json:"resEvaluation"`
	ResReason     *string `json:"resReason"`
	NetEvaluation *string `json:"netEvaluation"`
	NetReason     *string `json:"netReason"`
}
