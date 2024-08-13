package App

import (
	"agastiya.github/backup-db/Config"
	scheduler "agastiya.github/backup-db/Scheduler"
)

func AppInitialization() {
	scheduler.SchedulerStart(Config.GetEnvironment())
}
