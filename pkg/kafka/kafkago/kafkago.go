package kafkago

import (
	"ClientWorkerService/pkg/helper"
	"context"
	"time"

	"github.com/appneuroncompany/light-logger/clogger"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

type KafkaGo struct {
}

func (k *KafkaGo) Produce(key *[]byte, value *[]byte, topic string) (err error) {

	writer, _ := writerConfigure([]string{helper.ResolvePath("KAFKA_HOST", "KAFKA_PORT")}, uuid.New().String(), topic)
	message := kafka.Message{
		Key:   *key,
		Value: *value,
		Time:  time.Now(),
	}
	err = writer.WriteMessages(context.Background(), message)
	return err
}

func (k *KafkaGo) Consume(topic string, groupId string, callback func(topic string, data []byte) error) {
	reader, _ := readerConfigure([]string{helper.ResolvePath("KAFKA_HOST", "KAFKA_PORT")}, groupId, topic)
	defer func(reader *kafka.Reader) {
		err := reader.Close()
		if err != nil {
			clogger.Error(&map[string]interface{}{
				"KafkaGo consume err message ": err,
			})
		}
	}(reader)
	clogger.Info(&map[string]interface{}{
		"KafkaGo consume message ": reader.Stats().ClientID,
	})

	for {
		m, err := reader.FetchMessage(context.Background())
		if err != nil {
			clogger.Error(&map[string]interface{}{
				"KafkaGo consume error while receiving message ": err,
			})
			continue
		}
		callErr := callback(topic, m.Value)
		if callErr == nil {
			if err := reader.CommitMessages(context.Background(), m); err != nil {
				clogger.Error(&map[string]interface{}{
					"KafkaGo failed to commit messages": err,
				})
			}
		}
	}
}

func writerConfigure(kafkaBrokerUrls []string, clientId string, topic string) (w *kafka.Writer, err error) {
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
	return w, nil
}

func readerConfigure(kafkaBrokerUrls []string, groupID string, topic string) (r *kafka.Reader, err error) {
	config := kafka.ReaderConfig{
		Brokers:         kafkaBrokerUrls,
		GroupID:         groupID,
		Topic:           topic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka_go.
		ReadLagInterval: -1,
	}

	reader := kafka.NewReader(config)
	return reader, nil
}
