package main

import (
	"fmt"
	"log"

	"net/http"
	"net/http/httputil"
	"time"

	"github.com/Shopify/sarama"
	"github.com/matzew/ws-kafka/pkg/config"
)

func bridgeToKafka(message string) {

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

func dumpToKafka(w http.ResponseWriter, r *http.Request) {
	// Simulate at least a bit of processing time.
	time.Sleep(100 * time.Millisecond)

	w.WriteHeader(http.StatusOK)
	if reqBytes, err := httputil.DumpRequest(r, true); err == nil {
		log.Printf("Openshift Http Request Dumper received a message: %+v", string(reqBytes))
		bridgeToKafka(string(reqBytes))
		w.Write(reqBytes)
	} else {
		log.Printf("Error dumping the request: %+v :: %+v", err, r)
	}
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", dumpToKafka)

	http.ListenAndServe(":8080", m)
}
