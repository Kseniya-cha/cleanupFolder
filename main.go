package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Kseniya-cha/cleanupFolder/pkg/config"
)

func main() {
	// Чтение конфигурационного файла
	cfg, err := config.GetConfig()
	if err != nil {
		log.Println(err)
	}

	for {
		time.Sleep(cfg.TimeClean)
		trashs, err := os.ReadDir(cfg.FilePath)
		if err != nil {
			log.Println(err)
		}

		for _, trash := range trashs {
			fmt.Println(trash.Name())
			err = os.Remove(cfg.FilePath + trash.Name())
			if err != nil {
				log.Println(err)
			}
		}
		log.Println("clean trash!")
	}
}
