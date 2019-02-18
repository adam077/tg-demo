package scheduler

import (
	"github.com/bamzi/jobrunner"
)

func Run() {
	jobrunner.Start(32, 0)

	jobrunner.Schedule("@every 5m", SingleRunner{})
}
