package ens

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var ENS_GRAPH_URL = map[string]string{
	"1": "https://api.thegraph.com/subgraphs/name/ensdomains/ens",
	"3": "https://api.thegraph.com/subgraphs/name/ensdomains/ensropsten",
	"4": "https://api.thegraph.com/subgraphs/name/ensdomains/ensrinkeby",
	"5": "https://api.thegraph.com/subgraphs/name/ensdomains/ensgoerli",
}

type Args struct {
	Query     string
	Variables map[string]interface{}
	Network   string
}

func FetchENS[T any](args Args, res *T) error {
	// ENS URL check
	url, ok := ENS_GRAPH_URL[args.Network]
	if !ok {
		return errors.New("ENS url does not exist")
	}

	// Create the GraphQL request body
	requestBody := map[string]interface{}{
		"query":     args.Query,
		"variables": args.Variables,
	}

	// Convert the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	// Make the HTTP POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Define a struct to represent the GraphQL response structure
	type GraphQLResponse struct {
		Data   T `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	// Parse the GraphQL response
	var graphqlResponse GraphQLResponse
	err = json.Unmarshal(responseBody, &graphqlResponse)
	if err != nil {
		return err
	}

	// Check for GraphQL errors
	if len(graphqlResponse.Errors) > 0 {
		return fmt.Errorf("GraphQL error: %s", graphqlResponse.Errors[0].Message)
	}

	*res = graphqlResponse.Data

	return nil
}
