package Interface

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

type INewsRepository interface {
	GetPaginatedNewsList(searchValue, sortBy, newId, authorID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.News], error)
	GetAllNews() ([]BusinessObjects.News, error)
	GetNewByID(newsID string) (BusinessObjects.News, error)
	CreateNew(news BusinessObjects.News) error
	UpdateNew(news BusinessObjects.News) error
	DeleteNew(newsID string) error
}

type INewsService interface {
	GetPaginatedNewsList(searchValue, sortBy, newId, authorID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.News], error)
	GetAllNews() ([]BusinessObjects.News, error)
	GetNewsByID(newsID string) (BusinessObjects.News, error)
	CreateNews(title, content, authorID, category, ImageURL string) error
	UpdateNews(newsId, title, content, authorID, category, ImageURL string) error
	DeleteNews(newsID string) error
}
