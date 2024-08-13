package Config

import (
	"os"
	"path"
	"runtime"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

func GetEnvironment() (configEnv Env) {
	_, filename, _, _ := runtime.Caller(1)
	envPath := path.Join(path.Dir(filename), "../Environment/env.yml")

	if _, err := os.Stat(envPath); err != nil {
		log.Info().Msg(err.Error())
		panic(err)
	}

	readEnvfile, err := os.ReadFile(envPath)
	if err != nil {
		log.Info().Msg(err.Error())
		panic(err)
	}

	if err := yaml.Unmarshal(readEnvfile, &configEnv); err != nil {
		log.Info().Msg(err.Error())
		panic(err)
	}

	log.Info().Msg("Env load successfully!")
	return
}
