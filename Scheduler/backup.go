package scheduler

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"scheduler-backup-postgresql/Config"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

func StartSchedulerBackup(s *gocron.Scheduler, env Config.Env) {
	s.Every(1).Day().At("00:00").Do(func() {
		BackupDB(env)
	})
}

func BackupDB(env Config.Env) {
	log.Info().Msg("Scheduler Backup DB Start ....")

	location, _ := time.LoadLocation("Asia/Jakarta")
	time := time.Now().In(location).Format("2006_01_02")
	backupFile := fmt.Sprintf("/db_%s_%s.sql", env.Database.Name, time)

	pgDump := "pg_dump"
	cmd := exec.Command(pgDump,
		"-U", env.Database.User,
		"-h", env.Database.Host,
		"-d", env.Database.Name,
	)
	cmd.Env = append(cmd.Env, "PGPASSWORD="+env.Database.Password)

	outputFile, err := os.Create(backupFile)
	if err != nil {
		log.Error().Msgf("Failed to create backup file: %v", err)
		return
	}
	defer outputFile.Close()

	cmd.Stdout = outputFile
	err = cmd.Run()
	if err != nil {
		log.Error().Msgf("Failed to execute pg_dump: %v", err)
		return
	}

	SendFileBackup(backupFile, env)

	log.Info().Msg("Scheduler Backup DB Finish ....")
}
