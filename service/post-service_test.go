package service

import (
	"testing"

	"github.com/REST-API/entity"
	"github.com/REST-API/mocks"
	"github.com/stretchr/testify/assert"
)

// type MockRepository struct {
// 	mock.Mock
// }

// func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
// 	args := mock.Called()
// 	result := args.Get(0)
// 	return result.(*entity.Post), args.Error(1)
// }

// func (mock *MockRepository) FindAll() ([]entity.Post, error) {
// 	args := mock.Called()
// 	result := args.Get(0)
// 	return result.([]entity.Post), args.Error(1)
// }

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)

	assert.Equal(t, "The post is empty", err.Error())

}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{ID: 1, Title: "", Text: "B"}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "The post title is empty", err.Error())
}

func TestFindAll(t *testing.T) {
	mockRepo := &mocks.PostRepository{}

	var identifier int64 = 1

	post := entity.Post{ID: identifier, Title: "A", Text: "B"}

	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	// Mock assertion : Behavioor assertion
	mockRepo.AssertExpectations(t)
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)

	// Data Asseertiono
}

func TestCreate(t *testing.T) {
	mockRepo := &mocks.PostRepository{}

	post := entity.Post{Title: "A", Text: "B"}
	mockRepo.On("Save", &post).Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.Nil(t, err)

	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)

}
