package main

import (
	"fmt"
	"github.com/curio-research/keystone/game/server"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	godotenv.Load()

	mode := os.Getenv("MODE")
	if mode == "" {
		log.Fatal("Mode not set. Use 'dev' or 'prod'")
	}

	wsPortString := os.Getenv("INDEXER_PORT")
	if wsPortString == "" {
		log.Fatal("websocket port not set in .env file. Set using INDEXER_PORT=xxxx")
	}

	wsPort, err := strconv.Atoi(wsPortString)
	if err != nil {
		log.Fatal("websocket port is not a number")
	}

	mySQLdsn := os.Getenv("MYSQL_DSN")
	if mode == "prod" && mySQLdsn == "" {
		panic("missing MYSQL_DSN env variable")
	}

	// add the default randomness object
	randSeed := os.Getenv("RAND_SEED")
	if randSeed == "" {
		log.Fatal("missing RAND_SEED env variable")
	}

	randSeedNumber, err := strconv.Atoi(randSeed)
	if err != nil {
		log.Fatal(err)
	}

	s, _, err := server.StartMainServer(mode, wsPort, randSeedNumber)
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("missing PORT env variable")
	}

	color.HiWhite("Websocket port:    " + port)
	fmt.Println()

	log.Fatal(s.Run(":" + port))
}
