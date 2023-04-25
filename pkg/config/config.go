package config

import "time"

type Config struct {
	TimeClean time.Duration `yaml:"timeClean"`
	FilePath  string        `yaml:"filePath"`
}
