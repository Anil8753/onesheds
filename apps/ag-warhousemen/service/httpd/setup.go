package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Setup() {

	// os.Setenv("TOKEN_HOUR_LIFESPAN", "1")
	// os.Setenv("API_SECRET", "anilkumar")

	// os.Setenv("TLS_CERT_PATH", "/Users/anikumar/Development/MyProjets/github/onesheds/dev/vars/keyfiles/peerOrganizations/warehousemen.onesheds.com/peers/peer0.warehousemen.onesheds.com/tls/ca.crt")
	// os.Setenv("PEER_ENDPOINT", "localhost:7003")
	// os.Setenv("PEER_URL", "peer0.warehousemen.onesheds.com")
	err := godotenv.Load("service.env")
	if err != nil {
		log.Fatal("Error loading service.env file")
	}
	fmt.Println(os.Getenv("TLS_CERT_PATH"), os.Getenv("PEER_URL"))
	//
	gin.SetMode(gin.DebugMode)

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("service.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
