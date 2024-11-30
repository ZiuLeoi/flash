package master

import (
	"flash/internal/master"
	"flash/internal/scheduler"
)

func main() {
	m := master.NewDefaultMaster()
	m.SetScheduler(&scheduler.RandomScheduler{})
	m.Start(":5000")
}
