package dto

type MessageBatcExecRequest struct {
	Batch int64 `json:"batch" form:"batch" url:"batch" uri:"uri"`
}
