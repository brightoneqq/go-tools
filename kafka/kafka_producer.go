package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/brightoneqq/go-tools/flow"
	"log"
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

var CoreChannel = &Channel{}

func SendMessage(channel *Channel, key string, message string) error {
	//fmt.Println(message)
	return flow.Retry(3, 10, func() error {
		return SimpleSendMessage(channel, key, message)
	})
}

func SimpleSendMessage(channel *Channel, key string, message string) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = channel.Topic
	msg.Key = sarama.StringEncoder(key)
	msg.Timestamp = time.Now()
	msg.Value = sarama.ByteEncoder(message)

	client, err := sarama.NewSyncProducer(channel.Url, config)
	if err != nil {
		log.Println("FAILED TO INIT KAFKA CLIENT:", err)
		//todo: call http for this fail
		return err
	}
	defer client.Close()

	_, _, err = client.SendMessage(msg)
	if err != nil {
		log.Println("FAILED TO SEND KAFKA MSG,", err)
		return err
	}
	//fmt.Printf("pid:%v offset:%v\n", pid, offset)
	return nil
}
