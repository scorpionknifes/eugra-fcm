package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/appleboy/go-fcm"
)

type Status struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SendFollowers(w http.ResponseWriter, r *http.Request) {
	log.Print("./SendFollowers")
	id, ok := r.URL.Query()["id"]
	if !ok || len(id[0]) < 1 {
		response := Status{
			Status:  "400",
			Message: "Url Param 'id' is missing",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	var display_name string
	var topic_id int64
	var table_master_id int64
	var topic_name string

	row := SQL.QueryRow("SELECT users.id, users.name, tables.topic_id FROM eugra_users as users INNER JOIN eugra_tables as tables ON users.id = tables.table_master_id WHERE tables.id =?", id[0])
	row.Scan(&table_master_id, &display_name, &topic_id)
	row = SQL.QueryRow("SELECT `title` FROM `eugra_topics` WHERE `id`=?", topic_id)
	row.Scan(&topic_name)
	rows, _ := SQL.Query("SELECT `follower_id` FROM `eugra_follows` WHERE `user_id`=?", table_master_id)
	for rows.Next() {
		var followerid int64
		rows.Scan(&followerid)
		log.Println(followerid)
		rows, _ := SQL.Query("SELECT `token` FROM `eugra_firebase` WHERE `user_id`=?", followerid)
		for rows.Next() {
			var token string
			rows.Scan(&token)
			fcmMessage(token, display_name, topic_name)
		}
	}
}

func fcmMessage(token string, display_name string, topic_name string) {
	// Create the message to be sent.
	msg := &fcm.Message{
		To:           token,
		Notification: &fcm.Notification{Title: display_name + " has Grab a new Table", Body: topic_name},
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
