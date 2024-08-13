package App

import (
	"scheduler-backup-postgresql/Config"
	scheduler "scheduler-backup-postgresql/Scheduler"
)

func AppInitialization() {
	scheduler.SchedulerStart(Config.GetEnvironment())
}
