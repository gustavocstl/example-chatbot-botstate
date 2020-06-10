package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gucastiliao/example-chatbot-botstate/pkg/api"
	"github.com/gucastiliao/example-chatbot-botstate/pkg/bot"
)

func main() {
	var PORT string = os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		PORT = "8000"
	}

	router := mux.NewRouter()

	api.SetupRoutes(router)
	bot.SetupBot()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "content-type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	fmt.Println("API listen on port", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handlers.CORS(headersOk, originsOk, methodsOk)(loggedRouter)))
}
