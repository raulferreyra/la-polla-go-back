package main

import (
	"fmt"
	"log"
	"os"

	"go-backend/internal/config"
	"go-backend/internal/db"
	"go-backend/internal/logger"
	"go-backend/internal/models"
	"go-backend/internal/router"
)

func showBanner() {
	path := fmt.Sprint("internal/banner/banner.txt")

	b, err := os.ReadFile(path)
	if err != nil {
		log.Printf("No se pudo leer el banner (%s): %v", path, err)
		return
	}
	fmt.Println(string(b))
}

func main() {
	showBanner()

	cfg := config.Load()
	lg := logger.New()
	defer lg.Sync()

	database, err := db.Open(cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	// Migraciones mínimas
	if err := database.AutoMigrate(
		&models.Company{}, &models.User{},
		&models.Region{}, &models.Country{}, &models.Team{},
		&models.Tournament{}, &models.Match{}, &models.Bet{},
	); err != nil {
		log.Fatal(err)
	}

	app := router.Setup(router.Deps{
		DB: database, JWTSecret: cfg.JWTSecret,
	})

	if err := app.Run(cfg.HttpAddr); err != nil {
		log.Fatal(err)
	}
}
