package worker

import (
	"flash/framework"
	"flash/proto"
	"log"
)

type DefaultWorker struct {
	id      string
	master  string
	handler framework.TaskHandler
}

// 确保 DefaultWorker 实现了 Worker 接口
var _ framework.Worker = (*DefaultWorker)(nil)

func NewDefaultWorker(id, master string) *DefaultWorker {
	return &DefaultWorker{id: id, master: master}
}

func (w *DefaultWorker) Start() error {
	log.Printf("Worker %s started, connected to Master %s", w.id, w.master)
	return nil
}

func (w *DefaultWorker) RegisterTaskHandler(handler framework.TaskHandler) {
	w.handler = handler
}

func (w *DefaultWorker) handleTask(task *proto.Task) {
	if w.handler == nil {
		log.Printf("No handler registered for worker %s", w.id)
		return
	}
	result, err := w.handler.HandleTask(task)
	if err != nil {
		log.Printf("Failed to handle task %s: %v", task.Id, err)
		return
	}
	log.Printf("Task %s processed: %+v", task.Id, result)
}
