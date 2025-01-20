package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Name string
	Text string
}

func logPotentialErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleFuncHome(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		dat, err := os.ReadFile("site.html")
		logPotentialErr(err)
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, string(dat))
	}
}

func handleFuncUploadSound(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var msg Message
		json.NewDecoder(r.Body).Decode(&msg)
		data, err := base64.StdEncoding.DecodeString(msg.Text)
		logPotentialErr(err)
		err = os.WriteFile(msg.Name, data, 0644)
		logPotentialErr(err)
	}
}

func main() {
	http.HandleFunc("/", handleFuncHome)
	http.HandleFunc("/upload_sound", handleFuncUploadSound)
	fmt.Println("OK")
	go http.ListenAndServe(":3000", nil)

	var s string
	for {
		fmt.Scan(&s)
		if s == "/clear" {
			break
		}
	}
}
