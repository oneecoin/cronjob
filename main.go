package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

// "* 0 * * *"
func main() {
	godotenv.Load()
	url := "https://oneecoin-explorer-server.onrender.com/cron?secretKey=" + os.Getenv("SECRET_KEY")

	c := cron.New()

	c.AddFunc("@midnight", func() {
		http.Get(url)
	})

	c.Start()

	select {}
}
