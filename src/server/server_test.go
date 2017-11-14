package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPost(t *testing.T) {
	// Set up server
	server := httptest.NewServer(WordHandler{})
	defer server.Close()

	// Set up test cases
	testCases := []struct {
		testCaseName   string
		text           string
		limit          int
		expectedOutput string
	}{
		{"Simple case", "a a a", -1, "{\"word_list\":[{\"text\":\"a\",\"frequency\":3}]}"},
		{"Testing defaults", "a b c d e f g h i j k l", -1, "{\"word_list\":[{\"text\":\"a\"," +
			"\"frequency\":1},{\"text\":\"b\",\"frequency\":1},{\"text\":\"c\",\"frequency\":1},{\"text\":\"d\",\"frequency\":1},{" +
			"\"text\":\"e\",\"frequency\":1},{\"text\":\"f\",\"frequency\":1},{\"text\":\"g\",\"frequency\":1},{\"text\":\"h\"" +
			",\"frequency\":1},{\"text\":\"i\",\"frequency\":1},{\"text\":\"j\",\"frequency\":1}]}"},
	}

	for _, singleTestCase := range testCases {
		// If limit is less than 0, assume that it was not given in the request
		url := server.URL
		if singleTestCase.limit >= 0 {
			url = url + "?q=" + string(singleTestCase.limit)
		}

		// Create the required JSON
		requestValue := WordRequest{singleTestCase.text}
		requestJSON, err := json.Marshal(requestValue)
		if err != nil {
			t.Fatal("Stop all tests. Error with producing JSON")
		}

		// Retrieve and check status code
		response, err := http.Post(url, "application/json", bytes.NewReader(requestJSON))
		if err != nil {
			t.Error(singleTestCase.testCaseName, "Unable to obtain the required response")
		}
		if response.StatusCode != http.StatusOK {
			t.Error(singleTestCase.testCaseName, "Did not obtain the required status")
		}

		// Checking data
		data, err := ioutil.ReadAll(response.Body)
		if strings.TrimSpace(string(data)) != singleTestCase.expectedOutput {
			t.Error(singleTestCase.testCaseName, "failed. Expected:", singleTestCase.expectedOutput, "Actual:", string(data))
		}

	}
}
