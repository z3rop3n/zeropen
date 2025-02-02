package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zeropen/app/benghazi/postgresdriver"
	"github.com/zeropen/app/spector"
	"github.com/zeropen/app/spector/config"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
func main() {
	emailAuthGoogle := os.Getenv("EMAIL_AUTH_CODE_GOOGLE")
	emailFromAddress := os.Getenv("EMAIL_FROM")

	appConfig := config.AppConfig{
		EmailAuthGoogle: emailAuthGoogle,
		EmailFrom:       emailFromAddress,
	}

	postgresdb := postgresdriver.NewPostgresDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	conf, err := postgresdb.Connect()

	if err != nil {
		log.Fatalf("could not connect to postgres: %s\n", err)
	}

	spector.Run(*conf, appConfig)
}
