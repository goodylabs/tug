package main

import (
	"github.com/goodylabs/docker-swarm-cli/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	cmd.Execute()
}
