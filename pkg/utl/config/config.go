package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Load returns Configuration struct
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

type Configuration struct {
	DB     Database `yaml:"database,omitempty"`
	Server *Server  `yaml:"server,omitempty"`
	Nats   *Nats    `json:"nats,omitempty"`
}

type Database struct {
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
}

// Server holds data necessary for server configuration
type Server struct {
	Port         string `yaml:"port,omitempty"`
	Debug        bool   `yaml:"debug,omitempty"`
	ReadTimeout  int    `yaml:"read_timeout_seconds,omitempty"`
	WriteTimeout int    `yaml:"write_timeout_seconds,omitempty"`
}

type Nats struct {
	QOS  string `json:"qos"`
	Host string `json:"host"`
}