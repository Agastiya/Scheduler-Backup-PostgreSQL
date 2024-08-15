package scheduler

import (
	"time"

	"scheduler-backup-postgresql/Config"
	"scheduler-backup-postgresql/Notifiers"

	"github.com/go-co-op/gocron"
)

const (
	TimeLocation = "Asia/Jakarta"
)

func SchedulerStart(env Config.Env) {
	time, _ := time.LoadLocation(TimeLocation)
	s := gocron.NewScheduler(time)
	StartSchedulerBackup(s, env)
	s.StartAsync()
}

func SendFileBackup(backupFile string, env Config.Env) {

	// Send File di Discord Channel
	Notifiers.SendToDiscord(backupFile, env)
}
