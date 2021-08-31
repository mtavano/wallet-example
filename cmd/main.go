package main

import (
	"os"
	"strconv"
	"time"

	"github.com/mtavano/wallet-example/api"
	"github.com/mtavano/wallet-example/pkg/database"
)

func main() {
	portString := os.Getenv("PORT")
	port, err := strconv.Atoi(portString)
	check(err)

	dbStore := database.NewStore(time.Now)

	server := api.NewServer(dbStore, port)

	server.Run()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
