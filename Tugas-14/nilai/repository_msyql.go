package nilai

import (
	"context"
	"fmt"
	"Tugas-14/config"
	"Tugas-14/models"
	"log"
	"errors"
	"database/sql"
)

const (
	table = "nilaimhs"
)

func GetAll(ctx context.Context) ([]models.NilaiMahasiswa, error){
	var nilaiMhs []models.NilaiMahasiswa
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	query := fmt.Sprintf("SELECT * FROM %v", table)

	rowQuery, err := db.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var nilai models.NilaiMahasiswa
		if err = rowQuery.Scan(&nilai.ID,
			&nilai.Nama,
			&nilai.MataKuliah,
			&nilai.Nilai,
			&nilai.IndeksNilai); err != nil {
				return nil, err
			}

		if err != nil {
			log.Fatal(err)
		}
		nilaiMhs = append(nilaiMhs, nilai)
	}	
	return nilaiMhs, nil
}

func Insert(ctx context.Context, nilai models.NilaiMahasiswa) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	query := fmt.Sprintf("INSERT INTO %v (nama, mata_kuliah, nilai, indeks_nilai) values('%v', '%v', %v, '%v')", 
	table, nilai.Nama, nilai.MataKuliah, nilai.Nilai, nilai.IndeksNilai)
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}

func Update(ctx context.Context, nilai models.NilaiMahasiswa, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	query := fmt.Sprintf("UPDATE %v set nama = '%s', mata_kuliah = '%s', nilai = %d, indeks_nilai = '%s' where id = %s", 
	table, nilai.Nama, nilai.MataKuliah, nilai.Nilai, nilai.IndeksNilai, id)
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}

func Delete(ctx context.Context, idMhs string) error {
	db, err := config.MySQL()
    if err != nil {
        log.Fatal("Can't connect to MySQL", err)
    }

    queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, idMhs)

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