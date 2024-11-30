package framework

import "flash/proto"

// Scheduler 定义任务调度策略
type Scheduler interface {
	ScheduleTask(workers map[string]string, task *proto.Task) (workerID, address string, err error)
}
