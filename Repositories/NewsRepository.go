package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetPaginatedNewsList(searchValue, sortBy, newId, authorID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.News], error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return Util.PaginatedList[BusinessObjects.News]{}, err
	}

	var news []BusinessObjects.News
	query := db.Model(&BusinessObjects.News{})

	if newId != "" {
		query = query.Where("news_id = ?", newId)
	}

	if authorID != "" {
		query = query.Where("author_id = ?", authorID)
	}

	if searchValue != "" {
		query = query.Where("title LIKE ?", "%"+searchValue+"%")
	}

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	switch sortBy {
	case "news_id_asc":
		query = query.Order("news_id ASC")
	case "news_id_desc":
		query = query.Order("news_id DESC")
	case "title_asc":
		query = query.Order("title ASC")
	case "title_desc":
		query = query.Order("title DESC")
	case "content_asc":
		query = query.Order("content ASC")
	case "content_desc":
		query = query.Order("content DESC")
	case "published_date_asc":
		query = query.Order("published_date ASC")
	case "published_date_desc":
		query = query.Order("published_date DESC")
	case "author_id_asc":
		query = query.Order("author_id ASC")
	case "author_id_desc":
		query = query.Order("author_id DESC")
	case "category_asc":
		query = query.Order("category ASC")
	case "category_desc":
		query = query.Order("category DESC")
	case "status_asc":
		query = query.Order("status ASC")
	case "status_desc":
		query = query.Order("status DESC")
	default:
		query = query.Order("created_at ASC")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.News]{}, err
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&news).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.News]{}, err
	}

	return Util.NewPaginatedList(news, total, pageIndex, pageSize), nil
}

// GetAllNews retrieves all freight news from the database
func GetAllNews() ([]BusinessObjects.News, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var news []BusinessObjects.News
	if err := db.Find(&news).Error; err != nil {
		return nil, err
	}

	return news, nil
}

// GetNewByID retrieves a freight news by its ID
func GetNewByID(newsID string) (BusinessObjects.News, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.News{}, err
	}

	var news BusinessObjects.News
	if err := db.First(&news, "news_id = ?", newsID).Error; err != nil {
		return BusinessObjects.News{}, err
	}

	return news, nil
}

// CreateNew adds a new freight news to the database
func CreateNew(news BusinessObjects.News) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Create(&news).Error; err != nil {
		return err
	}

	return nil
}

// UpdateNew updates an existing freight news
func UpdateNew(news BusinessObjects.News) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Save(&news).Error; err != nil {
		return err
	}

	return nil
}

// DeleteNew removes a freight news from the database by its ID
func DeleteNew(newsID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	// if err := db.Delete(&BusinessObjects.News{}, "news_id = ?", newsID).Error; err != nil {
	// 	return err
	// }

	if err := db.Model(&BusinessObjects.News{}).Where("news_id = ?", newsID).Update("status", false).Error; err != nil {
		return err
	}

	return nil
}
