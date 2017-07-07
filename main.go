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



func main() {
	//url := os.Args[1]
	// Build the request
	db, err = sql.Open("mysql", "root:agrajag@/nhl")
	if err != nil {
		panic(err.Error())
	}
	err= db.Ping()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	ch := make(chan NHL)
	for gameId := 2016020001; gameId < 2016021231; gameId++ {
		url := "http://statsapi.web.nhl.com/api/v1/game/" + strconv.Itoa(gameId) + "/feed/live/"
		fmt.Println(url)
		go fetch(url, ch)
	}

	for gameId := 2016020001; gameId < 2016021231; gameId++ {
		// handle channel return
		record := <-ch
		for _, element := range record.LiveData.Plays.AllPlays {
			/*fmt.Println("Event id: ", element.About.EventID)
			fmt.Println("eventIdx: ", element.About.EventIdx)
			fmt.Println("Description: ", element.Result.Description)
			fmt.Println("Event: ", element.Result.Event)
			fmt.Println("eventCode: ", element.Result.EventCode)
			fmt.Println("eventTypeId: ", element.Result.EventTypeID, "\n")*/
			_, err = db.Exec("INSERT INTO event (event_type_id) VALUES (?)", element.Result.EventTypeID)
			if err != nil {
				panic(err.Error())
			}
		}
	}
}

func fetch(url string, ch chan<- NHL) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
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
		return
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
	ch<- record
}
