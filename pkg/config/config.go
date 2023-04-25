package config

import "time"

type Config struct {
	TimeClean time.Duration `yaml:"timeClean"`
	FilePath  string        `yaml:"filePath"`
	Logger    `yaml:"logger"`
}

type Logger struct {
	LogLevel        string `mapstructure:"logLevel"`
	LogFileEnable   bool   `mapstructure:"logFileEnable"`
	LogStdoutEnable bool   `mapstructure:"logStdoutEnable"`
	LogFile         string `mapstructure:"logpath"`
	RewriteLog      bool   `mapstructure:"rewriteLog"`
}
