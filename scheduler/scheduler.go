package scheduler

import "fmt"

type Scheduler interface {
	SelectWorker(workers map[string]*WorkerInfo, task Task) (string, error)
}

type RoundRobinScheduler struct {
	currentIndex int
}

func (s *RoundRobinScheduler) SelectWorker(workers map[string]*WorkerInfo, task Task) (string, error) {
	if len(workers) == 0 {
		return "", fmt.Errorf("no workers available")
	}
	workerIDs := []string{}
	for id := range workers {
		workerIDs = append(workerIDs, id)
	}
	s.currentIndex = (s.currentIndex + 1) % len(workerIDs)
	return workerIDs[s.currentIndex], nil
}
