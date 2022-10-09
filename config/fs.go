package config

import (
	"bytes"
	"errors"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path"
)

type Config struct {
	AuthToken string
}

func getConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != err {
		// TODO: Somehow handle not finding the users home
		log.Fatal(err)
	}
	return path.Join(homeDir, ".gliConfig")
}

func ReadFromFilesystem() (Config, error) {
	var conf Config
	if _, err := toml.DecodeFile(getConfigPath(), &conf); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = nil
		}
		return Config{}, err
	}
	return conf, nil
}

func (c *Config) WriteToFilesystem() error {
	var buffer bytes.Buffer
	if err := toml.NewEncoder(&buffer).Encode(c); err != nil {
		return err
	}
	return os.WriteFile(getConfigPath(), buffer.Bytes(), 0666)
}
