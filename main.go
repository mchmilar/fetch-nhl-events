package main

import (
	"encoding/json"
	//"fmt"
	"log"
	"net/http"
	//"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"strconv"
)

var db *sql.DB
var err error

// channels
var token = make(chan struct{}, 100)
var wait = make(chan struct{})
var urls = make(chan string)
var records = make(chan *NHL)
var events = make(chan *AllPlay)
var recordToken = make(chan struct{}, 100)

func main() {
	// Setup DB Connection
	db, err = sql.Open("mysql", "root:agrajag@/nhl")
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(1000)
	if err != nil {
		panic(err.Error())
	}
	err= db.Ping()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// URLS
	go func() {
		for gameId := 2016020001; gameId < 2016021231; gameId++ {
			url := "http://statsapi.web.nhl.com/api/v1/game/" + strconv.Itoa(gameId) + "/feed/live/"
			fmt.Println(url)
			urls <- url
		}
		close(urls)
	}()
	fmt.Println("Inbetween URLS and Fethcer")
	// Fetcher
	go func() {
		fmt.Println("Start of Fetcher")
		for {
			url, ok := <- urls
			if !ok {
				fmt.Println("not ok")
				break
			}
			recordToken <- struct{}{}
			go func() {
				records <- fetch(url)
				<-recordToken
			}()
		}
		fmt.Println("Closed?")
		close(records)
	}()
	fmt.Println("Inbetween Fetcher and Events")

	// Events
	go func() {
		for record := range records{
			for _, allPlay := range record.LiveData.Plays.AllPlays {
				events <-&allPlay
			}
		}
		close(events)
	}()

	// DB
	go func() {
		for event := range events {
			token <-struct{}{}
			go func() {
				_, err = db.Exec("INSERT INTO event (event_type_id) VALUES (?)", event.Result.EventTypeID)
				if err != nil {
					panic(err.Error())
				}
				<-token
			}()


		}
		wait <-struct{}{}
	}()


	<-wait
}

func fetch(url string) *NHL {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record NHL
	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	return &record
}
