package worker

import "log"

type Worker interface {
	Start() error                            // 启动 Worker 服务
	RegisterTaskHandler(handler TaskHandler) // 注册任务处理逻辑
	ReportStatus() error                     // 定期报告状态
}

type TaskHandler interface {
	HandleTask(taskID string, payload string) (result string, err error) // 处理任务
}

type DefaultWorker struct {
	id      string
	master  string
	handler TaskHandler
}

// Start 实现 Worker 接口
func (w *DefaultWorker) Start() error {
	log.Printf("Worker %s started", w.id)
	err := w.ReportStatus()
	if err != nil {
		return err
	}
	return nil
}

func (w *DefaultWorker) RegisterTaskHandler(handler TaskHandler) {
	w.handler = handler
}

func (w *DefaultWorker) ReportStatus() error {
	log.Printf("Worker %s reporting status", w.id)
	return nil
}

func NewDefaultWorker(workerID, masterAddress string) Worker {
	return &DefaultWorker{
		id:     workerID,
		master: masterAddress,
	}
}
