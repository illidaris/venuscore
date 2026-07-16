package vao

type RetryType int8

const (
	RetryTypeNil          RetryType = iota // 无
	RetryTypeNo                            // 不重试
	RetryTypeImmediately                   // 立刻重试
	RetryTypeQuickly                       // 尽快重试
	RetryTypeCompensation                  // 补偿重试
	RetryTypeHybrid                        // 混合重试
)
