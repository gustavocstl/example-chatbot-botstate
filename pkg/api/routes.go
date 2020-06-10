package api

import (
	"net/http"

	"github.com/gorilla/mux"
	message "github.com/gucastiliao/example-chatbot-botstate/pkg/api/message"
	user "github.com/gucastiliao/example-chatbot-botstate/pkg/api/user"
)

func SetupRoutes(r *mux.Router) {
	router := r.PathPrefix("/api/v1").Subrouter()

	router.HandleFunc("/bot/message", message.Handler).Methods(http.MethodPost)
	router.HandleFunc("/bot/user", user.Handler).Methods(http.MethodPost)
}
