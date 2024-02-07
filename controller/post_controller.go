package controller

import (
	"encoding/json"

	"net/http"

	"github.com/REST-API/entity"
	"github.com/REST-API/errors"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type controller struct {
	postService PostService
}

func NewPostController(service PostService) *controller {

	return &controller{
		postService: service,
	}
}

func (c *controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	posts, err := c.postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the posts"})
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func (c *controller) AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the posts"})
		return
	}
	err = c.postService.Validate(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})

	}
	result, err := c.postService.Create(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error saving the post"})
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
	return

}
