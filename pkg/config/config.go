package config

import (
	"os"
	"strings"
)

type config struct {
	BootStrapServers string
	WebSocketServer  string
	KafkaTopic       string
	LogLevel         string
	LogFormat        string
}

// config object
func GetConfig() config {
	bootstrapServers := getEnv("KAFKA_BOOTSTRAP_SERVERS", "")
	if bootstrapServers == "" {
		host := getEnv("KAFKA_HOST", "localhost")
		port := getEnv("KAFKA_PORT", "9092")
		bootstrapServers = host + ":" + port
	}
	return config{
		BootStrapServers: strings.ToLower(bootstrapServers),
		WebSocketServer:  strings.ToLower(getEnv("WEBSOCKET_SERVER", "wss://localhost/echo")),
		KafkaTopic:       os.Getenv("KAFKA_TOPIC"),
		LogLevel:         strings.ToLower(getEnv("LOG_LEVEL", "info")),
		LogFormat:        strings.ToLower(getEnv("LOG_FORMAT", "text")), //cann be text or json
	}
}

func getEnv(key string, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
