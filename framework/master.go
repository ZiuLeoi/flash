package framework

import "flash/proto"

// Master 定义 Master 的行为
type Master interface {
	Start(address string) error                    // 启动 Master 服务
	SubmitTask(task *proto.Task) error             // 提交任务
	RegisterWorker(workerID, address string) error // 注册 Worker
	SetScheduler(scheduler Scheduler)              // 设置任务调度器
}
