package httptransport

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
	"math/rand"
	"myGitCode/codeDataBroker/kafka"
	"net/http"
	"strconv"

	log "myGitCode/mylog"
)

var posts []Post

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var producer sarama.SyncProducer
var producerTopic string

func ServeProducerApi(listenAddrApi string, kafkaProducer sarama.SyncProducer, topic string) error{
	//Creating new router
	router := mux.NewRouter()

	posts = append(posts, Post{ID: "1", Title: "My first post", Body: "This is the content of my first post"})

	router.HandleFunc("/api/v1/data", postDataToKafka).Methods("POST")

	/*router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")*/

	producer = kafkaProducer
	producerTopic = topic
	log.Info("Starting serving Producer Api")
	err := http.ListenAndServe(listenAddrApi, router)
	if err != nil{
		return err
	}
	return nil
}

//Request Handlers

func postDataToKafka(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	if post.ID == ""{
		post.ID = strconv.Itoa(rand.Intn(1000000))
	}
	posts = append(posts, post)
	log.Debug(fmt.Sprintf("Before: Message send is : %#v", post))
	_ = json.NewEncoder(w).Encode(&post)

	message, err := json.Marshal(post)
	if err != nil{
		log.Error("Failed to Convert data to string. " + err.Error())
	}

	log.Info(string(message))

	err = kafka.Push(string(message), producer, producerTopic)
	if err != nil{
		log.Error("Failed to push data to kafka. " + err.Error())
	}
}

/*func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000000))
	posts = append(posts, post)
	json.NewEncoder(w).Encode(&post)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Post{})
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			var post Post
			_ = json.NewDecoder(r.Body).Decode(&post)
			post.ID = params["id"]
			posts = append(posts, post)
			json.NewEncoder(w).Encode(&post)
			return
		}
	}
	json.NewEncoder(w).Encode(posts)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
}*/
