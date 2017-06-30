package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Player struct {
	Active             bool   `json:"active"`
	AlternateCaptain   bool   `json:"alternateCaptain"`
	BirthCity          string `json:"birthCity"`
	BirthCountry       string `json:"birthCountry"`
	BirthDate          string `json:"birthDate"`
	BirthStateProvince string `json:"birthStateProvince"`
	Captain            bool   `json:"captain"`
	CurrentAge         int    `json:"currentAge"`
	CurrentTeam        struct {
		ID      int    `json:"id"`
		Link    string `json:"link"`
		Name    string `json:"name"`
		TriCode string `json:"triCode"`
	} `json:"currentTeam"`
	FirstName       string `json:"firstName"`
	FullName        string `json:"fullName"`
	Height          string `json:"height"`
	ID              int    `json:"id"`
	LastName        string `json:"lastName"`
	Link            string `json:"link"`
	Nationality     string `json:"nationality"`
	PrimaryNumber   string `json:"primaryNumber"`
	PrimaryPosition struct {
		Abbreviation string `json:"abbreviation"`
		Code         string `json:"code"`
		Name         string `json:"name"`
		Type         string `json:"type"`
	} `json:"primaryPosition"`
	Rookie        bool   `json:"rookie"`
	RosterStatus  string `json:"rosterStatus"`
	ShootsCatches string `json:"shootsCatches"`
	Weight        int    `json:"weight"`
}

type NHL struct {
	Copyright string `json:"copyright"`
	GameData  struct {
		DateTime struct {
			DateTime    string `json:"dateTime"`
			EndDateTime string `json:"endDateTime"`
		} `json:datetime`
		Game struct {
			Pk     int    `json:"pk"`
			Season string `json:"season"`
			Type   string `json:"type"`
		} `json:game`
		Players map[string]Player `json:"players"`
		Status struct {
			AbstractGameState string `json:"abstractGameState"`
			CodedGameState    string `json:"codedGameState"`
			DetailedState     string `json:"detailedState"`
			StartTimeTBD      bool   `json:"startTimeTBD"`
			StatusCode        string `json:"statusCode"`
		} `json:"status"`
		Teams struct {
			Away struct {
				Abbreviation string `json:"abbreviation"`
				Active       bool   `json:"active"`
				Conference   struct {
					ID   int    `json:"id"`
					Link string `json:"link"`
					Name string `json:"name"`
				} `json:"conference"`
				Division struct {
					ID   int    `json:"id"`
					Link string `json:"link"`
					Name string `json:"name"`
				} `json:"division"`
				FirstYearOfPlay string `json:"firstYearOfPlay"`
				Franchise       struct {
					FranchiseID int    `json:"franchiseId"`
					Link        string `json:"link"`
					TeamName    string `json:"teamName"`
				} `json:"franchise"`
				FranchiseID     int    `json:"franchiseId"`
				ID              int    `json:"id"`
				Link            string `json:"link"`
				LocationName    string `json:"locationName"`
				Name            string `json:"name"`
				OfficialSiteURL string `json:"officialSiteUrl"`
				ShortName       string `json:"shortName"`
				TeamName        string `json:"teamName"`
				TriCode         string `json:"triCode"`
				Venue           struct {
					City     string `json:"city"`
					Link     string `json:"link"`
					Name     string `json:"name"`
					TimeZone struct {
						ID     string `json:"id"`
						Offset int    `json:"offset"`
						Tz     string `json:"tz"`
					} `json:"timeZone"`
				} `json:"venue"`
			} `json:"away"`
			Home struct {
				Abbreviation string `json:"abbreviation"`
				Active       bool   `json:"active"`
				Conference   struct {
					ID   int    `json:"id"`
					Link string `json:"link"`
					Name string `json:"name"`
				} `json:"conference"`
				Division struct {
					ID   int    `json:"id"`
					Link string `json:"link"`
					Name string `json:"name"`
				} `json:"division"`
				FirstYearOfPlay string `json:"firstYearOfPlay"`
				Franchise       struct {
					FranchiseID int    `json:"franchiseId"`
					Link        string `json:"link"`
					TeamName    string `json:"teamName"`
				} `json:"franchise"`
				FranchiseID     int    `json:"franchiseId"`
				ID              int    `json:"id"`
				Link            string `json:"link"`
				LocationName    string `json:"locationName"`
				Name            string `json:"name"`
				OfficialSiteURL string `json:"officialSiteUrl"`
				ShortName       string `json:"shortName"`
				TeamName        string `json:"teamName"`
				TriCode         string `json:"triCode"`
				Venue           struct {
					City     string `json:"city"`
					Link     string `json:"link"`
					Name     string `json:"name"`
					TimeZone struct {
						ID     string `json:"id"`
						Offset int    `json:"offset"`
						Tz     string `json:"tz"`
					} `json:"timeZone"`
				} `json:"venue"`
			} `json:"home"`
		} `json:"teams"`
		Venue struct {
			Link string `json:"link"`
			Name string `json:"name"`
		} `json:"venue"`
	} `json:"gameData"`
	GamePk   int    `json:"gamePk"`
	Link     string `json:"link"`
}

func main() {
	url := os.Args[1]
	// Build the request
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

	for _, element := range record.Data.Game.Plays.Play {
		fmt.Println(element.Type, " ", element.Desc)
	}
}
