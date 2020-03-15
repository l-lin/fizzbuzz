package kafka

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/l-lin/fizzbuzz/stats"
	"github.com/l-lin/fizzbuzz/stats/memory"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

// Repository that produces request stats to Kafka
type Repository struct {
	memoryRepository *memory.Repository
	*kafka.Writer
	*kafka.Reader
}

// NewRepository returns a new instance the repository to store stats in memory
func NewRepository(brokers, topic string) *Repository {
	return &Repository{
		memoryRepository: memory.NewRepository(),
		Writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers:  strings.Split(brokers, ","),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		}),
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  strings.Split(brokers, ","),
			Topic:    topic,
			MinBytes: 10e3, // 10KB
			MaxBytes: 10e6, // 10MB
		}),
	}
}

// GetAll stats from Kafka
func (r *Repository) GetAll() []*stats.Request {
	return r.memoryRepository.GetAll()
}

// Increment number of hits in Kafka
func (r *Repository) Increment(path string, parameters map[string]interface{}) error {
	parent := context.Background()
	defer parent.Done()
	req := map[string]interface{}{}
	req["path"] = path
	if parameters != nil {
		req["parameters"] = parameters
	}
	v, err := json.Marshal(req)
	if err != nil {
		return err
	}
	msg := kafka.Message{
		Key:   nil,
		Value: v,
		Time:  time.Now(),
	}
	log.Debug().Str("value", string(v)).Msg("Producing message to Kafka")
	return r.Writer.WriteMessages(parent, msg)
}

// Close kafka writer & reader
func (r *Repository) Close() error {
	log.Info().Msg("Closing connection to Kafka")
	return r.Writer.Close()
}

func (r *Repository) Listen() {
	for {
		parent := context.Background()
		defer parent.Done()
		msg, err := r.Reader.ReadMessage(parent)
		if err != nil {
			log.Err(err)
			continue
		}
		log.Debug().Str("value", string(msg.Value)).Msg("Consumed message from Kafka")
		req := map[string]interface{}{}
		if err := json.Unmarshal(msg.Value, &req); err != nil {
			log.Err(err)
			continue
		}
		path, ok := req["path"].(string)
		if !ok {
			log.Error().Msg("No path found in kafka message")
			continue
		}
		// ignore errors because some endpoints do not have any parameters
		parameters, _ := req["parameters"].(map[string]interface{})
		r.memoryRepository.Increment(path, parameters)
	}
}
