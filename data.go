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
	if data == nil {
		return ""
	}

	// Add another one of these because of some weirdnesss with the app's state
	// after Heroku puts this to sleep ..
	rand.NewSource(time.Now().UnixNano())

	return data.Questions[rand.Intn(len(data.Questions))]
}
