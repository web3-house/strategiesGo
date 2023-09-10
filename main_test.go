package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/This-Is-Prince/strategiesGo/routes"
	"github.com/This-Is-Prince/strategiesGo/utils"
)

func TestPingRoute(t *testing.T) {
	// Create a test HTTP server
	r := routes.SetupRoutes(utils.NewClients())
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Make an HTTP GET request to the /ping route
	res, err := http.Get(ts.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	// Check the response status code
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}

	// Check the response body
	expectedBody := `{"message":"pong"}`
	actualBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(actualBody) != expectedBody {
		t.Errorf("Expected response body %s but got %s", expectedBody, string(actualBody))
	}
}

func TestScores(t *testing.T) {
	// Create a test HTTP server
	r := routes.SetupRoutes(utils.NewClients())
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Read the JSON data from the file
	jsonData, err := ioutil.ReadFile("scores_test.json")
	if err != nil {
		t.Fatal(err)
	}

	// Parse the JSON data into a map
	var testData map[string]interface{}
	if err := json.Unmarshal(jsonData, &testData); err != nil {
		t.Fatal(err)
	}

	// Extract the "input" object from the test data
	inputData, ok := testData["input"].(map[string]interface{})
	if !ok {
		t.Fatal("Failed to extract input data from JSON")
	}

	// Convert the input data to JSON
	inputJSON, err := json.Marshal(inputData)
	if err != nil {
		t.Fatal(err)
	}

	// Create an HTTP POST request with the JSON payload
	req, err := http.NewRequest("POST", ts.URL+"/scores", bytes.NewBuffer(inputJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Make the HTTP request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	// Check the response status code
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}

	// Read the response body
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Extract the "input" object from the test data
	expectedOutput, ok := testData["output"].([]interface{})
	if !ok {
		t.Fatal("Failed to extract input data from JSON")
	}

	// Unmarshal the actual response body
	var actualOutput []interface{}
	if err := json.Unmarshal(responseBody, &actualOutput); err != nil {
		t.Fatal(err)
	}

	// Compare the expected and actual output
	if !reflect.DeepEqual(expectedOutput, actualOutput) {
		t.Errorf("Expected response body %v but got %v", expectedOutput, actualOutput)
	}
}
