package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/gorilla/mux"
)

type Job struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Salary      string `json:"salary"`
}

// this file is for getting data to the kafka

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/jobs", createjob).Methods("POST")
	log.Fatal(http.ListenAndServe(":9090", router))
}

func createjob(w http.ResponseWriter, r *http.Request) {
	//receive body from http request
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	//save data into job struct
	var _job Job
	err = json.Unmarshal(b, &_job)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	saveJobToKafka(_job)

	//convert job struct to json

	jsonString, err := json.Marshal(_job)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//set content-type http header
	w.Header().Set("content-type", "application/json")

	//send back data as response
	w.Write(jsonString)
}

func saveJobToKafka(job Job) {

	fmt.Println("save to kafka")

	jsonString, err := json.Marshal(job)

	jobString := string(jsonString)
	fmt.Println(jobString)

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)

	}
	//produce message to topic
	topic := "job-topic1"
	for _, word := range []string{string(jobString)} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(0)},
			Value:          []byte(word),
		}, nil)
	}
}

//kafka.PartitionAny
