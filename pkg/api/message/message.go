package api

import (
	"encoding/json"
	"net/http"

	"github.com/gucastiliao/example-chatbot-botstate/pkg/bot"
	"github.com/gucastiliao/example-chatbot-botstate/pkg/user"
)

//Handler process text from user and call bot.GetBotAnswer
func Handler(w http.ResponseWriter, r *http.Request) {
	var result bot.Message

	err := json.NewDecoder(r.Body).Decode(&result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !user.Exists(result.UserID) {
		http.Error(w, "User undefined.", http.StatusBadRequest)
		return
	}

	answer := bot.GetBotAnswer(result)

	j, _ := json.Marshal(answer)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(j))
}
