package main

import (
	"bitmagnet-io/bitmagnet/internal/dev/app"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app.New().Run()
}
