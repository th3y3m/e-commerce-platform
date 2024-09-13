package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

func GetPaginatedNewsList(searchValue, sortBy, newID, authorID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.News], error) {
	return Repositories.GetPaginatedNewsList(searchValue, sortBy, newID, authorID, pageIndex, pageSize, status)
}

func GetAllNews() ([]BusinessObjects.News, error) {
	return Repositories.GetAllNews()
}

func GetNewsByID(id string) (BusinessObjects.News, error) {
	return Repositories.GetNewByID(id)
}

func CreateNews(title, content, authorID, category, ImageURL string) error {
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

	err := Repositories.CreateNew(news)
	if err != nil {
		return err
	}

	return nil
}

func UpdateNews(newsId, title, content, authorID, category, ImageURL string) error {
	news, err := Repositories.GetNewByID(newsId)
	if err != nil {
		return err
	}

	news.Title = title
	news.Content = content
	news.AuthorID = authorID
	news.Category = category
	news.ImageURL = ImageURL

	return Repositories.UpdateNew(news)
}

func DeleteNews(id string) error {
	return Repositories.DeleteNew(id)
}
