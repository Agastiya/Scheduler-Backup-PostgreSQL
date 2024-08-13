package Config

type EnvFile []byte

type Env struct {
	Database Database `yaml:"database"`
	Discord  Discord  `yaml:"discord"`
}

type Database struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Discord struct {
	Webhook string `yaml:"webhook"`
}
