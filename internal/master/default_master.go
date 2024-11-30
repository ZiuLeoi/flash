package master

import (
	"flash/framework"
	"flash/proto"
	"log"
	"sync"
)

type DefaultMaster struct {
	address   string
	tasks     chan *proto.Task
	workers   map[string]string
	mutex     sync.Mutex
	scheduler framework.Scheduler
}

// 确保 DefaultMaster 实现了 Master 接口
var _ framework.Master = (*DefaultMaster)(nil)

func NewDefaultMaster() *DefaultMaster {
	return &DefaultMaster{
		tasks:   make(chan *proto.Task, 100),
		workers: make(map[string]string),
	}
}

func (m *DefaultMaster) Start(address string) error {
	m.address = address
	log.Printf("Master started at %s", address)
	go m.dispatchTasks()
	return nil
}

func (m *DefaultMaster) SubmitTask(task *proto.Task) error {
	m.tasks <- task
	log.Printf("Task %s submitted", task.Id)
	return nil
}

func (m *DefaultMaster) RegisterWorker(workerID, address string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.workers[workerID] = address
	log.Printf("Worker %s registered at %s", workerID, address)
	return nil
}

func (m *DefaultMaster) SetScheduler(scheduler framework.Scheduler) {
	m.scheduler = scheduler
}

func (m *DefaultMaster) dispatchTasks() {
	for task := range m.tasks {
		if m.scheduler == nil {
			log.Println("Scheduler not set, skipping task dispatch")
			continue
		}
		m.mutex.Lock()
		workerID, address, err := m.scheduler.ScheduleTask(m.workers, task)
		if err != nil {
			log.Printf("Failed to schedule task %s: %v", task.Id, err)
			m.mutex.Unlock()
			continue
		}
		log.Printf("Task %s dispatched to worker %s at %s", task.Id, workerID, address)
		m.mutex.Unlock()
	}
}
