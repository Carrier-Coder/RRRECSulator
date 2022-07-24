package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"errors"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"rrecsulator.com/dataSets"
	"rrecsulator.com/drive"
	"rrecsulator.com/fullRRECS"
)

// globals used for tracking hits
var driveCount uint64
var fixedCount uint64
var dailyCount uint64

// hours b/n saving the hit counters
const SAVE_HOURS = 24

// tiles to save and load hit counters
var HIT_FILE = "./hitfile"

// handle post requests with drive time data
func driveDataPost(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var dd drive.DriveData

	err := decoder.Decode(&dd)
	if err != nil {
		log.Printf("Error decoding drive data\n%v", err)
		http.Error(w, "Bad Request", 400)
		return
	}

	temp := 0.0
	for _, v := range dd.Distances {
		temp += v
	}
	temp = temp / 5280

	//increment the hit counter
	atomic.AddUint64(&driveCount, 1)

	result := fmt.Sprintf(`{"time": %4.2f}`, dd.GetTime())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

// handle post requests with fixed data
func fixedDataPost(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var fd dataSets.FixedData

	err := decoder.Decode(&fd)
	if err != nil {
		log.Printf("Error decoding fixed data\n%   v", err)
		http.Error(w, "Bad Request", 400)
		return
	}
	log.Println("Processed fixed data post")

	//increment the hit counter
	atomic.AddUint64(&fixedCount, 1)

	w.WriteHeader(http.StatusOK)
}

//handle requests to see the hit counters
func hitCounter(w http.ResponseWriter, r *http.Request) {
	log.Println("accessed the hit counter")
	fmt.Fprintf(w, "drive: %d\nfixed: %d\ndaily: %d\n",
		driveCount, fixedCount, dailyCount)
}

// handle post requests with daily data
func dailyDataPost(w http.ResponseWriter, r *http.Request) {

	type fullData struct {
		dataSets.DailyData
		dataSets.FixedData
	}

	decoder := json.NewDecoder(r.Body)
	var data fullData

	err := decoder.Decode(&data)
	if err != nil {
		log.Printf("Error decoding daily data\n%   v", err)
		http.Error(w, "Bad Request", 400)
		return
	}

	log.Println("Processed a daily data post")
	report := fullRRECS.GenerateReport(data.FixedData, data.DailyData)
	result := "{\"result\" :" + "\"" + report + "\"" + "}"

	//json requires control characters to be escaped
	// maybe there is a less hacky way to do this?
	test := strings.Replace(result, "\n", "\\n", -1)

	//increment the hit counter
	atomic.AddUint64(&dailyCount, 1)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(test))
}

func main() {

	//load hit counters
	counters := loadHitCounters()
	//these don't have to be atomic, I don't think
	//should only be happening in the main thread no threat of race
	atomic.AddUint64(&driveCount, counters[0])
	atomic.AddUint64(&fixedCount, counters[1])
	atomic.AddUint64(&dailyCount, counters[2])

	//start routine for saving hit counters
	go func() {
		for {
			time.Sleep(SAVE_HOURS * time.Hour)
			saveHitCounters()
		}
	}()

	r := mux.NewRouter()

	//Handle api requests
	api := r.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/drive", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get"}`))
	}).Methods("GET")
	api.HandleFunc("/drive", driveDataPost).Methods("POST")
	api.HandleFunc("/daily", dailyDataPost).Methods("POST")
	api.HandleFunc("/fixed", fixedDataPost).Methods("POST")
	api.HandleFunc("/secret_hits", hitCounter).Methods("GET")

	log.Println("http://localhost:8888")

	//handle Cors
	//TODO: should be env or flags
	corsOrigin := handlers.AllowedOrigins([]string{
		"http://localhost:8001",
		"https://rrecsulator.com",
	})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"})
	corsMethods := handlers.AllowedMethods([]string{"POST", "GET"})

	log.Fatal(http.ListenAndServe(":8888", handlers.CORS(
		corsOrigin,
		corsHeaders,
		corsMethods,
	)(r)))

}

func saveHitCounters() {
	//these should probably be atomic, but impact will be minimal
	s := fmt.Sprintf("%d\n%d\n%d", driveCount, fixedCount, dailyCount)
	err := ioutil.WriteFile(HIT_FILE, []byte(s), 0777)
	if err != nil {
		log.Println(err)
	}
}

func loadHitCounters() []uint64 {

	//read the hit counts from disk, if they exist
	_, err := os.Stat(HIT_FILE)
	if err != nil {
		// if file doesn't exist, create and seed with zeros
		if errors.Is(err, os.ErrNotExist) {
			f, err := os.Create(HIT_FILE)
			if err != nil {
				log.Println("Could not create hit file")
				log.Println(err)
			}
			defer f.Close()
			f.Write([]byte("0\n0\n0"))
		} else {
			// some other disk error
			log.Println("Error accessing hit file")
			log.Println(err)
		}
	}

	//read in our values
	content, err := ioutil.ReadFile(HIT_FILE)
	if err != nil {
		log.Println("Could not read hit file")
		return []uint64{0, 0, 0}
	}
	data := strings.Split(strings.Trim(string(content), " \n\r\t"), "\n")
	out := []uint64{}
	for _, v := range data {
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			//if can't parse value, just zero out, now worth crashing over
			log.Println("Error parsing hit file")
			log.Println(err)
			out = append(out, 0)
			continue
		}
		out = append(out, i)
	}

	return out
}
