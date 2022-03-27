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
	book_table = "books"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAllBooks(ctx context.Context) ([]models.Books, error){
	var books []models.Books
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	query := fmt.Sprintf("SELECT * FROM %v", book_table)

	rowQuery, err := db.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book models.Books
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageUrl,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&createdAt,
			&updatedAt,
			&book.CategoryId); err != nil {
				return nil, err
			}

		book.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		book.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}	
	return books, nil
}

func InsertBook(ctx context.Context, book models.Books) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	query := fmt.Sprintf("INSERT INTO %v (title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id) values('%v', '%v', '%v', %v, %v, %v, '%v', NOW(), NOW(), %v)", 
	book_table, book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryId)
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}

func UpdateBook(ctx context.Context, book models.Books, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	query := fmt.Sprintf("UPDATE %v set title = '%s', description = '%s', image_url = '%s', release_year = %d, price = '%s', total_page = '%s', thickness = '%s', updated_at = NOW(), category_id = %d where id = %d",
		book_table,
		book.Title,
		book.Description,
		book.ImageUrl,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CategoryId,
		id)
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(ctx context.Context, id string) error {
	db, err := config.MySQL()
    if err != nil {
        log.Fatal("Can't connect to MySQL", err)
    }

    queryText := fmt.Sprintf("DELETE FROM %v where id = %s", book_table, id)

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

func GetBooksByCategory(ctx context.Context, id string) ([]models.Books, error){
	var books []models.Books
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	query := fmt.Sprintf("SELECT * FROM %v WHERE category_id = %d", book_table, id)

	rowQuery, err := db.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book models.Books
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageUrl,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&createdAt,
			&updatedAt,
			&book.CategoryId); err != nil {
				return nil, err
			}

		book.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		book.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}	
	return books, nil
}