package framework

import "flash/proto"

// Worker 定义 Worker 的行为
type Worker interface {
	Start() error                            // 启动 Worker 服务
	RegisterTaskHandler(handler TaskHandler) // 注册任务处理器
}

// TaskHandler 定义任务处理逻辑
type TaskHandler interface {
	HandleTask(task *proto.Task) (*proto.TaskStatus, error) // 处理任务
}
