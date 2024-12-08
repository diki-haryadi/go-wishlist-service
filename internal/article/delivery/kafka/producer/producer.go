package articleKafkaProducer

import (
	"context"

	"github.com/segmentio/kafka-go"

	articleDomain "github.com/diki-haryadi/go-micro-template/internal/article/domain"
	kafkaProducer "github.com/diki-haryadi/ztools/kafka/producer"
)

type producer struct {
	createWriter *kafkaProducer.Writer
}

func NewProducer(w *kafkaProducer.Writer) articleDomain.KafkaProducer {
	return &producer{createWriter: w}
}

func (p *producer) PublishCreateEvent(ctx context.Context, messages ...kafka.Message) error {
	return p.createWriter.Client.WriteMessages(ctx, messages...)
}
