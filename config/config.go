package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	godotenv.Load()
	log.Info("Stage : " + os.Getenv("stage"))
}
