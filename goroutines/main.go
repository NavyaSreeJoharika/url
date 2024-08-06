package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go receiveOrders(&wg)
	wg.Wait()
	fmt.Println(orders)
}

func receiveOrders(wg *sync.WaitGroup) {
	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
			continue
		}
		orders = append(orders, newOrder)
	}
	wg.Done()
}

var rawOrders = []string{
	`{"productCode":1111, "quality":5, "status":1}`,
	`{"productCode":2222, "quality":42.3, "status":1}`,
	`{"productCode":3333, "quality":19, "status":1}`,
	`{"productCode":4444, "quality":8, "status":1}`,
}

var orders = []order{}

// go run main.go order.go
