package vao

type ExecutorType int8

const (
	ExecutorTypeNil      ExecutorType = iota
	ExecutorTypeNormal                // 一般执行器，需要配置
	ExecutorTypeDiscover              // 内置注册发现服务的执行器
)
