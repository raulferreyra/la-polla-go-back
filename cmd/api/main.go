package main

import (
	"log"

	"polla/internal/config"
	"polla/internal/db"
	"polla/internal/logger"
	"polla/internal/models"
	"polla/internal/router"
)

func main() {
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
