package main

import (
	"io"
	"log"
	"os"
	"path"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Setup() {

	// os.Setenv("TOKEN_HOUR_LIFESPAN", "1")
	// os.Setenv("API_SECRET", "anilkumar")

	// TLS_CERT_PATH=/Users/anikumar/Development/MyProjets/github/onesheds/dev/vars/keyfiles/peerOrganizations/warehousemen.onesheds.com/peers/peer0.warehousemen.onesheds.com/tls/ca.crt
	// PEER_ENDPOINT=localhost:7003
	// PEER_URL=peer0.warehousemen.onesheds.com

	err := godotenv.Load("service.env")
	if err != nil {
		log.Fatal("Error loading service.env file")
	}

	// SetupDataDir
	SetupDataDir()

	//
	gin.SetMode(gin.DebugMode)

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	logFile := path.Join(os.Getenv("DATA_DIR"), "logs", "service.log")
	f, _ := os.Create(logFile)
	//gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func SetupCORS(engine *gin.Engine) {
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	engine.Use(cors.New(cors.Config{
		//AllowAllOrigins: true,
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
}

func SetupDataDir() {

	userHomeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	dataDir := path.Join(userHomeDir, "onesheds", os.Getenv("NODE_TYPE"))

	_, err = os.Stat(dataDir)
	if os.IsNotExist(err) {
		if err := os.Mkdir(dataDir, 0755); err != nil {
			panic(err)
		}
	}

	os.Setenv("DATA_DIR", dataDir)
}
