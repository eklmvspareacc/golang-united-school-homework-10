package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", nameHandler).Methods("GET")
	router.HandleFunc("/bad", badHandler).Methods("GET")
	router.HandleFunc("/data", dataHandler).Methods("POST")
	router.HandleFunc("/headers", headersHandler).HeadersRegexp("a", `^[+-]?\d+$`, "b", `^[+-]?\d+$`).Methods("POST")
	router.PathPrefix("/").HandlerFunc(defaultHandler)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!", vars["PARAM"])
}

func badHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I got message:\n%s", body)
}

func headersHandler(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("a"))
	// should be no error because of regex?
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("a+b", strconv.Itoa(a+b))
	w.WriteHeader(http.StatusOK)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
