package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

type NewsService struct {
	newsRepository Interface.INewsRepository
}

func NewNewsService(newsRepository Interface.INewsRepository) Interface.INewsService {
	return &NewsService{newsRepository: newsRepository}
}

func (n *NewsService) GetPaginatedNewsList(searchValue, sortBy, newID, authorID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.News], error) {
	return n.newsRepository.GetPaginatedNewsList(searchValue, sortBy, newID, authorID, pageIndex, pageSize, status)
}

func (n *NewsService) GetAllNews() ([]BusinessObjects.News, error) {
	return n.newsRepository.GetAllNews()
}

func (n *NewsService) GetNewsByID(id string) (BusinessObjects.News, error) {
	return n.newsRepository.GetNewByID(id)
}

func (n *NewsService) CreateNews(title, content, authorID, category, ImageURL string) error {
	news := BusinessObjects.News{
		NewsID:        "NEWS" + Util.GenerateID(10),
		Title:         title,
		Content:       content,
		AuthorID:      authorID,
		Category:      category,
		PublishedDate: time.Now(),
		ImageURL:      ImageURL,
		Status:        true,
	}

	err := n.newsRepository.CreateNew(news)
	if err != nil {
		return err
	}

	return nil
}

func (n *NewsService) UpdateNews(newsId, title, content, authorID, category, ImageURL string) error {
	news, err := n.newsRepository.GetNewByID(newsId)
	if err != nil {
		return err
	}

	news.Title = title
	news.Content = content
	news.AuthorID = authorID
	news.Category = category
	news.ImageURL = ImageURL

	return n.newsRepository.UpdateNew(news)
}

func (n *NewsService) DeleteNews(id string) error {
	return n.newsRepository.DeleteNew(id)
}
