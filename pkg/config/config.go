package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"gopkg.in/yaml.v2"
)

const (
	defaultConfigPath     = "notification_handler.properties"
	configFileEnvKey      = "CONFIG_FILE_PATH"
)

var configInstance struct {
	instance *Config
	once     sync.Once
}

// Config ...
type Config struct {
	EmailServerPort               string `yaml:"EmailServerPort"`
	EmailServerAddress            string `yaml:"EmailServerAddress"`
	SlackServerPort               string `yaml:"SlackServerPort"`
	SlackServerAddress            string `yaml:"SlackServerAddress"`
	KafkaTopic                    string `yaml:"KafkaTopic"`
	KafkaConsumerGroups         []string `yaml:"KafkaConsumerGroup"`
	KafkaProducers              []string `yaml:"KafkaProducers"`
	KafkaConsumerGroupID          string `yaml:"KafkaConsumerGroupID"`
}

// GetInstance ...
func GetInstance() *Config {
	configInstance.once.Do(func() {
		configInstance.instance = &Config{}
		if err := configInstance.instance.parseAndLoadConfig(); err != nil {
			log.Fatal("failed to load configuration exiting ", err)
		}
	})
	return configInstance.instance
}

func (c *Config) parseAndLoadConfig() error {
	configFile := getEnv(configFileEnvKey, defaultConfigPath)
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Print("mcn trace manager config: Unable to read config/properties file", err.Error())
		return err
	}

	if err = yaml.Unmarshal(data, c); err != nil {
		log.Print("mcn trace manager config: error un-marshaling config file", err.Error())
		return err
	}

	fmt.Printf("%+v", c)

	return nil
}

// getEnv returns if env variable value if set else defaultvalue
func getEnv(env, defaultValue string) string {
	val := os.Getenv(env)
	if val == "" {
		val = defaultValue
	}
	return val
}
