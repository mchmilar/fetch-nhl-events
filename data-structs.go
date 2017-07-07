package main

import "time"

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

type TeamsPlayer struct {
	Person struct {
		ID int `json:"id"`
		FullName string `json:"fullName"`
		Link string `json:"link"`
		ShootsCatches string `json:"shootsCatches"`
		RosterStatus string `json:"rosterStatus"`
	} `json:"person"`
	JerseyNumber string `json:"jerseyNumber"`
	Position struct {
		Code string `json:"code"`
		Name string `json:"name"`
		Type string `json:"type"`
		Abbreviation string `json:"abbreviation"`
	} `json:"position"`
	Stats struct {
		SkaterStats struct {
			TimeOnIce string `json:"timeOnIce"`
			Assists int `json:"assists"`
			Goals int `json:"goals"`
			Shots int `json:"shots"`
			Hits int `json:"hits"`
			PowerPlayGoals int `json:"powerPlayGoals"`
			PowerPlayAssists int `json:"powerPlayAssists"`
			PenaltyMinutes int `json:"penaltyMinutes"`
			FaceOffPct float64 `json:"faceOffPct"`
			FaceOffWins int `json:"faceOffWins"`
			FaceoffTaken int `json:"faceoffTaken"`
			Takeaways int `json:"takeaways"`
			Giveaways int `json:"giveaways"`
			ShortHandedGoals int `json:"shortHandedGoals"`
			ShortHandedAssists int `json:"shortHandedAssists"`
			Blocked int `json:"blocked"`
			PlusMinus int `json:"plusMinus"`
			EvenTimeOnIce string `json:"evenTimeOnIce"`
			PowerPlayTimeOnIce string `json:"powerPlayTimeOnIce"`
			ShortHandedTimeOnIce string `json:"shortHandedTimeOnIce"`
		} `json:"skaterStats"`
	} `json:"stats"`
}


