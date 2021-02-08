package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
	"github.com/twsm000/imersao-fullstack-fullcycle/application/factory"
	appmodel "github.com/twsm000/imersao-fullstack-fullcycle/application/model"
	"github.com/twsm000/imersao-fullstack-fullcycle/application/usecase"
	"github.com/twsm000/imersao-fullstack-fullcycle/domain/model"
)

// Processor ...
type Processor struct {
	Database     *gorm.DB
	Producer     *ckafka.Producer
	DeliveryChan chan ckafka.Event
}

// NewKafkaProcessor ...
func NewKafkaProcessor(database *gorm.DB, producer *ckafka.Producer, delivery chan ckafka.Event) *Processor {
	return &Processor{
		Database:     database,
		Producer:     producer,
		DeliveryChan: delivery,
	}
}

// Consume ...
func (p *Processor) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
		"group.id":          os.Getenv("kafkaConsumerGroupId"),
		"auto.offset.reset": "earliest",
	}

	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalln(err)
	}

	topics := []string{
		os.Getenv("kafkaTransactionTopic"),
		os.Getenv("kafkaTransactionConfirmationTopic"),
	}
	c.SubscribeTopics(topics, nil)
	log.Println("kafka consumer has been started")
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			p.processMessage(msg)
		}
	}
}

func (p *Processor) processMessage(msg *ckafka.Message) {
	transactionsTopic := "transactions"
	transactionConfirmationTopic := "transaction_confirmation"

	switch topic := *msg.TopicPartition.Topic; topic {
	case transactionsTopic:
		p.processTransaction(msg)
	case transactionConfirmationTopic:
		p.processTransactionConfirmation(msg)
	default:
		fmt.Println("not a valid topic", string(msg.Value))
	}
}

func (p *Processor) processTransaction(msg *ckafka.Message) error {
	t := appmodel.NewTransaction()
	if err := t.ParseJSON(msg.Value); err != nil {
		return err
	}

	uc := factory.TransactionUseCaseFactory(p.Database)
	newTransaction, err := uc.Register(
		t.AccountID,
		t.Amount,
		t.PixKeyTo,
		t.PixKeyKindTo,
		t.Description,
	)
	if err != nil {
		return fmt.Errorf("error registering transaction %v", err)
	}

	topic := "bank" + newTransaction.PixKeyTo.Account.Bank.Code
	t.ID = newTransaction.ID
	t.Status = string(newTransaction.Status)
	jsonData, err := t.ToJSON()
	if err != nil {
		return err
	}

	return Publish(string(jsonData), topic, p.Producer, p.DeliveryChan)
}

func (p *Processor) processTransactionConfirmation(msg *ckafka.Message) error {
	t := appmodel.NewTransaction()
	var err error
	err = t.ParseJSON(msg.Value); 
	if err != nil {
		return err
	}

	uc := factory.TransactionUseCaseFactory(p.Database)

	
	if t.Status == string(model.TransactionConfirmed) {
		err = p.confirmTransaction(t, uc)			
	} else if t.Status == string(model.TransactionCompleted) {
		_, err = uc.Complete(t.ID)
	}

	return err
}

func (p *Processor) confirmTransaction(transaction *appmodel.Transaction, uc *usecase.TransactionUseCase) error {
	confirmedTransaction, err := uc.Confirm(transaction.ID)
	if err != nil {
		return err
	}

	jsonData, err := transaction.ToJSON()
	if err != nil {
		return err
	}

	topic := "bank" + confirmedTransaction.AccountFrom.Bank.Code
	return Publish(string(jsonData), topic, p.Producer, p.DeliveryChan)
}
