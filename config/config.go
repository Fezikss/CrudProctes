package config

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
}

func Load() Config {
	cfg := Config{}

	cfg.PostgresHost = "localHost"
	cfg.PostgresPort = "5432"
	cfg.PostgresUser = "postgres"
	cfg.PostgresPassword = "Master0101"
	cfg.PostgresDB = "catalog"

	return cfg
}
