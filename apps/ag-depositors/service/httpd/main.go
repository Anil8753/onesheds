package main

import (
	"log"
	"os"

	"github.com/anil8753/onesheds/apps/warehousemen/service/db"
	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
	"github.com/gin-gonic/gin"
)

func main() {

	Setup()

	engine := gin.Default()
	SetupCORS(engine)

	db := db.NewLevelDB("generic")
	ledger := ledger.Ledger{}
	ledger.Init()

	InitRoutes(engine, db, &ledger)

	PrintConfig()
	engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func PrintConfig() {

	log.Println("---------------------------------- Configs ----------------------------------")
	log.Println("PORT:", os.Getenv("PORT"))
	log.Println("TLS_CERT_PATH:", os.Getenv("TLS_CERT_PATH"))
	log.Println("PEER_ENDPOINT:", os.Getenv("PEER_ENDPOINT"))
	log.Println("PEER_URL:", os.Getenv("PEER_URL"))
	log.Println("LEDGER_CHANNEL:", os.Getenv("LEDGER_CHANNEL"))
	log.Println("LEDGER_CHAINCODE:", os.Getenv("LEDGER_CHAINCODE"))
	log.Println("NODE_TYPE:", os.Getenv("NODE_TYPE"))
	log.Println("DATA_DIR:", os.Getenv("DATA_DIR"))
	log.Println("-----------------------------------------------------------------------------")
}
