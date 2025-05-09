package logger

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

type LogEvent struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	Source  string `json:"source"`
}

func InitLogger(redisClient *redis.Client) {
	rdb = redisClient
}

func PublishLog(level, message, source string) {
	if rdb == nil {
		log.Println("Logger not initialized")
		return
	}

	event := LogEvent{
		Level:   level,
		Message: message,
		Source:  source,
	}

	data, err := json.Marshal(event)
	if err != nil {
		log.Println("Failed to serialize log event:", err)
		return
	}

	err = rdb.Publish(context.Background(), "log_channel", data).Err()
	if err != nil {
		log.Println("Failed to publish log event:", err)
	}
}

func StartLogSubscriber() {
	go func() {
		ctx := context.Background()
		sub := rdb.Subscribe(ctx, "log_channel")
		ch := sub.Channel()

		for msg := range ch {
			var event LogEvent
			if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
				log.Println("Ошибка при разборе лога:", err)
				continue
			}
			log.Printf("[LOG] [%s] %s: %s", event.Level, event.Source, event.Message)
		}
	}()
}
