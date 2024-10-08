package Services

import (
	"context"
	"testing"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"th3y3m/e-commerce-platform/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPaginatedNewsList_Success(t *testing.T) {
	newsRepository := &mocks.INewsRepository{}

	newsService := NewNewsService(newsRepository)

	ctx := context.TODO()
	searchValue := ""
	sortBy := ""
	newID := ""
	authorID := ""
	pageIndex := 1
	pageSize := 10
	status := true

	newsRepository.On("GetPaginatedNewsList", ctx, searchValue, sortBy, newID, authorID, pageIndex, pageSize, &status).Return(Util.PaginatedList[BusinessObjects.News]{}, nil)

	_, err := newsService.GetPaginatedNewsList(ctx, searchValue, sortBy, newID, authorID, pageIndex, pageSize, &status)

	assert.NoError(t, err)
	newsRepository.AssertExpectations(t)
}

func TestGetAllNews_Success(t *testing.T) {
	newsRepository := &mocks.INewsRepository{}

	newsService := NewNewsService(newsRepository)

	ctx := context.TODO()

	newsRepository.On("GetAllNews", ctx).Return([]BusinessObjects.News{}, nil)

	_, err := newsService.GetAllNews(ctx)

	assert.NoError(t, err)
	newsRepository.AssertExpectations(t)
}

func TestGetNewsByID_Success(t *testing.T) {
	newsRepository := &mocks.INewsRepository{}

	newsService := NewNewsService(newsRepository)

	ctx := context.TODO()
	newID := "1"

	newsRepository.On("GetNewByID", ctx, newID).Return(BusinessObjects.News{}, nil)

	_, err := newsService.GetNewsByID(ctx, newID)

	assert.NoError(t, err)
	newsRepository.AssertExpectations(t)
}

func TestCreateNews_Success(t *testing.T) {
	newsRepository := &mocks.INewsRepository{}

	newsService := NewNewsService(newsRepository)

	ctx := context.TODO()

	title := "title"
	content := "content"
	authorID := "1"
	category := "category"
	imageURL := "imageURL"

	newsRepository.On("CreateNew", ctx, mock.AnythingOfType("BusinessObjects.News")).Return(nil)

	err := newsService.CreateNews(ctx, title, content, authorID, category, imageURL)

	assert.NoError(t, err)
	newsRepository.AssertExpectations(t)
}
func TestCreateNews_Error(t *testing.T) {
	newsRepository := &mocks.INewsRepository{}

	newsService := NewNewsService(newsRepository)

	ctx := context.TODO()
	error := assert.AnError

	title := "title"
	content := "content"
	authorID := "1"
	category := "category"
	imageURL := "imageURL"

	newsRepository.On("CreateNew", ctx, mock.AnythingOfType("BusinessObjects.News")).Return(error)

	err := newsService.CreateNews(ctx, title, content, authorID, category, imageURL)

	assert.Error(t, err)
	assert.Equal(t, error, err)
	newsRepository.AssertExpectations(t)
}

func TestUpdateNews_Success(t *testing.T) {
	newsRepository := &mocks.INewsRepository{}

	newsService := NewNewsService(newsRepository)

	ctx := context.TODO()

	news := BusinessObjects.News{NewsID: "1", Title: "title", Content: "content", AuthorID: "1", Category: "category", ImageURL: "imageURL"}

	newsRepository.On("GetNewByID", ctx, news.NewsID).Return(BusinessObjects.News{}, nil)
	newsRepository.On("UpdateNew", ctx, mock.AnythingOfType("BusinessObjects.News")).Return(nil)

	err := newsService.UpdateNews(ctx, news.NewsID, news.Title, news.Content, news.AuthorID, news.Category, news.ImageURL)

	assert.NoError(t, err)
	newsRepository.AssertExpectations(t)
}

func TestUpdateNews_Error_GetByID(t *testing.T) {
	newsRepository := &mocks.INewsRepository{}

	newsService := NewNewsService(newsRepository)

	ctx := context.TODO()
	error := assert.AnError

	news := BusinessObjects.News{NewsID: "1", Title: "title", Content: "content", AuthorID: "1", Category: "category", ImageURL: "imageURL"}

	newsRepository.On("GetNewByID", ctx, news.NewsID).Return(BusinessObjects.News{}, error)

	err := newsService.UpdateNews(ctx, news.NewsID, news.Title, news.Content, news.AuthorID, news.Category, news.ImageURL)

	assert.Error(t, err)
	assert.Equal(t, error, err)
	newsRepository.AssertExpectations(t)
}

func TestUpdateNews_Error_Update(t *testing.T) {
	newsRepository := &mocks.INewsRepository{}

	newsService := NewNewsService(newsRepository)

	ctx := context.TODO()
	error := assert.AnError

	news := BusinessObjects.News{NewsID: "1", Title: "title", Content: "content", AuthorID: "1", Category: "category", ImageURL: "imageURL"}

	newsRepository.On("GetNewByID", ctx, news.NewsID).Return(BusinessObjects.News{}, nil)
	newsRepository.On("UpdateNew", ctx, mock.AnythingOfType("BusinessObjects.News")).Return(error)

	err := newsService.UpdateNews(ctx, news.NewsID, news.Title, news.Content, news.AuthorID, news.Category, news.ImageURL)

	assert.Error(t, err)
	assert.Equal(t, error, err)
	newsRepository.AssertExpectations(t)
}

func TestDeleteNews_Success(t *testing.T) {
	newsRepository := &mocks.INewsRepository{}

	newsService := NewNewsService(newsRepository)

	ctx := context.TODO()
	newID := "1"

	newsRepository.On("DeleteNew", ctx, newID).Return(nil)

	err := newsService.DeleteNews(ctx, newID)

	assert.NoError(t, err)
	newsRepository.AssertExpectations(t)
}
