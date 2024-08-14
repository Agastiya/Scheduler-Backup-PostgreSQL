package App

import (
	"os"
	"os/signal"
	"scheduler-backup-postgresql/Config"
	scheduler "scheduler-backup-postgresql/Scheduler"
	"syscall"
)

func ServiceInit() {
	environment := Config.GetEnvironment()
	scheduler.SchedulerStart(environment)
}

func AppInitialization() {
	ServiceInit()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
