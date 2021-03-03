package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Quiz struct {
	Id           int       `json:"id"`
	Question     string    `json:"question"`
	Propositions [4]string `json:"propositions"`
	Réponse      string    `json:"réponse"`
	Anecdote     string    `json:"anecdote"`
}

var quizzs []Quiz

func getRandomQuiz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	question := quizzs[rand.Intn(len(quizzs))]

	json.NewEncoder(w).Encode(question)
}

func getQuizByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range quizzs {
		toCompare, _ := strconv.Atoi(params["id"])
		if item.Id == toCompare {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
}

func readFile(path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &quizzs)
}

func main() {

	readFile("quizzs.json")

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/quiz/", getRandomQuiz).Methods("GET")
	r.HandleFunc("/quiz/{id}", getQuizByID).Methods("GET")

	port := os.Getenv("PORT")

    	if port == "" {
        	port = "8000"
        	log.Printf("defaulting to port %s\n", port)
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
