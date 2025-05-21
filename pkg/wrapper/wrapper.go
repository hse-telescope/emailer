package wrapper

import (
	"context"
	"encoding/json"

	"github.com/hse-telescope/utils/queues/kafka"

	"github.com/IBM/sarama"
)

type Emailer struct {
	topic    string
	producer sarama.AsyncProducer
}

func New(kafkaCfg kafka.QueueCredentials) (Emailer, error) {
	saramaConfig := sarama.NewConfig()
	saramaProducer, err := sarama.NewAsyncProducer(
		kafkaCfg.URLs,
		saramaConfig,
	)
	if err != nil {
		return Emailer{}, err
	}

	return Emailer{
		topic:    kafkaCfg.Topic,
		producer: saramaProducer,
	}, nil
}

type Message struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

const (
	defaultPartition = 0
)

func (e Emailer) SendEmail(ctx context.Context, message Message) error {
	messageEncoded, err := json.Marshal(message)
	if err != nil {
		return err
	}

	go func() {
		e.producer.Input() <- &sarama.ProducerMessage{
			Topic:     e.topic,
			Partition: defaultPartition,
			Offset:    sarama.OffsetOldest,
			Value:     sarama.ByteEncoder(messageEncoded),
		}
	}()
	return nil
}
