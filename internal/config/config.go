package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Environment      string `envconfig:"env"`
	DatabaseHost     string `envconfig:"database_host"`
	DatabasePort     string `envconfig:"database_port"`
	DatabaseDBName   string `envconfig:"database_name"`
	DatabaseSchema   string `envconfig:"database_schema"`
	DatabaseUsername string `envconfig:"database_username"`
	DatabasePassword string `envconfig:"database_password"`
	MigrationsDir    string `envconfig:"migrations_dir"`

	KafkaKitchenTopic string `envconfig:"kafka_kitchen_topic"`
	KafkaBroker       string `envconfig:"kafka_broker"`
}

func New() (cfg Config, err error) {
	err = envconfig.Process("", &cfg)
	return cfg, err
}
