package kafka

import (
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/segmentio/kafka-go"
)

type Reader struct {
	Client *kafka.Reader
}

type ReaderConf struct {
	Brokers []string
	GroupID string
	Topic   string
}

func NewKafkaReader(cfg *ReaderConf) *Reader {
	rc := kafka.ReaderConfig{
		Brokers:                cfg.Brokers,
		GroupID:                cfg.GroupID,
		Topic:                  cfg.Topic,
		QueueCapacity:          queueCapacity,
		MinBytes:               minBytes,
		MaxBytes:               maxBytes,
		HeartbeatInterval:      heartbeatInterval,
		CommitInterval:         commitInterval,
		PartitionWatchInterval: partitionWatchInterval,
		Logger:                 kafka.LoggerFunc(logger.Zap.Sugar().Debugf),
		ErrorLogger:            kafka.LoggerFunc(logger.Zap.Sugar().Errorf),
		MaxAttempts:            maxAttempts,
		Dialer: &kafka.Dialer{
			Timeout: dialTimeout,
		},
	}
	return &Reader{
		Client: kafka.NewReader(rc),
	}
}
