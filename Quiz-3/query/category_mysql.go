package query

import (
	"context"
	"fmt"
	"Quiz-3/config"
	"Quiz-3/models"
	"log"
	"errors"
	"database/sql"
	"time"
)

const (
	category_table = "books"
	layoutDate = "2006-01-02 15:04:05"
)

func GetAllCategory(ctx context.Context) ([]models.Category, error){
	var categories []models.Category
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	query := fmt.Sprintf("SELECT * FROM %v", category_table)

	rowQuery, err := db.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var category models.Category
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&category.ID,
			&category.Name,
			&createdAt,
			&updatedAt); err != nil {
				return nil, err
			}

		category.CreatedAt, err = time.Parse(layoutDate, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		category.UpdatedAt, err = time.Parse(layoutDate, updatedAt)

		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, category)
	}	
	return categories, nil
}

func InsertCategory(ctx context.Context, category models.Category) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	query := fmt.Sprintf("INSERT INTO %v (name, created_at, updated_at) values('%v', NOW(), NOW())", 
	category_table, category.Name)
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}

func UpdateCategory(ctx context.Context, category models.Category, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	query := fmt.Sprintf("UPDATE %v set name = '%s', updated_at = NOW() where id = %s", 
	category_table, category.Name, id)
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}

func DeleteCategory(ctx context.Context, id string) error {
	db, err := config.MySQL()
    if err != nil {
        log.Fatal("Can't connect to MySQL", err)
    }

    queryText := fmt.Sprintf("DELETE FROM %v where id = %s", category_table, id)

    s, err := db.ExecContext(ctx, queryText)

    if err != nil && err != sql.ErrNoRows {
        return err
    }

    check, err := s.RowsAffected()
    fmt.Println(check)
    if check == 0 {
        return errors.New("id tidak ada")
    }

    if err != nil {
        fmt.Println(err.Error())
    }

    return nil
}

