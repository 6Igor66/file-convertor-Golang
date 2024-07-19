package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	Db *sql.DB
}

func NewPostgres(storagePath string) (*PostgreSQL, error) {
	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, err
	}

	return &PostgreSQL{Db: db}, nil
}

func (p *PostgreSQL) CreateUser(tgID int64) error {
	query := "INSERT INTO Users ( telegram_id ) VALUES ($1 ) ON CONFLICT DO NOTHING"
	_, err := p.Db.Exec(query, tgID)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgreSQL) SetPaylaod(tgID int64, payload string) error {
	query := "UPDATE Users SET payload=$1 WHERE telegram_id=$2"
	_, err := p.Db.Exec(query, payload, tgID)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgreSQL) UpdateMessageStatus(tgID int64, status string) error {
	query := "UPDATE Users SET message_status=$1 WHERE telegram_id=$2"
	_, err := p.Db.Exec(query, status, tgID)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgreSQL) GetMessageStatus(tgID int64) (string, error) {
	var status string
	query := "SELECT message_status FROM Users WHERE telegram_id = $1"
	row := p.Db.QueryRow(query, tgID)

	err := row.Scan(&status)
	if err != nil {
		return "", err
	}
	return status, nil
}

func (p *PostgreSQL) GetPayload(tgID int64) (string, error) {
	var payload string
	query := "SELECT payload FROM Users WHERE telegram_id = $1"
	row := p.Db.QueryRow(query, tgID)

	err := row.Scan(&payload)
	if err != nil {
		return "", err
	}
	return payload, nil
}

func (p *PostgreSQL) SetOperations(tgID int64, operations int) error {
	query := "UPDATE Users SET operations=$1 WHERE telegram_id=$2"
	_, err := p.Db.Exec(query, operations, tgID)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgreSQL) GetOperations(tgID int64) (int, error) {
	var operations int
	query := "SELECT operations FROM Users WHERE telegram_id = $1"
	row := p.Db.QueryRow(query, tgID)

	err := row.Scan(&operations)
	if err != nil {
		return 0, err
	}
	return operations, nil
}

func (p *PostgreSQL) SetFileName(tgID int64, fileName string) error {
	query := "UPDATE Users SET file_name=$1 WHERE telegram_id=$2"
	_, err := p.Db.Exec(query, fileName, tgID)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgreSQL) GetFileName(tgID int64) (string, error) {
	var fileName string
	query := "SELECT file_name FROM Users WHERE telegram_id = $1"
	row := p.Db.QueryRow(query, tgID)

	err := row.Scan(&fileName)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
