package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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

const SOUNDS_PATH_PREFIX = ""
const SOUNDS_NAMES_PATH = "sounds.txt"

func handleFuncUploadSound(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var msg Message
		json.NewDecoder(r.Body).Decode(&msg)
		data, err := base64.StdEncoding.DecodeString(msg.Text)
		logPotentialErr(err)
		err = os.WriteFile(SOUNDS_PATH_PREFIX+msg.Name, data, 0644)
		logPotentialErr(err)

		f, err := os.OpenFile(SOUNDS_NAMES_PATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write([]byte(strings.Split(msg.Name, ".")[0] + " " + msg.Name + "\n")); err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
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