type NHL struct {
	Copyright string `json:"copyright"`
	MetaData struct {
		Wait int `json:"wait"`
		TimeStamp string `json:"timeStamp"`
	} `json:"metaData"`
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
	LiveData struct {
		Plays struct {
			AllPlays []struct {
				Result struct {
					Event string `json:"event"`
					EventCode string `json:"eventCode"`
					EventTypeID string `json:"eventTypeId"`
					Description string `json:"description"`
				} `json:"result"`
				About struct {
					EventIdx int `json:"eventIdx"`
					EventID int `json:"eventId"`
					Period int `json:"period"`
					PeriodType string `json:"periodType"`
					OrdinalNum string `json:"ordinalNum"`
					PeriodTime string `json:"periodTime"`
					PeriodTimeRemaining string `json:"periodTimeRemaining"`
					DateTime time.Time `json:"dateTime"`
					Goals struct {
						Away int `json:"away"`
						Home int `json:"home"`
					} `json:"goals"`
				} `json:"about"`
				Coordinates struct {
				} `json:"coordinates"`
				Players []struct {
					Player struct {
						ID int `json:"id"`
						FullName string `json:"fullName"`
						Link string `json:"link"`
					} `json:"player"`
					PlayerType string `json:"playerType"`
				} `json:"players,omitempty"`
				Team struct {
					ID int `json:"id"`
					Name string `json:"name"`
					Link string `json:"link"`
					TriCode string `json:"triCode"`
				} `json:"team,omitempty"`
			} `json:"allPlays"`
			ScoringPlays []int `json:"scoringPlays"`
			PenaltyPlays []int `json:"penaltyPlays"`
			PlaysByPeriod []struct {
				StartIndex int `json:"startIndex"`
				Plays []int `json:"plays"`
				EndIndex int `json:"endIndex"`
			} `json:"playsByPeriod"`
			CurrentPlay struct {
				Result struct {
					Event string `json:"event"`
					EventCode string `json:"eventCode"`
					EventTypeID string `json:"eventTypeId"`
					Description string `json:"description"`
				} `json:"result"`
				About struct {
					EventIdx int `json:"eventIdx"`
					EventID int `json:"eventId"`
					Period int `json:"period"`
					PeriodType string `json:"periodType"`
					OrdinalNum string `json:"ordinalNum"`
					PeriodTime string `json:"periodTime"`
					PeriodTimeRemaining string `json:"periodTimeRemaining"`
					DateTime time.Time `json:"dateTime"`
					Goals struct {
						Away int `json:"away"`
						Home int `json:"home"`
					} `json:"goals"`
				} `json:"about"`
				Coordinates struct {
				} `json:"coordinates"`
			} `json:"currentPlay"`
		} `json:"plays"`
		Linescore struct {
			CurrentPeriod int `json:"currentPeriod"`
			CurrentPeriodOrdinal string `json:"currentPeriodOrdinal"`
			CurrentPeriodTimeRemaining string `json:"currentPeriodTimeRemaining"`
			Periods []struct {
				PeriodType string `json:"periodType"`
				StartTime time.Time `json:"startTime"`
				EndTime time.Time `json:"endTime"`
				Num int `json:"num"`
				OrdinalNum string `json:"ordinalNum"`
				Home struct {
					Goals int `json:"goals"`
					ShotsOnGoal int `json:"shotsOnGoal"`
					RinkSide string `json:"rinkSide"`
				} `json:"home"`
				Away struct {
					Goals int `json:"goals"`
					ShotsOnGoal int `json:"shotsOnGoal"`
					RinkSide string `json:"rinkSide"`
				} `json:"away"`
			} `json:"periods"`
			ShootoutInfo struct {
				Away struct {
					Scores int `json:"scores"`
					Attempts int `json:"attempts"`
				} `json:"away"`
				Home struct {
					Scores int `json:"scores"`
					Attempts int `json:"attempts"`
				} `json:"home"`
			} `json:"shootoutInfo"`
			Teams struct {
				Home struct {
					Team struct {
						ID int `json:"id"`
						Name string `json:"name"`
						Link string `json:"link"`
						Abbreviation string `json:"abbreviation"`
						TriCode string `json:"triCode"`
					} `json:"team"`
					Goals int `json:"goals"`
					ShotsOnGoal int `json:"shotsOnGoal"`
					GoaliePulled bool `json:"goaliePulled"`
					NumSkaters int `json:"numSkaters"`
					PowerPlay bool `json:"powerPlay"`
				} `json:"home"`
				Away struct {
					Team struct {
						ID int `json:"id"`
						Name string `json:"name"`
						Link string `json:"link"`
						Abbreviation string `json:"abbreviation"`
						TriCode string `json:"triCode"`
					} `json:"team"`
					Goals int `json:"goals"`
					ShotsOnGoal int `json:"shotsOnGoal"`
					GoaliePulled bool `json:"goaliePulled"`
					NumSkaters int `json:"numSkaters"`
					PowerPlay bool `json:"powerPlay"`
				} `json:"away"`
			} `json:"teams"`
			PowerPlayStrength string `json:"powerPlayStrength"`
			HasShootout bool `json:"hasShootout"`
		} `json:"linescore"`
		Boxscore struct {
			Teams struct {
				Away struct {
					Team struct {
						ID int `json:"id"`
						Name string `json:"name"`
						Link string `json:"link"`
						Abbreviation string `json:"abbreviation"`
						TriCode string `json:"triCode"`
					} `json:"team"`
					TeamStats struct {
						TeamSkaterStats struct {
							Goals int `json:"goals"`
							Pim int `json:"pim"`
							Shots int `json:"shots"`
							PowerPlayPercentage string `json:"powerPlayPercentage"`
							PowerPlayGoals float64 `json:"powerPlayGoals"`
							PowerPlayOpportunities float64 `json:"powerPlayOpportunities"`
							FaceOffWinPercentage string `json:"faceOffWinPercentage"`
							Blocked int `json:"blocked"`
							Takeaways int `json:"takeaways"`
							Giveaways int `json:"giveaways"`
							Hits int `json:"hits"`
						} `json:"teamSkaterStats"`
					} `json:"teamStats"`
					Players map[string]TeamsPlayer
					Goalies []int `json:"goalies"`
					Skaters []int `json:"skaters"`
					OnIce []int `json:"onIce"`
					OnIcePlus []struct {
						PlayerID int `json:"playerId"`
						ShiftDuration int `json:"shiftDuration"`
						Stamina int `json:"stamina"`
					} `json:"onIcePlus"`
					Scratches []int `json:"scratches"`
					PenaltyBox []interface{} `json:"penaltyBox"`
					Coaches []struct {
						Person struct {
							FullName string `json:"fullName"`
							Link string `json:"link"`
						} `json:"person"`
						Position struct {
							Code string `json:"code"`
							Name string `json:"name"`
							Type string `json:"type"`
							Abbreviation string `json:"abbreviation"`
						} `json:"position"`
					} `json:"coaches"`
				} `json:"away"`
				Home struct {
					Team struct {
						ID int `json:"id"`
						Name string `json:"name"`
						Link string `json:"link"`
						Abbreviation string `json:"abbreviation"`
						TriCode string `json:"triCode"`
					} `json:"team"`
					TeamStats struct {
						TeamSkaterStats struct {
							Goals int `json:"goals"`
							Pim int `json:"pim"`
							Shots int `json:"shots"`
							PowerPlayPercentage string `json:"powerPlayPercentage"`
							PowerPlayGoals float64 `json:"powerPlayGoals"`
							PowerPlayOpportunities float64 `json:"powerPlayOpportunities"`
							FaceOffWinPercentage string `json:"faceOffWinPercentage"`
							Blocked int `json:"blocked"`
							Takeaways int `json:"takeaways"`
							Giveaways int `json:"giveaways"`
							Hits int `json:"hits"`
						} `json:"teamSkaterStats"`
					} `json:"teamStats"`
					Players map[string]TeamsPlayer
					Goalies []int `json:"goalies"`
					Skaters []int `json:"skaters"`
					OnIce []int `json:"onIce"`
					OnIcePlus []struct {
						PlayerID int `json:"playerId"`
						ShiftDuration int `json:"shiftDuration"`
						Stamina int `json:"stamina"`
					} `json:"onIcePlus"`
					Scratches []int `json:"scratches"`
					PenaltyBox []interface{} `json:"penaltyBox"`
					Coaches []struct {
						Person struct {
							FullName string `json:"fullName"`
							Link string `json:"link"`
						} `json:"person"`
						Position struct {
							Code string `json:"code"`
							Name string `json:"name"`
							Type string `json:"type"`
							Abbreviation string `json:"abbreviation"`
						} `json:"position"`
					} `json:"coaches"`
				} `json:"home"`
			} `json:"teams"`
		} `json:"boxscore"`
	} `json:"liveData"`
}