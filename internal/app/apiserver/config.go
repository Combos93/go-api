package apiserver

type Config struct {
	BindAddr string
	LogLevel string
}

func NewConfig(env map[string]string) *Config {
	return &Config{
		BindAddr: env["BIND_ADDR"],
		LogLevel: env["LOG_LEVEL"],
	}
}
