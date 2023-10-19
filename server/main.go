package main

import (
	"fmt"
	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/server"
	"github.com/curio-research/keystone/logging"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// TODO use urfave cli?
func main() {
	godotenv.Load()

	// get randomness seed (controls how random numbers are created)
	randSeed := os.Getenv("RAND_SEED")
	if randSeed == "" {
		logging.Log().Infof("missing RAND_SEED env variable, seeding randomness with local variable", randSeed)
	}

	randSeedNumber, err := strconv.Atoi(randSeed)
	if err != nil {
		log.Fatal(err)
	}

	s, err := server.MainServer(randSeedNumber)
	if err != nil {
		log.Fatal(err)
	}

	// get listening port
	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(constants.DefaultListeningPort)
		logging.Log().Infof("missing PORT env variable, using %s", port)
	}

	color.HiWhite("Listening on port:    " + port)
	fmt.Println()

	log.Fatal(s.Run(":" + port))
}
