package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	URL = "http://127.0.0.1:8080/producers"
)

type Producer struct {
	Id           int32   `json:id`
	Manufacturer *string `json:manufacturer`
	Price        int32   `json:price`
}

func main() {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Printf("Errror when calling server %s\n", URL, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var producers []Producer
	err = json.Unmarshal(body, &producers)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for _, producer := range producers {
		if producer.Manufacturer != nil {
			fmt.Printf("Producer Producer%d with price %d and manufacturer %v\n", producer.Id, producer.Price, *producer.Manufacturer)
		} else {
			fmt.Printf("Producer Producer%d with price %d and no manufacturer\n", producer.Id, producer.Price)
		}
	}
}
