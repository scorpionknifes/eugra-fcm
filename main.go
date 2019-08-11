package main

import (
	"log"

	"github.com/appleboy/go-fcm"
)

func main() {
	// Create the message to be sent.
	msg := &fcm.Message{
		To: "sample_device_token",
		Data: map[string]interface{}{
			"foo": "bar",
		},
	}

	// Create a FCM client to send the message.
	client, err := fcm.NewClient("AAAA1YSlFuQ:APA91bGjAbnVhpFYWLfr2l2qafYpsFJhAbIl5eGvcew-eWf1w0myBwrP8TPVI0c2czNLqKEYWvXM3Y32vv0wb25_nNxcIWgvJaSzIQD3qn5KE_6ywQtX599mXb61aKUwA_bbx7CBIzMv")
	if err != nil {
		log.Fatalln(err)
	}

	// Send the message and receive the response without retries.
	response, err := client.Send(msg)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", response)
}
