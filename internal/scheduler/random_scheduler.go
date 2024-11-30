package scheduler

import (
	"errors"
	"flash/framework"
	"flash/proto"
	"math/rand"
)

type RandomScheduler struct{}

// 确保 RandomScheduler 实现了 Scheduler 接口
var _ framework.Scheduler = (*RandomScheduler)(nil)

func (s *RandomScheduler) ScheduleTask(workers map[string]string, task *proto.Task) (workerID, address string, err error) {
	if len(workers) == 0 {
		return "", "", errors.New("no workers available")
	}
	keys := make([]string, 0, len(workers))
	for id := range workers {
		keys = append(keys, id)
	}
	selected := keys[rand.Intn(len(keys))]
	return selected, workers[selected], nil
}
