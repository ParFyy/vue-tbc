package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Fixture struct {
	Location      string
	Date          string
	Time          string
	Teamfor       string
	Teamagainst   string
	Pointsfor     string `json:",omitempty"`
	Pointsagainst string `json:",omitempty"`
}

var fixtures []Fixture

func main() {
	fileServer := http.FileServer(http.Dir("dist"))
	http.Handle("/", fileServer)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	UpdateTeamFixtures()

	http.HandleFunc("/GetTbcFixtures", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "https://test.tasburghbadmintonclub.co.uk")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(fixtures)
	})

	http.HandleFunc("/GetTbcFixtures/Refresh", func(w http.ResponseWriter, r *http.Request) {
		UpdateTeamFixtures()
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":"+port, nil)
}

func UpdateTeamFixtures() {
	var updatedFixtures []Fixture
	updatedFixtures = append(updatedFixtures, GetFixturesFromSheets("M1", "https://docs.google.com/spreadsheets/d/1KUhnkSxCt9gPtGjn1lx8I2kY0Pj23070z6WeyPSWySc/export?format=csv&id=1KUhnkSxCt9gPtGjn1lx8I2kY0Pj23070z6WeyPSWySc&gid=434795360")...)
	updatedFixtures = append(updatedFixtures, GetFixturesFromSheets("M2", "https://docs.google.com/spreadsheets/d/1KUhnkSxCt9gPtGjn1lx8I2kY0Pj23070z6WeyPSWySc/export?format=csv&id=1KUhnkSxCt9gPtGjn1lx8I2kY0Pj23070z6WeyPSWySc&gid=1194335473")...)
	updatedFixtures = append(updatedFixtures, GetFixturesFromSheets("M3", "https://docs.google.com/spreadsheets/d/1KUhnkSxCt9gPtGjn1lx8I2kY0Pj23070z6WeyPSWySc/export?format=csv&id=1KUhnkSxCt9gPtGjn1lx8I2kY0Pj23070z6WeyPSWySc&gid=918711655")...)
	updatedFixtures = append(updatedFixtures, GetFixturesFromSheets("X1", "https://docs.google.com/spreadsheets/d/1KUhnkSxCt9gPtGjn1lx8I2kY0Pj23070z6WeyPSWySc/export?format=csv&id=1KUhnkSxCt9gPtGjn1lx8I2kY0Pj23070z6WeyPSWySc&gid=1247946663")...)
	updatedFixtures = append(updatedFixtures, GetFixturesFromSheets("X2", "https://docs.google.com/spreadsheets/d/1KUhnkSxCt9gPtGjn1lx8I2kY0Pj23070z6WeyPSWySc/export?format=csv&id=1KUhnkSxCt9gPtGjn1lx8I2kY0Pj23070z6WeyPSWySc&gid=737070727")...)
	updatedFixtures = append(updatedFixtures, GetFixturesFromSheets("L1", "https://docs.google.com/spreadsheets/d/1KUhnkSxCt9gPtGjn1lx8I2kY0Pj23070z6WeyPSWySc/export?format=csv&id=1KUhnkSxCt9gPtGjn1lx8I2kY0Pj23070z6WeyPSWySc&gid=669783534")...)
	fixtures = updatedFixtures
}

func FixtureLengthCheck(field int, val []string) string {
	if len(val) > field {
		return val[field]
	}

	return ""
}

func GetFixturesFromSheets(teamName string, url string) []Fixture {
	var fixtures []Fixture

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	csvReader := csv.NewReader(resp.Body)
	csvReader.FieldsPerRecord = -1
	csvReader.Read()

	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, fixture := range csvData {
		fixtures = append(fixtures,
			Fixture{
				Teamfor:       teamName,
				Teamagainst:   fixture[0],
				Date:          fixture[1],
				Location:      fixture[2],
				Time:          fixture[3],
				Pointsfor:     FixtureLengthCheck(4, fixture),
				Pointsagainst: FixtureLengthCheck(5, fixture),
			})
	}

	return fixtures
}
