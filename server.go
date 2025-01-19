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
	Text string
}

// const MXM_CHAT_SIZE int = 20000

// var chat []Message
// var msg_id int = 0

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// func updateChatFile() {
// 	chat_content, err := json.Marshal(chat)
// 	logErr(err)
// 	err = os.WriteFile("chat.json", chat_content, 0666)
// 	logErr(err)
// }

// func handleFuncChat(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		json.NewEncoder(w).Encode(chat)
// 	case http.MethodPost:
// 		var msg Message
// 		json.NewDecoder(r.Body).Decode(&msg)
// 		chat = append(chat, msg)
// 		if len(chat) > MXM_CHAT_SIZE {
// 			chat = chat[:MXM_CHAT_SIZE]
// 		}

// 		if msg_id%5 == 0 {
// 			updateChatFile()
// 		}
// 		msg_id++
// 	}
// }

func handleFuncHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
	switch r.Method {
	case http.MethodGet:
		dat, err := os.ReadFile("site.html")
		logErr(err)
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, string(dat))
	}
}

func handleFuncUploadSound(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received")
	switch r.Method {
	case http.MethodPost:
		var msg Message
		json.NewDecoder(r.Body).Decode(&msg)
		fmt.Print(len(msg.Text))
		fmt.Print(msg.Text)
		fmt.Println("a")
		data, err := base64.StdEncoding.DecodeString(msg.Text)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile("img.png", data, 0644)
		if err != nil {
			log.Fatal(err)
		}
		// mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
		// reader := multipart.NewReader(r.Body, params["boundary"])
	}
}

func main() {
	// chat_content, _ := os.ReadFile("chat.json")
	// err := json.Unmarshal(chat_content, &chat)
	// logErr(err)

	http.HandleFunc("/", handleFuncHome)
	http.HandleFunc("/upload_sound", handleFuncUploadSound)
	fmt.Println("OK")
	go http.ListenAndServe(":3000", nil)

	var s string
	for {
		fmt.Scan(&s)
		if s == "/clear" {
			// chat = chat[:0]
		}
	}
}
