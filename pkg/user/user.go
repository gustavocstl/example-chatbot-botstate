package user

import (
	"strconv"

	"github.com/gucastiliao/example-chatbot-botstate/pkg/bot"
	"github.com/gucastiliao/example-chatbot-botstate/pkg/util"
)

//Create generate random ID to user and save in redis
func Create() map[string]int {
	var id int

	for {
		id = util.Random(0, 1000000000000000)

		if Exists(id) == false {
			bot.RedisClient.HSet(strconv.Itoa(id), "user_id", "")
			break
		}
	}

	return map[string]int{
		"user_id": id,
	}
}

//Exists return if ID exists in redis
func Exists(id int) bool {
	user, _ := bot.RedisClient.HGetAll(strconv.Itoa(id)).Result()

	if len(user) <= 0 {
		return false
	}

	return true
}
