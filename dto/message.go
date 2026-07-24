package dto

import (
	"encoding/json"

	"github.com/illidaris/venuscore/vao"
)

type Message struct {
	Id              string           `json:"id,omitempty"  gorm:"column:id;type:varchar(36);primaryKey;comment:唯一ID"`                  // 消息ID
	BizId           int64            `json:"bizId,omitempty" gorm:"column:bizId;type:int;index:busi_topic;comment:业务Id"`               // 业务Id
	Topic           string           `json:"topic,omitempty" gorm:"column:topic;type:varchar(64);index:busi_topic;comment:消息主题"`       // 消息主题
	Partition       int8             `json:"partition" gorm:"column:partition;type:tinyint;comment:分区"`                                // 分区
	Group           string           `json:"group,omitempty" gorm:"column:group;type:varchar(36);index;comment:分组"`                    // 分组
	RType           string           `json:"rType,omitempty" gorm:"column:rType;type:varchar(32);comment:数据类型"`                        // 数据类型
	Raw             json.RawMessage  `json:"raw,omitempty" gorm:"column:raw;type:text;serializer:json;comment:原始响应数据"`                 // 消息体
	MType           vao.MessageType  `json:"mType,omitempty" gorm:"column:mType;type:tinyint;default:1;comment:消息类型"`                  // 消息类型
	QType           vao.QueueType    `json:"qType,omitempty" gorm:"column:qType;type:tinyint;default:1;comment:队列类型"`                  // 队列类型
	ExecutorType    vao.ExecutorType `json:"executorType,omitempty" gorm:"column:executorType;type:tinyint;default:1;comment:执行器类型"`   // 执行器类型
	RetryType       vao.RetryType    `json:"retryType,omitempty" gorm:"column:retryType;type:tinyint;default:1;comment:重试类型"`          // 重试类型
	Executor        string           `json:"executor,omitempty" gorm:"column:executor;type:varchar(32);comment:执行器"`                   // 执行器
	InDscvrName     string           `json:"inDscvrName,omitempty"  gorm:"column:inDscvrName;type:varchar(36);comment:发现服务名"`          // 防竞争锁
	InDscvrGroup    string           `json:"inDscvrGroup,omitempty"  gorm:"column:inDscvrGroup;type:varchar(36);comment:发现服务组"`        // 防竞争锁
	InDscvrClusters string           `json:"inDscvrClusters,omitempty"  gorm:"column:inDscvrClusters;type:varchar(64);comment:发现服务集群"` // 防竞争锁
	Locker          string           `json:"locker,omitempty"  gorm:"column:locker;type:varchar(36);default:0;index;comment:防竞争锁"`     // 防竞争锁
	LockExp         int64            `json:"lockExp,omitempty"  gorm:"column:expire;type:bigint;default:0;index;comment:锁有效期"`         // 锁有效期
	Timeout         int64            `json:"timeout,omitempty" gorm:"column:timeout;<-:create;type:bigint;default:0;comment:任务超时（秒）"`  // 任务超时时间
	Retries         int32            `json:"retries,omitempty" gorm:"column:retries;type:int;default:0;index;comment:重试次数"`            // 重试次数
	ExecAt          int64            `json:"execAt,omitempty" gorm:"column:execAt;type:bigint;comment:执行时间"`                           // 执行时间
	ExecErr         string           `json:"execErr,omitempty" gorm:"column:execErr;type:varchar(255);comment:执行错误"`                   // 执行错误
	Issuer          string           `json:"iss,omitempty" gorm:"column:iss;type:varchar(32);comment:寄件人"`                             // 寄件人
	Audience        string           `json:"aud,omitempty" gorm:"column:aud;type:varchar(32);comment:收件人"`                             // 收件人
	IssuedAt        int64            `json:"iat,omitempty" gorm:"column:iat;type:bigint;comment:投递时间"`                                 // 寄件时间（秒级时间戳）
	NotBefore       int64            `json:"nbf,omitempty" gorm:"column:nbf;type:bigint;comment:生效时间"`                                 // 生效时间（秒级时间戳）
	Expire          int64            `json:"exp,omitempty" gorm:"column:exp;type:bigint;comment:失效时间"`                                 // 失效时间（秒级时间戳）
	TradeId         string           `json:"traceId" gorm:"column:traceId;type:varchar(36);comment:追踪Id"`                              // 追踪Id
	Status          int8             `json:"status" gorm:"column:status;type:tinyint;default:1;comment:状态"`                            // 状态 0-默认
	CreateAt        int64            `json:"createAt" gorm:"column:createAt;<-:create;index;autoCreateTime;comment:创建时间"`              // 创建时间
	UpdateAt        int64            `json:"updateAt" gorm:"column:updateAt;index;autoUpdateTime;comment:修改时间"`                        // 修改时间
	Describe        string           `json:"describe" gorm:"column:describe;type:varchar(255);comment:描述"`                             // 描述
}
