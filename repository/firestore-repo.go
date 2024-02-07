package repository

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/REST-API/entity"
	"github.com/joho/godotenv"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

type repo struct{}

// NewFireStoreRepository creates a new repo
func NewFirestoreRepository() *repo {
	return &repo{}
}

const (
	projectId      string = "true-elevator-413315"
	collectionName string = "posts"
	databaseID     string = "go-test"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()

	dotenv := goDotEnvVariable("GOOGLE_APPLICATION_CREDENTIALS")

	client, err := firestore.NewClientWithDatabase(ctx, projectId, databaseID, option.WithCredentialsFile(dotenv))
	// ctx := context.Background()
	// client, err := firestore.NewClientWithDatabase(ctx, projectId, databaseID)
	if err != nil {
		log.Fatalf("fail to create firebase client: %v", err)
		return nil, err
	}
	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("fail to add new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {

	ctx := context.Background()

	dotenv := goDotEnvVariable("GOOGLE_APPLICATION_CREDENTIALS")

	client, err := firestore.NewClientWithDatabase(ctx, projectId, databaseID, option.WithCredentialsFile(dotenv))
	// ctx := context.Background()
	// client, err := firestore.NewClientWithDatabase(ctx, projectId, databaseID)
	if err != nil {
		log.Fatalf("fail to create firebase client: %v", err)
		return nil, err
	}
	defer client.Close()

	posts := []entity.Post{}
	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("fail to iterate the list of posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}

	return posts, nil

}
