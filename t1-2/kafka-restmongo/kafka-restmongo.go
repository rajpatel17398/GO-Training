package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gopkg.in/mgo.v2"
)

const (
	hosts      = "localhost:27017"
	database   = "kafka"
	username   = ""
	password   = ""
	collection = "kafka"
)

type Job struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Salary      string `json:"salary"`
}
type MongoStore struct {
	session *mgo.Session
}

var mongoStore = MongoStore{}

func main() {
	//create mongoDB session
	session := initialiseMongo()
	mongoStore.session = session

	receiveFromKafka()
}
func initialiseMongo() (session *mgo.Session) {

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}
	return

}

func receiveFromKafka() {
	fmt.Println("Start receiveing from kafka")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "group-id-1",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)

	}
	c.SubscribeTopics([]string{"job-topic1"}, nil)

	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			fmt.Printf("Received from Kafka %s: %s\n", msg.TopicPartition, string(msg.Value))
			job := string(msg.Value)
			saveJobToMongo(job)

		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}

	}
	c.Close()
}
func saveJobToMongo(jobString string) {
	fmt.Println("save to mongoDB")
	col := mongoStore.session.DB(database).C(collection)

	//save data into job struct
	var _job Job
	b := []byte(jobString)
	err := json.Unmarshal(b, &_job)
	if err != nil {
		panic(err)

	}
	// insert job into mongoDB

	errMongo := col.Insert(_job)
	if errMongo != nil {
		panic(errMongo)

	}
	fmt.Printf("saved to MongoDB: %s", jobString)
}
