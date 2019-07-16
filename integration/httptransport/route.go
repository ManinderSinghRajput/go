package httptransport

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"

	log "github.com/luciferdocker/go/mylog"
)

var posts []Post

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func ServeProducerApi() {
	//Creating new router
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/data", postDataToKafka).Methods("POST")

	http.ListenAndServe(":7891", router)
}


func postDataToKafka(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(post)
	if post.ID == ""{
		post.ID = strconv.Itoa(rand.Intn(1000000))
	}
	posts = append(posts, post)
	log.Info(fmt.Sprintf("Before: Message send is : %v", post))
	json.NewEncoder(w).Encode(&post)
	log.Info(fmt.Sprintf("After: Message send is : %v", post))
}