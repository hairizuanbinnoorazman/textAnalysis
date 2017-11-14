package main

import (
	"analyse"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type WordRequest struct {
	Text string
}

type WordResponse struct {
	WordList []analyse.Word `json:"word_list"`
}

type WordHandler struct{}

func (h WordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	queryLimit := queryValues.Get("q")

	// Dealing with retrieving the limit value
	var limit int
	if queryLimit == "" {
		limit = 10
	}
	limit, err := strconv.Atoi(queryLimit)
	if err != nil {
		log.Println("Unable to convert the limit accordingly. Will go back to defaut")
		limit = 10
	}

	// Dealing with retrieving the request body
	rawData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read the raw data accordingly")
	}
	data := WordRequest{}
	json.Unmarshal(rawData, &data)

	// Send it off for analysis
	results := analyse.TopNwords(data.Text, limit)
	response := WordResponse{results}

	// Encode and return
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func main() {
	port := ":3000"
	log.Println("Server Started on port", port)

	http.Handle("/analysetext", WordHandler{})
	log.Fatal(http.ListenAndServe(port, nil))
}
