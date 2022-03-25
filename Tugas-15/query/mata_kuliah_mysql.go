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
	mk_table = "mata_kuliah"
)

func GetMk(ctx context.Context) ([]models.MataKuliah, error){
	var matkul []models.MataKuliah
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	query := fmt.Sprintf("SELECT * FROM %v", mk_table)

	rowQuery, err := db.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var mk models.MataKuliah
		if err = rowQuery.Scan(&mk.ID,
			&mk.Nama); err != nil {
				return nil, err
			}

		if err != nil {
			log.Fatal(err)
		}
		matkul = append(matkul, mk)
	}	
	return matkul, nil
}

func InsertMk(ctx context.Context, matkul models.MataKuliah) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	query := fmt.Sprintf("INSERT INTO %v (nama) values('%v')", 
	mk_table, matkul.Nama)
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}

func UpdateMk(ctx context.Context, matkul models.MataKuliah, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	query := fmt.Sprintf("UPDATE %v set nama = '%s' where id = %s", 
	mk_table, matkul.Nama, id)
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}

func DeleteMk(ctx context.Context, idMk string) error {
	db, err := config.MySQL()
    if err != nil {
        log.Fatal("Can't connect to MySQL", err)
    }

    queryText := fmt.Sprintf("DELETE FROM %v where id = %s", mk_table, idMk)

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
