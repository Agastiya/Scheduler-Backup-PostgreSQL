package App

import (
	"agastiya.github/backup-db/Config"
	"agastiya.github/backup-db/Scheduler"
)

func AppInitialization() {
	Scheduler.SchedulerStart(Config.GetEnvironment())
}
