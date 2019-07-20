package scheduler

import (
	"github.com/bamzi/jobrunner"
)

func Run() {
	jobrunner.Start(32, 0)

	jobrunner.Schedule("@every 5m", HotSpotRunner{})

	jobrunner.Schedule("0 0 11,17 * * *", EatWhat{do: Choose})
	jobrunner.Schedule("0 0 12,18 * * *", EatWhat{do: Result})
}
