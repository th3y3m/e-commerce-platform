package Repositories

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetAllNews() ([]BusinessObjects.News, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM News")
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}
	defer rows.Close()

	news := []BusinessObjects.News{}
	for rows.Next() {
		var newsItem BusinessObjects.News
		err := rows.Scan(&newsItem.NewsID, &newsItem.Title, &newsItem.Content, &newsItem.PublishedDate, &newsItem.AuthorID, &newsItem.Status, &newsItem.ImageURL, &newsItem.Category)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		news = append(news, newsItem)
	}

	return news, nil
}

func GetNewsById(id string) (BusinessObjects.News, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.News{}, err
	}
	defer db.Close()

	var news BusinessObjects.News
	err = db.QueryRow("SELECT * FROM News WHERE NewsID = ?", id).Scan(&news.NewsID, &news.Title, &news.Content, &news.PublishedDate, &news.AuthorID, &news.Status, &news.ImageURL, &news.Category)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.News{}, err
	}

	return news, nil
}

func CreateNews(news BusinessObjects.News) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO News (Title, Content, DatePublished) VALUES (?, ?, ?)", news.Title, news.Content, news.PublishedDate, news.AuthorID, news.Status, news.ImageURL, news.Category)
	if err != nil {
		log.Fatalf("Error inserting into the database: %v", err)
		return err
	}

	return nil
}

func UpdateNews(news BusinessObjects.News) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE News SET Title = ?, Content = ?, DatePublished = ? WHERE NewsID = ?", news.Title, news.Content, news.PublishedDate, news.AuthorID, news.Status, news.ImageURL, news.Category, news.NewsID)
	if err != nil {
		log.Fatalf("Error updating the database: %v", err)
		return err
	}

	return nil
}

func DeleteNews(news BusinessObjects.News) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM News WHERE NewsID = ?", news.NewsID)
	if err != nil {
		log.Fatalf("Error deleting from the database: %v", err)
		return err
	}

	return nil
}
