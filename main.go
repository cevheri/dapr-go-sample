package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type incomingEvent struct {
	Data interface{} `json:"data"`
}

func main() {
	http.HandleFunc("/greeting",
		func(w http.ResponseWriter, r *http.Request) {
			var event incomingEvent
			decoder := json.NewDecoder(r.Body)
			decoder.Decode(&event)
			fmt.Println(event.Data)
			fmt.Fprint(w, "Hello World!")
		})
	log.Fatal(http.ListenAndServe(":8088", nil))
}

// dapr run --app-id hello-dapr --app-port 8088 --port 8089 go run main.go
// http://localhost:8088/greeting
// .../<version>/invoke/<action-id>/method/<methodname>
// http://localhost:8089/v1.0/invoke/hello-dapr/method/greeting