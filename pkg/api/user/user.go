package api

import (
	"encoding/json"
	"net/http"

	"github.com/gucastiliao/example-chatbot-botstate/pkg/user"
)

type User struct {
	ID int `json:"user_id"`
}

//Handler create and return new user
func Handler(w http.ResponseWriter, r *http.Request) {
	u := user.Create()

	j, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(j))
}
