package bot

import (
	"fmt"
	"strconv"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v7"
	"github.com/gucastiliao/botstate"
	"github.com/gucastiliao/example-chatbot-botstate/pkg/util"
)

type Message struct {
	UserID int    `json:"user_id"`
	Text   string `json:"text"`
}

var states []botstate.State
var commands map[string]string
var RedisClient *redis.Client

func init() {
	states = []botstate.State{
		{
			Name:     "start",
			Executes: sendInitialMessage,
			Callback: callbackInitialMessage,
			Next:     "success",
		},
		{
			Name:     "success",
			Executes: sendMessageSuccess,
			Next:     "start",
		},
	}

	mr, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
}

//SetupBot set storage client as redis
func SetupBot() {
	botstate.SetStorageClient(botstate.DefaultStorage(RedisClient))
}

//GetBotAnswer create new instance of botstate with user id coming from API
//Call botstate.ExecuteState based on current state value
//Store text sended by user in redis so that it is possible to recover this value in the states
//After call ExecuteState, return all user data to API
func GetBotAnswer(message Message) botstate.Data {
	b := botstate.New(states)
	b.Data.User(message.UserID)

	current, _ := b.Data.GetCurrentState()

	if current == "" {
		current, _ = b.Data.SetCurrentState("start")
	}

	b.Data.SetData(botstate.Data{
		"text": message.Text,
	})

	b.ExecuteState(current)

	data, _ := b.Data.GetData()
	b.GetMessages()

	return data
}

func sendInitialMessage(b *botstate.Bot) bool {
	fn := util.Random(0, 100)
	sn := util.Random(0, 100)

	b.Data.SetData(botstate.Data{
		"first_number":  strconv.Itoa(fn),
		"second_number": strconv.Itoa(sn),
	})

	b.SendMessage([]string{
		"Hello!",
		fmt.Sprintf("How much is %d + %d?", fn, sn),
	})

	return true
}

func callbackInitialMessage(b *botstate.Bot) bool {
	fn, _ := strconv.Atoi(b.Data.Current["first_number"])
	sn, _ := strconv.Atoi(b.Data.Current["second_number"])

	sum := fn + sn

	answer, err := strconv.Atoi(b.Data.Current["text"])

	if err != nil {
		b.SendMessage([]string{
			"Sorry, can you repeat, please?",
		})
		return false
	}

	if sum != answer {
		b.SendMessage([]string{
			"Wrong answer!",
		})
		return false
	}

	return true
}

func sendMessageSuccess(b *botstate.Bot) bool {
	b.SendMessage([]string{
		"You are right! Good bye :)",
	})

	b.Data.ResetCurrentState()

	return true
}
