package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"log"
)

type DatabaseConfigurations struct {
	Connection_string string `koanf:"connectionstring"`
}

type KafkaConfiguration struct {
	Kafka_brokers string `koanf:"kafkabroker"`
	Kafka_topic   string `koanf:"topic"`
	MaxBytes      int    `koanf:"maxBytes"`
}

type Configurations struct {
	Db          DatabaseConfigurations `koanf:"database"`
	Kafkastream KafkaConfiguration     `koanf:"kafka"`
}

func LoadConfig() Configurations {
	k := koanf.New(".")
	err := k.Load(file.Provider("./configure_variables/configure_variables.yaml"), yaml.Parser())
	if err != nil {
		log.Fatalf("There is something wrong with the configuration file .%v", err)
	}
	var configuration Configurations
	err = k.Unmarshal("", &configuration)
	if err != nil {
		log.Fatalf("There is something wrong with the configuration file .%v", err)
	}
	return configuration
}
