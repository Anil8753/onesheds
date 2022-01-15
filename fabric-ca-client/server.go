package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anil8753/fabric-ca-client/handlers"
)

func StartServer() {

	http.HandleFunc("/v1/enrolladmin", handlers.EnrollAdminHandler())
	http.HandleFunc("/v1/createidentity", handlers.CreateIdentityHandler())
	http.HandleFunc("/v1/allindentities", handlers.AllIdentitiesHandler())

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	fmt.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
