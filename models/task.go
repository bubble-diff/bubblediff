package models

// Task diff任务结构体
type Task struct {
	// ID 任务id
	ID int64 `json:"id" bson:"id"`
	// Name 任务名称 全局唯一
	Name string `json:"name" bson:"name"`
	// Description 任务描述
	Description string `json:"description" bson:"description"`
	// Owner 任务负责人 todo：可以有多个负责人
	Owner User `json:"owner" bson:"owner"`
	// IsRunning Diff任务运行状态
	IsRunning bool `json:"is_running" bson:"is_running"`

	TrafficConfig *TrafficConfig `json:"traffic_config" bson:"traffic_config"`
	FilterConfig  *FilterConfig  `json:"filter_config" bson:"filter_config"`
	AdvanceConfig *AdvanceConfig `json:"advance_config" bson:"advance_config"`

	// TotalRecord 流量回放的流量数
	TotalRecord int64 `json:"total_record" bson:"total_record"`
	// SuccessRecord 成功回放的流量数
	SuccessRecord int64 `json:"success_record" bson:"success_record"`
	// CreatedTime 任务创建时间
	CreatedTime string `json:"created_time" bson:"created_time"`
	// UpdatedTime 配置最后变更时间
	UpdatedTime string `json:"updated_time" bson:"updated_time"`
}
