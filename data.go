package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

type DataStore struct {
	Questions []string
}

var (
	data *DataStore
)

func init() {
	// Read in the questions
	f, err := os.Open("data/questions.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	// Seed the random generator
	rand.NewSource(time.Now().UnixNano())
}

func getQuestion() string {
	return data.Questions[rand.Intn(len(data.Questions))]
}
