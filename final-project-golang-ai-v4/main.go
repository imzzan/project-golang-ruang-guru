package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type AIModelConnector struct {
	Client *http.Client
}

type Inputs struct {
	Table map[string][]string `json:"table"`
	Query string              `json:"query"`
}

type Response struct {
	Answer      string   `json:"answer"`
	Coordinates [][]int  `json:"coordinates"`
	Cells       []string `json:"cells"`
	Aggregator  string   `json:"aggregator"`
}

func CsvToSlice(data string) (map[string][]string, error) {
	reader := csv.NewReader(strings.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("not enough data in CSV")
	}

	headers := records[0]
	columns := make(map[string][]string)

	for _, header := range headers {
		columns[header] = []string{}
	}

	for _, record := range records[1:] {
		for i, value := range record {
			columns[headers[i]] = append(columns[headers[i]], value)
		}
	}

	return columns, nil
}

func (c *AIModelConnector) ConnectAIModel(payload interface{}, token string) (Response, error) {
	url := "https://api-inference.huggingface.co/models/google/tapas-base-finetuned-wtq"
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return Response{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return Response{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return Response{}, err
	}

	return result, nil
}

func main() {
	csvFile, err := os.Open("data-series.csv")
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer csvFile.Close()

	csvData, err := ioutil.ReadAll(csvFile)
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	data, err := CsvToSlice(string(csvData))
	if err != nil {
		log.Fatalf("Failed to parse CSV data: %v", err)
	}

	connector := &AIModelConnector{Client: &http.Client{}}
	token := "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"

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter your query (type 'exit' to quit):")
		query, _ := reader.ReadString('\n')
		query = strings.TrimSpace(query)

		if query == "exit" {
			break
		}

		payload := Inputs{
			Table: data,
			Query: query,
		}

		response, err := connector.ConnectAIModel(payload, token)
		if err != nil {
			fmt.Printf("Error connecting to AI model: %v\n", err)
			continue
		}

		fmt.Printf("Answer: %s\n", response.Answer)
		fmt.Printf("Details: %+v\n", response)
	}
}
