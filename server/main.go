package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/curio-research/keystone-starter-kit/constants"
	"github.com/curio-research/keystone-starter-kit/server"
	"github.com/fatih/color"
	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load()

	// get randomness seed (controls how random numbers are created)
	randSeed := os.Getenv("RAND_SEED")
	if randSeed == "" {
		log.Info("missing RAND_SEED env variable, seeding randomness with local variable", randSeed)
	}

	// get listening port
	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(constants.DefaultListeningPort)
		log.Infof("missing PORT env variable, using %s", port)
	}

	// get websocket port (for streaming updates)
	wsPortStr := os.Getenv("WS_PORT")
	if wsPortStr == "" {
		wsPortStr = strconv.Itoa(constants.DefaultWSPort)
		log.Infof("missing WS_PORT env variable, using %s", wsPortStr)
	}

	wsPort, err := strconv.Atoi(wsPortStr)
	if err != nil {
		log.Fatal(err)
	}

	s, err := server.MainServer(wsPort)
	if err != nil {
		log.Fatal(err)
	}

	color.HiWhite("Listening on port:    " + port)
	color.HiWhite("WS port:    " + wsPortStr)
	fmt.Println()

	log.Fatal(s.Run(":" + port))
}
