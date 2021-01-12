package controller

import (
	"../entity"
	"../repository"
	"encoding/json"
	"math/rand"
	"net/http"
)
type Post struct {
	Id 		int		`json:"id"`
	Title 	string	`json:"title"`
	Text	string	`json:"text"`
}

var (
	repo = repository.NewPostRepository()
)


func GetPosts(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error getting the posts"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func AddPost(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}
	post.ID = rand.Int()
	repo.Save(&post)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}
