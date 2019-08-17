package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

const local string = "LOCAL"

//SQL & MQTT settings
var SQL *sql.DB

//Function to Start APP
func main() {
	var (
		// environment variables
		env        = os.Getenv("ENV")  // LOCAL, DEV, STG, PRD
		port       = os.Getenv("PORT") // server traffic on this port
		connection = os.Getenv("SQL")
	)
	if env == "" || env == local {
		env = local
		port = "3004"
		connection = "eugra:eugra@tcp(114.23.220.120:3306)/eugra-api?parseTime=true"
	}
	//Initialise Database Connection
	var errs error
	SQL, errs = sql.Open("mysql", connection)
	if errs != nil {
		log.Fatal(errs)
	}
	defer SQL.Close()

	StartAPI(port)
}

//Start Simple REST API
func StartAPI(port string) {
	router := mux.NewRouter()

	//Handlers for REST API
	router.HandleFunc("/", SendFollowers)

	srv := &http.Server{
		Addr:    ":" + port, //Connection Port
		Handler: router,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	//Check REST API Connection
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	//Function to Stop Server
	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
