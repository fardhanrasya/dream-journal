package repository

import (
	"database/sql"
	"time"

	"fardhan.dev/dreamjournal/internal/db"
	"fardhan.dev/dreamjournal/internal/model"
)

type DreamRepository struct {
	DB *sql.DB
}

func NewDreamRepository(db *sql.DB) *DreamRepository {
	return &DreamRepository{DB: db}
}

func (r *DreamRepository) CreateDream(dream *model.Dream) error {
	query := `INSERT INTO dreams (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if dream.CreatedAt.IsZero() {
		dream.CreatedAt = time.Now()
	}
	if dream.UpdatedAt.IsZero() {
		dream.UpdatedAt = time.Now()
	}

	result, err := stmt.Exec(dream.Title, dream.Content, dream.CreatedAt, dream.UpdatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	dream.ID = int(id)
	return nil
}

func (r *DreamRepository) GetDreams() ([]model.Dream, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM dreams ORDER BY created_at DESC`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dreams []model.Dream
	for rows.Next() {
		var d model.Dream
		if err := rows.Scan(&d.ID, &d.Title, &d.Content, &d.CreatedAt, &d.UpdatedAt); err != nil {
			return nil, err
		}
		dreams = append(dreams, d)
	}
	return dreams, nil
}

func (r *DreamRepository) GetDreamByID(id int) (*model.Dream, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM dreams WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var d model.Dream
	if err := row.Scan(&d.ID, &d.Title, &d.Content, &d.CreatedAt, &d.UpdatedAt); err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *DreamRepository) UpdateDream(dream *model.Dream) error {
	query := `UPDATE dreams SET title = ?, content = ?, updated_at = ? WHERE id = ?`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	dream.UpdatedAt = time.Now()
	_, err = stmt.Exec(dream.Title, dream.Content, dream.UpdatedAt, dream.ID)
	return err
}

func (r *DreamRepository) DeleteDream(id int) error {
	query := `DELETE FROM dreams WHERE id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *DreamRepository) SearchDreams(query string) ([]model.Dream, error) {
	sqlQuery := `SELECT id, title, content, created_at, updated_at FROM dreams WHERE title LIKE ? OR content LIKE ? ORDER BY created_at DESC`
	searchTerm := "%" + query + "%"
	rows, err := r.DB.Query(sqlQuery, searchTerm, searchTerm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dreams []model.Dream
	for rows.Next() {
		var d model.Dream
		if err := rows.Scan(&d.ID, &d.Title, &d.Content, &d.CreatedAt, &d.UpdatedAt); err != nil {
			return nil, err
		}
		dreams = append(dreams, d)
	}
	return dreams, nil
}
