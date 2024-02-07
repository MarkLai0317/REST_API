package main

import (
	"fmt"

	"net/http"

	"github.com/REST-API/repository"
	"github.com/REST-API/router"

	"github.com/REST-API/controller"
	"github.com/REST-API/service"
)

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

var (
	httpRouter router.Router = router.NewChiRouter()

	postRepository service.PostRepository = repository.NewFirestoreRepository()
	postService    controller.PostService = service.NewPostService(postRepository)
	postController PostController         = controller.NewPostController(postService)
)

func main() {

	const port string = ":8080"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Up and running")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
