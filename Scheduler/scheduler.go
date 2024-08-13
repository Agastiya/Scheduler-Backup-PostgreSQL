package Scheduler

import (
	"time"

	"agastiya.github/backup-db/Config"
	"agastiya.github/backup-db/Scheduler/Backup"
	"github.com/go-co-op/gocron"
)

const (
	TimeLocation = "Asia/Jakarta"
)

func SchedulerStart(env Config.Env) {
	time, _ := time.LoadLocation(TimeLocation)
	s := gocron.NewScheduler(time)
	Backup.StartSchedulerBackup(s, env)
	s.StartAsync()
}
