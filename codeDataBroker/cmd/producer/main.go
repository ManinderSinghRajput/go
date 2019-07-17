package main

import (
	log "myGitCode/mylog"
	"myGitCode/codeDataBroker/httptransport"
)

func main() {

	log.Debug("Calling to start Producer API Server")

	startAndServeProducerApi()

	log.Debug("Serving Producer API is Done")
}

func startAndServeProducerApi() {

	var posts []httptransport.Post
	posts = append(posts, httptransport.Post{ID: "1", Title: "My first post", Body: "This is the content of my first post"})
	posts = append(posts, httptransport.Post{ID: "2", Title: "My second post", Body: "This is the content of my second post"})

	httptransport.ServeProducerApi()
	log.Info("")
}
