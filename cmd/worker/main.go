package worker

import "flash/internal/worker"

func main() {
	w := worker.NewDefaultWorker("worker-1", "localhost:5000")
	w.Start()
}
