package consumer

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/hse-telescope/emailer/internal/providers/email"
	"github.com/hse-telescope/emailer/pkg/wrapper"

	"github.com/hse-telescope/utils/queues/kafka"

	"github.com/IBM/sarama"
)

var (
	ErrNoMessages   = errors.New("no messages")
	ErrMessageParse = errors.New("failed to parse message")
)

// Consumer ...
type Consumer struct {
	emailProvider     email.Provider
	partitionConsumer sarama.PartitionConsumer
}

// New ...
func New(
	emailProvider email.Provider,
	cfg kafka.QueueCredentials,
) (Consumer, error) {
	saramaConfig := sarama.NewConfig()
	saramaConsumer, err := sarama.NewConsumer(
		cfg.URLs,
		saramaConfig,
	)
	if err != nil {
		return Consumer{}, err
	}

	partitionConsumer, err := saramaConsumer.ConsumePartition(cfg.Topic, 0, sarama.OffsetOldest)
	if err != nil {
		return Consumer{}, err
	}

	return Consumer{
		emailProvider:     emailProvider,
		partitionConsumer: partitionConsumer,
	}, nil
}

func (c Consumer) Consume(ctx context.Context) error {
	for message := range c.partitionConsumer.Messages() {
		// logger.Info("got message", zap.Any("message", message))

		decodedMessage := wrapper.Message{}
		err := json.Unmarshal(message.Value, &decodedMessage)
		if err != nil {
			// logger.Error("failed to decode message", zap.Any("message", message.Value), zap.Error(err))
			return ErrMessageParse
		}

		err = c.emailProvider.SendEmail(ctx, decodedMessage.EMail, email.WrapperMessageToProviderMessage(decodedMessage))
		if err != nil {
			// logger.Error("failed to send email", zap.Error(err))
		}
	}

	// logger.Error("channel closed")
	return ErrNoMessages
}
