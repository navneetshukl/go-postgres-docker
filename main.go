package main

import (
	"log"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	ConnectToDB()

	w.Write([]byte("Inserted to DB"))
}

func main() {
	http.HandleFunc("/", Test)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error in opening connection ", err)
	}
}
