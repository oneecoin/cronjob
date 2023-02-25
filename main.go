package main

import (
	"net/http"
	"os"
)

// "* 0 * * *"
func main() {
	url := "https://oneecoin-explorer-server.onrender.com/cron?secretKey=" + os.Getenv("SECRET_KEY")
	http.Get(url)
}
