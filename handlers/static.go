package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Page struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type Candidate struct {
	ID         int    `json:"id"`
	FirstName  string `json:"fname"`
	MiddleName string `json:"mname"`
	LastName   string `json:"lname"`
	Suffix     string `json:"suffix"`
	UseSuffix  bool   `json:"usesuffix"`
	Party      string `json:"party"`
	Winner     bool   `json:"winner"`
	Vpct       int    `json:"vpct"`
	PctDecimal string `json:"pctDecimal"`
	Inc        bool   `json:"inc"`
	Votes      int    `json:"votes"`
	Cvotes     string `json:"cvotes"`
	Evotes     int    `json:"evotes"`
}

type Race struct {
	Bgs          bool        `json:"bgs"`
	Status       string      `json:"status"`
	RaceID       string      `json:"raceid"`
	State        string      `json:"state"`
	ElectionDate string      `json:"electiondate"`
	Ncon         bool        `json:"ncon"`
	KeyRace      bool        `json:"keyrace"`
	Rtype        string      `json:"rtype"`
	Race         string      `json:"race"`
	Title2       string      `json:"title2"`
	Evotes       int         `json:"evotes"`
	Ahead        int         `json:"ahead"`
	PollClose    int64       `json:"pollclose"`
	Xpoll        bool        `json:"xpoll"`
	Cresults     bool        `json:"cresults"`
	Cmap         bool        `json:"cmap"`
	CallTime     int64       `json:"calltime"`
	Ts           int64       `json:"ts"`
	PctsRep      int         `json:"pctsrep"`
	Sw           bool        `json:"sw"`
	Ip           string      `json:"ip"`
	Candidates   []Candidate `json:"candidates"`
}

type ElectionResult struct {
	Pages      []Page      `json:"pages"`
	Race       string      `json:"race"`
	Races      []Race      `json:"races"`
	Candidates []Candidate `json:"candidates"`
	Lts        int64       `json:"lts"`
}

func ElectionResultHandler(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("results.json")
	if err != nil {
		log.Fatalf("Failed to open JSON file: %v", err)
	}
	defer jsonFile.Close()

	// Read the JSON file into a byte slice
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	// Unmarshal the JSON data into the structs
	var results ElectionResult
	err = json.Unmarshal(byteValue, &results)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
