package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Server        ServerConfiguration
	Database      DatabaseConfiguration
	Credential    GCStorageConfiguration
	Microservices Microservices
}

type Microservices struct {
	Auth string
}

type GCStorageConfiguration struct {
	Gcbucket                       string
	GOOGLE_APPLICATION_CREDENTIALS string
}

type DatabaseConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

type ServerConfiguration struct {
	Port   string
	Secret string
	Mode   string
}

// SetupDB initialize configuration
func Setup(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error al leer el archivo de configuración, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("No se puede decodificar en estructura, %v", err)
	}

	Config = configuration
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
