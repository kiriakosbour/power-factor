package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"power-factor/app"
	pfhttp "power-factor/http"
)

func main() {
	router := mux.NewRouter()
	timestampsService := app.NewTimestampInteractor()
	timestampsHandler := pfhttp.NewTimestampDataHandler(*timestampsService)
	router.HandleFunc("/ptlist", timestampsHandler.TimestampsMatching).Methods(http.MethodGet)
	port := "8080"
	log.Printf("Defaulting to port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
