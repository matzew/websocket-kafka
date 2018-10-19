package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"github.com/matzew/ws-kafka/pkg/config"
	"github.com/sacOO7/gowebsocket"
)

func bridgeToKafka(message string, socket gowebsocket.Socket) {

	config := config.GetConfig()
	cfg := sarama.NewConfig()
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Retry.Max = 5
	cfg.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{config.BootStrapServers}, cfg)
	if err != nil {
		// Should not reach here
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			// Should not reach here
			panic(err)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: config.KafkaTopic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", config.KafkaTopic, partition, offset)

}

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	config := config.GetConfig()

	socket := gowebsocket.New(config.WebSocketServer)

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Fatal("Received connect error - ", err)
	}

	// attach event handler
	socket.OnTextMessage = bridgeToKafka

	socket.Connect()

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			socket.Close()
			return
		}
	}
}
