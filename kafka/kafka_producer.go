package kafka

import (
	"github.com/Shopify/sarama"
	"log"
	"strings"
	"time"
)

type ProducerSetting struct {
	Url    []string
	Option struct {
	}
}

type Channel struct {
	Url   []string
	Topic string
}
type Client struct {
	Url    []string
	config *sarama.Config
}

func NewClient(bootstrapServers string, config *sarama.Config) *Client {
	client := &Client{}
	client.Url = strings.Split(bootstrapServers, ",")
	if config != nil {
		client.config = config
	} else {
		defaultConfig := sarama.NewConfig()
		defaultConfig.Producer.RequiredAcks = sarama.WaitForLocal
		defaultConfig.Producer.Partitioner = sarama.NewRandomPartitioner
		defaultConfig.Producer.Return.Successes = true
		client.config = defaultConfig
	}
	return client
}

func (client *Client) SendMessage(topic, key, message string) error {

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Key = sarama.StringEncoder(key)
	msg.Timestamp = time.Now()
	msg.Value = sarama.ByteEncoder(message)

	producer, err := sarama.NewSyncProducer(client.Url, client.config)
	if err != nil {
		log.Println("FAILED TO INIT KAFKA PRODUCER:", err)
		return err
	}
	defer producer.Close()

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Println("FAILED TO SEND KAFKA MSG,", err)
		return err
	}
	//fmt.Printf("pid:%v offset:%v\n", pid, offset)
	return nil
}
