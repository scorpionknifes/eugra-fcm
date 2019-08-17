package main

import (
	"log"

	"github.com/appleboy/go-fcm"
)

func fcmsend() {

	// Create the message to be sent.
	msg := &fcm.Message{
		To:           "dDuANNNc6YY:APA91bF82k21H-TNuDa-IlAY56EWmxJeH987RFGHAoxxZO2OTBObfnCiarwU6ZtrgPt0p7VTKUDIT1Ts0gbOdtSFOZPadRDLvLgvkgB49RltTvaXRddnNAHuUHd5adBQxMLNahgUcaA7",
		Notification: &fcm.Notification{Title: "JIMMY" + " has Grab a new Table", Body: "TOPIC TITLE"},
		Data: map[string]interface{}{
			"table_id": 10,
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
