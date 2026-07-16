package dto

import "encoding/json"

type Message struct {
	Id    string `json:"id,omitempty"`  // 消息ID
	BizId int64  `json:"biz,omitempty"` // 业务

	Topic string          `json:"topic,omitempty"` // 消息主题
	Raw   json.RawMessage `json:"raw,omitempty"`   // 消息体

	Executor string `json:"executor,omitempty"` // 执行器

	Retries uint32 `json:"retries,omitempty"` // 重试次数
	ErrMsg  string `json:"err,omitempty"`     // 错误消息

	Issuer    string `json:"iss,omitempty"` // 寄件人
	Audience  string `json:"aud,omitempty"` // 收件人
	IssuedAt  int64  `json:"iat,omitempty"` // 寄件时间（秒级时间戳）
	NotBefore int64  `json:"nbf,omitempty"` // 生效时间（秒级时间戳）
	Expire    int64  `json:"exp,omitempty"` // 失效时间（秒级时间戳）
}
