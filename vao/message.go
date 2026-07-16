package vao

type MessageType int8

const (
	MessageTypeNil   MessageType = iota
	MessageTypeOnce              // 一次性，执行一次
	MessageTypeSync              // 同步执行，同步执行接过当场返回
	MessageTypeAsync             // 异步消息，只确认是否投递成功
)
