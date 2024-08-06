package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	receiveOrders()
	fmt.Println(orders)
}

func receiveOrders() {
	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
			continue
		}
		orders = append(orders, newOrder)
	}
}

var rawOrders = []string{
	`{"productCode":1111, "quality":5, "status":1}`,
	`{"productCode":2222, "quality":42.3, "status":1}`,
	`{"productCode":3333, "quality":19, "status":1}`,
	`{"productCode":4444, "quality":8, "status":1}`,
}

//this is synchronously done program
