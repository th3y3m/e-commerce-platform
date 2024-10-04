package Interface

import (
	"context"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

type INewsRepository interface {
	GetPaginatedNewsList(ctx context.Context, searchValue, sortBy, newId, authorID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.News], error)
	GetAllNews(ctx context.Context) ([]BusinessObjects.News, error)
	GetNewByID(ctx context.Context, newsID string) (BusinessObjects.News, error)
	CreateNew(ctx context.Context, news BusinessObjects.News) error
	UpdateNew(ctx context.Context, news BusinessObjects.News) error
	DeleteNew(ctx context.Context, newsID string) error
}

type INewsService interface {
	GetPaginatedNewsList(ctx context.Context, searchValue, sortBy, newId, authorID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.News], error)
	GetAllNews(ctx context.Context) ([]BusinessObjects.News, error)
	GetNewsByID(ctx context.Context, newsID string) (BusinessObjects.News, error)
	CreateNews(ctx context.Context, title, content, authorID, category, ImageURL string) error
	UpdateNews(ctx context.Context, newsId, title, content, authorID, category, ImageURL string) error
	DeleteNews(ctx context.Context, newsID string) error
}
