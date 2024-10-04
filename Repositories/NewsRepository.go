package Repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Provider"
	"th3y3m/e-commerce-platform/Util"

	"github.com/sirupsen/logrus"
)

type NewsRepository struct {
	log *logrus.Logger
	db  Provider.IDb
}

func NewNewsRepository(log *logrus.Logger, db Provider.IDb) Interface.INewsRepository {
	return &NewsRepository{log: log, db: db}
}

func getPaginatedNewListCacheKey(searchValue, sortBy, newId, authorID string, pageIndex, pageSize int, status *bool) string {
	return fmt.Sprintf("paginatedNewsList:%s:%s:%s:%s:%d:%d:%v", searchValue, sortBy, newId, authorID, pageIndex, pageSize, status)
}

func getNewsCacheKey(newsID string) string {
	return "news:" + newsID
}

func (n *NewsRepository) GetPaginatedNewsList(ctx context.Context, searchValue, sortBy, newId, authorID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.News], error) {
	n.log.Infof("Fetching paginated news list with searchValue: %s, sortBy: %s, newId: %s, authorID: %s, pageIndex: %d, pageSize: %d, status: %v", searchValue, sortBy, newId, authorID, pageIndex, pageSize, status)
	redisClient, err := n.db.GetRedis()
	if err != nil {
		n.log.Error("Failed to get Redis client:", err)
		return Util.PaginatedList[BusinessObjects.News]{}, err
	}

	cacheKey := getPaginatedNewListCacheKey(searchValue, sortBy, newId, authorID, pageIndex, pageSize, status)
	val, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		n.log.Infof("Cache hit for paginated news list")
		var paginatedList Util.PaginatedList[BusinessObjects.News]
		if err := json.Unmarshal([]byte(val), &paginatedList); err == nil {
			return paginatedList, nil
		}
		n.log.Warn("Failed to unmarshal cached paginated list data:", err)
	}

	n.log.Infof("Cache miss for paginated news list")
	db, err := n.db.GetDB()
	if err != nil {
		n.log.Error("Failed to connect to PostgreSQL:", err)
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
		n.log.Error("Failed to count news:", err)
		return Util.PaginatedList[BusinessObjects.News]{}, err
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&news).Error; err != nil {
		n.log.Error("Failed to fetch paginated news:", err)
		return Util.PaginatedList[BusinessObjects.News]{}, err
	}

	paginatedList := Util.NewPaginatedList(news, total, pageIndex, pageSize)
	data, err := json.Marshal(paginatedList)
	if err == nil {
		redisClient.Set(ctx, cacheKey, data, time.Hour)
	}

	n.log.Infof("Successfully fetched paginated news list with total count: %d", total)
	return paginatedList, nil
}

// GetAllNews retrieves all news from the database
func (n *NewsRepository) GetAllNews(ctx context.Context) ([]BusinessObjects.News, error) {
	n.log.Info("Fetching all news")

	db, err := n.db.GetDB()
	if err != nil {
		n.log.Error("Failed to connect to PostgreSQL:", err)
		return nil, err
	}

	var news []BusinessObjects.News
	if err := db.Find(&news).Error; err != nil {
		n.log.Error("Failed to fetch all news:", err)
		return nil, err
	}

	n.log.Info("Successfully fetched all news")
	return news, nil
}

// GetNewByID retrieves a news by its ID
func (n *NewsRepository) GetNewByID(ctx context.Context, newsID string) (BusinessObjects.News, error) {
	n.log.Infof("Fetching news by ID: %s", newsID)
	redisClient, err := n.db.GetRedis()
	if err != nil {
		n.log.Error("Failed to get Redis client:", err)
		return BusinessObjects.News{}, err
	}

	cacheKey := getNewsCacheKey(newsID)
	val, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		n.log.Infof("Cache hit for news ID: %s", newsID)
		var news BusinessObjects.News
		if err := json.Unmarshal([]byte(val), &news); err == nil {
			return news, nil
		}
		n.log.Warn("Failed to unmarshal cached news data:", err)
	}

	n.log.Infof("Cache miss for news ID: %s", newsID)
	db, err := n.db.GetDB()
	if err != nil {
		n.log.Error("Failed to connect to PostgreSQL:", err)
		return BusinessObjects.News{}, err
	}

	var news BusinessObjects.News
	if err := db.First(&news, "news_id = ?", newsID).Error; err != nil {
		n.log.Error("Failed to fetch news by ID:", err)
		return BusinessObjects.News{}, err
	}

	data, err := json.Marshal(news)
	if err == nil {
		redisClient.Set(ctx, cacheKey, data, time.Hour)
	}

	n.log.Infof("Successfully fetched news by ID: %s", newsID)
	return news, nil
}

// CreateNew adds a new news to the database
func (n *NewsRepository) CreateNew(ctx context.Context, news BusinessObjects.News) error {
	n.log.Infof("Creating new news with title: %s", news.Title)
	db, err := n.db.GetDB()
	if err != nil {
		n.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Create(&news).Error; err != nil {
		n.log.Error("Failed to create new news:", err)
		return err
	}

	n.log.Infof("Successfully created new news with title: %s", news.Title)
	return nil
}

// UpdateNew updates an existing news
func (n *NewsRepository) UpdateNew(ctx context.Context, news BusinessObjects.News) error {
	n.log.Infof("Updating news with ID: %s", news.NewsID)
	db, err := n.db.GetDB()
	if err != nil {
		n.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Save(&news).Error; err != nil {
		n.log.Error("Failed to update news:", err)
		return err
	}

	n.log.Infof("Successfully updated news with ID: %s", news.NewsID)
	return nil
}

// DeleteNew removes a news from the database by its ID
func (n *NewsRepository) DeleteNew(ctx context.Context, newsID string) error {
	n.log.Infof("Deleting news with ID: %s", newsID)
	db, err := n.db.GetDB()
	if err != nil {
		n.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Model(&BusinessObjects.News{}).Where("news_id = ?", newsID).Update("status", false).Error; err != nil {
		n.log.Error("Failed to delete news:", err)
		return err
	}

	n.log.Infof("Successfully deleted news with ID: %s", newsID)
	return nil
}
