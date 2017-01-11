package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Home Page")
}

func cur_time(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	io.WriteString(w, now.String())
}

func news(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://feeds.reuters.com/reuters/MostRead")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	news, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(news))
}


func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/time", cur_time)
	http.HandleFunc("/news", news)
	http.ListenAndServe(":8080", nil)
}
