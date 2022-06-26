package config

import (
	"log"

	"github.com/spf13/viper"
)

type ConfigurationInterface interface {
	LoadConfiguration() error
	GetConfigDetails() ConfigurationDetails
}

type DatabaseDetails struct {
	Host             string `mapstructure:"host"`
	Port             string `mapstructure:"port"`
	Name             string `mapstructure:"name"`
	User             string `mapstructure:"user"`
	Password         string `mapstructure:"password"`
	ConnectionString string `mapstructure:"connectionString"`
}

type JWT struct {
	Secret     string `mapstructure:"secret"`
	ExpiryTime int    `mapstructure:"expiryTime"`
}

type ConfigurationDetails struct {
	ServiceName     string          `mapstructure:"serviceName"`
	ServerPort      string          `mapstructure:"serverPort"`
	DatabaseDetails DatabaseDetails `mapstructure:"database"`
	Jwt             JWT             `mapstructure:"jwt"`
}

type config struct {
	configDetails ConfigurationDetails
}

func NewConfigObject() *config {
	return &config{}
}

func (c *config) LoadConfiguration() error {
	log.Println("load configuration")

	var conf ConfigurationDetails
	viper.AddConfigPath("./configuration/conf/")
	// viper.AddConfigPath("/configuration/conf/")
	// viper.AddConfigPath(".")

	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error while reading configuration : ", err)
		return err
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Println("Error while unmarshalling configuration : ", err)
		return err
	}

	c.configDetails = conf
	return nil
}

func (c *config) GetConfigDetails() ConfigurationDetails {
	return c.configDetails
}
