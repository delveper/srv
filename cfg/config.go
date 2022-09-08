package cfg

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

const defaultPath = "./cfg/config.yml"

// Option contain all basic configuration
type Option struct {
	HTTP struct {
		Host         string        `yaml:"host"`
		Port         string        `yaml:"port"`
		ReadTimeout  time.Duration `yaml:"readTimeout"`
		WriteTimeout time.Duration `yaml:"writeTimeout"`
		IdleTimeout  time.Duration `yaml:"idleTimeout"`
	} `yaml:"http"`

	Params []string `yaml:"params"`

	Endpoints struct {
		Random string `yaml:"main"`
		Root   string `yaml:"root"`
		Index  string `yaml:"index"`
	} `yaml:"endpoints"`
}

// Load retrieves config from yaml file
// into predefined struct Option
func Load() (opt *Option, err error) {
	file, err := os.Open(defaultPath)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			log.Println(err)
		}
	}(file)

	if err = yaml.NewDecoder(file).Decode(&opt); err != nil {
		return nil, err
	}

	return opt, nil
}
