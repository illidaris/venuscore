package dto

import (
	acDto "github.com/illidaris/aphrodite/dto"
)

type MessagePageListRequest struct {
	acDto.Page
	BizId int64
	Busi  int64
	Topic string
	Group []string
	Ids   []string
}

type ExecResult struct {
	Id      string
	Success bool
	ErrMsg  string
}
type MessageBatchExecResult struct {
	Data []ExecResult `json:"data"`
}
type MessageBatchExecResponse struct {
	acDto.BaseResponse
	Data MessageBatchExecResult `json:"data"`
}

type MessagePageListResponse struct {
	acDto.BaseResponse
	Data acDto.GenericPager[*Message] `json:"data"`
}

type MessagePushRequest struct {
	Msgs []Message `json:"msgs"`
}
