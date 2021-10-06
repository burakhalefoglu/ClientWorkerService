package kafka

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

var writer *kafka.Writer

func Configure(kafkaBrokerUrls []string, clientId string, topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientId,
	}

	config := kafka.WriterConfig{
		Brokers:          kafkaBrokerUrls,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	w = kafka.NewWriter(config)
	writer = w
	return w, nil
}

func Produce(parent context.Context, key, value []byte, topic string) (err error) {
	Configure([]string{"192.168.43.178:9092"}, uuid.New().String(), topic)
	message := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}
	err = writer.WriteMessages(parent, message)
	return err
}