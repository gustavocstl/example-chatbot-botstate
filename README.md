# Chatbot Example with [Botstate](https://github.com/gucastiliao/botstate/)

A simple chatbot API to show how [botstate](https://github.com/gucastiliao/botstate/) works.

To see this chatbot working [https://web-chatbot-botstate.herokuapp.com/](https://web-chatbot-botstate.herokuapp.com/)

## Start API
```
go run main.go
```

API listen default port `8000`.

This example use [go-redis](https://github.com/go-redis/redis) to save conversation data. See [bot.go](https://github.com/gucastiliao/example-chatbot-botstate/blob/master/pkg/bot/bot.go#L50)

---

## Create new user
POST `/api/v1/bot/user`

**Response:**
```json
{
    "user_id": 119454012254717
}
```

---

## Send a message to the bot
POST `/api/v1/bot/message`

**Body:**
```json
{
    "user_id": 119454012254717,
    "text": "Hello!"
}
```

**Response:**
```json
{
    "current_state": "success",
    "first_number": "13",
    "messages": "[\"Hello!\",\"How much is 13 + 86?\"]",
    "second_number": "86",
    "state_with_callback": "start",
    "text": "Hello!",
    "user_id": "119454012254717"
}
```