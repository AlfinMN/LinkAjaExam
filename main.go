package main

import (
	"TestLinkAja/configuration"
	master "TestLinkAja/masters"
	"log"
)

func main() {
	db, err, serverHost, serverPort := configuration.Connection()
	if err != nil {
		log.Fatal(err)
	}

	myRoute := configuration.CreateRouter()
	master.InitData(myRoute, db)
	configuration.RunServer(myRoute, serverHost, serverPort)
}
