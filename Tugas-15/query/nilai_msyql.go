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
	nilai_table = "nilai"
)

func GetNilai(ctx context.Context) ([]models.NilaiMahasiswa, error){
	var nilaiMhs []models.NilaiMahasiswa
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	query := fmt.Sprintf("SELECT * FROM %v", nilai_table)

	rowQuery, err := db.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var nilai models.NilaiMahasiswa
		if err = rowQuery.Scan(&nilai.ID,
			&nilai.Indeks,
			&nilai.Skor,
			&nilai.MahasiswaId,
			&nilai.MataKuliahId); err != nil {
				return nil, err
			}

		if err != nil {
			log.Fatal(err)
		}
		nilaiMhs = append(nilaiMhs, nilai)
	}	
	return nilaiMhs, nil
}

func InsertNilai(ctx context.Context, nilai models.NilaiMahasiswa) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	query := fmt.Sprintf("INSERT INTO %v (indeks, skor, mahasiswa_id, mata_kuliah_id) values('%v', %v, %v, %v)", 
	nilai_table, nilai.Indeks, nilai.Skor, nilai.MahasiswaId, nilai.MataKuliahId)
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}

func UpdateNilai(ctx context.Context, nilai models.NilaiMahasiswa, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	query := fmt.Sprintf("UPDATE %v set indeks = '%s', skor = %d, mahasiswa_id = %d, mata_kuliah_id = %d where id = %s", 
	nilai_table, nilai.Indeks, nilai.Skor, nilai.MahasiswaId, nilai.MataKuliahId, id)
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}

func DeleteNilai(ctx context.Context, id string) error {
	db, err := config.MySQL()
    if err != nil {
        log.Fatal("Can't connect to MySQL", err)
    }

    queryText := fmt.Sprintf("DELETE FROM %v where id = %s", nilai_table, id)

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