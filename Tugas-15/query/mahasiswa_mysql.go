package query

import (
	"context"
	"fmt"
	"Tugas-15/config"
	"Tugas-15/models"
	"log"
	"errors"
	"database/sql"
)

const (
	mhs_table = "mahasiswa"
)

func GetMhs(ctx context.Context) ([]models.Mahasiswa, error){
	var mahasiswa []models.Mahasiswa
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	query := fmt.Sprintf("SELECT * FROM %v", mhs_table)

	rowQuery, err := db.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var mhs models.Mahasiswa
		if err = rowQuery.Scan(&mhs.ID,
			&mhs.Nama); err != nil {
				return nil, err
			}

		if err != nil {
			log.Fatal(err)
		}
		mahasiswa = append(mahasiswa, mhs)
	}	
	return mahasiswa, nil
}

func InsertMhs(ctx context.Context, mahasiswa models.Mahasiswa) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	query := fmt.Sprintf("INSERT INTO %v (nama) values('%v')", 
	mhs_table, mahasiswa.Nama)
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}

func UpdateMhs(ctx context.Context, mahasiswa models.Mahasiswa, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	query := fmt.Sprintf("UPDATE %v set nama = '%s' where id = %s", 
	mhs_table, mahasiswa.Nama, id)
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}

func DeleteMhs(ctx context.Context, idMhs string) error {
	db, err := config.MySQL()
    if err != nil {
        log.Fatal("Can't connect to MySQL", err)
    }

    queryText := fmt.Sprintf("DELETE FROM %v where id = %s", mhs_table, idMhs)

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
