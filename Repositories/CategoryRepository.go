package Repositories

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetAllCategories() ([]BusinessObjects.Category, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Categories")
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}
	defer rows.Close()

	categories := []BusinessObjects.Category{}

	for rows.Next() {
		var category BusinessObjects.Category
		err := rows.Scan(&category.CategoryID, &category.CategoryName)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func GetCategoryById(id string) (BusinessObjects.Category, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.Category{}, err
	}

	defer db.Close()

	var category BusinessObjects.Category

	err = db.QueryRow("SELECT * FROM Categories WHERE CategoryID = ?", id).Scan(&category.CategoryID, &category.CategoryName)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.Category{}, err
	}

	return category, nil
}

func CreateCategory(category BusinessObjects.Category) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Categories (CategoryID, CategoryName) VALUES (?, ?)", category.CategoryID, category.CategoryName)
	if err != nil {
		log.Fatalf("Error inserting into the database: %v", err)
		return err
	}

	return nil
}

func UpdateCategory(category BusinessObjects.Category) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE Categories SET CategoryName = ? WHERE CategoryID = ?", category.CategoryName, category.CategoryID)
	if err != nil {
		log.Fatalf("Error updating the database: %v", err)
		return err
	}

	return nil
}

func DeleteCategory(id string) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Categories WHERE CategoryID = ?", id)
	if err != nil {
		log.Fatalf("Error deleting from the database: %v", err)
		return err
	}

	return nil
}
