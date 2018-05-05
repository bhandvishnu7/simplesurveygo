package main

import (
	"encoding/json"
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"io/ioutil"
	"net/http"
	"sync"
)

type Movie struct {
	Title    string `json:"title"`
	Year     int16  `json:"year"`
	Director string `json:"director"`
	Cast     string `json:"cast"`
	Genre    string `json:"genre"`
	Notes    string `json:"notes"`
}

type Input struct {
	Movie []Movie `json:"movie"`
}

var jobs = make(chan Movie, 10)
var done = make(chan bool)
var MgoSession *mgo.Session

func init() {
	if MgoSession == nil {
		var err error
		MgoSession, err = mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
	}
}

func InsertSingleMovieDetails(movie Movie) {
	session := MgoSession.Clone()
	defer session.Close()

	clctn := session.DB("simplesurveys").C("movie")
	clctn.Insert(movie)
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		InsertSingleMovieDetails(job)
	}
	wg.Done()
}

func allocate(noOfMovies int, movies []Movie) {
	var index = 0

	for i := 0; i < noOfMovies; i++ {
		jobs <- movies[index]
		index += 1
	}
	close(jobs)
	done <- true
}

func createWorkerPool() {
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
}

func makeEntries(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var movieJsonInput Input

	body, _ := ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &movieJsonInput)

	noOfJobs := len(movieJsonInput.Movie)
	fmt.Println("No of records - ", noOfJobs)

	go allocate(noOfJobs, movieJsonInput.Movie)

	createWorkerPool()
	<-done
}

func main() {
	http.HandleFunc("/", makeEntries)

	http.ListenAndServe(":3000", nil)
}
