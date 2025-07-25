package main

import (
	"github.com/goodylabs/tug/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cmd.Execute()
}
