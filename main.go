package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Kseniya-cha/cleanupFolder/pkg/config"
	"github.com/Kseniya-cha/cleanupFolder/pkg/logger"
)

func main() {
	// Чтение конфигурационного файла
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	log := logger.NewLogger(cfg)

	for {
		time.Sleep(cfg.TimeClean)
		trashs, err := os.ReadDir(cfg.FilePath)
		if err != nil {
			log.Error(err.Error())
		}

		for _, trash := range trashs {
			fmt.Println(trash.Name())
			err = os.Remove(cfg.FilePath + trash.Name())
			if err != nil {
				log.Error(err.Error())
			}
		}
		log.Info("clean trash!")
	}
}
