package config

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

func NewConfig() *Config {
	return &Config{}
}

// GetConfig инициализирует и заполняет структуру конфигурационного файла
func GetConfig() (*Config, error) {
	var cfg Config

	configPath := "./"

	var v = viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(configPath)

	err := cfg.readParametersFromConfig(v)
	if err != nil {
		return &cfg, err
	}

	// Проверка наличия параметров в командной строке
	readFlags(&cfg)

	return &cfg, err
}

// readParametersFromConfig производит чтение и десериализацию конфигурационного файла
func (cfg *Config) readParametersFromConfig(v *viper.Viper) error {
	// Попытка чтения конфига
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("cannot read config: %v", err)
	}
	// Попытка заполнение структуры Config полученными данными
	if err := v.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("cannot read config: %v", err)
	}
	return nil
}

// readFlags реализует возможность передачи параметров
// конфигурационного файла при запуске из командной строки
func readFlags(cfg *Config) {
	flag.DurationVar(&cfg.TimeClean, "timeClean", cfg.TimeClean, "time for period clean")
	flag.StringVar(&cfg.FilePath, "filePath", cfg.FilePath, "file for clean")
	flag.Parse()
}
