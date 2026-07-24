package vao

type QueueType int8

const (
	QueueTypeNil      QueueType = iota
	QueueTypeQueue              // 消息队列
	QueueTypeSchedule           // 调度队列
	QueueTypeDelay              // 延时队列
)
