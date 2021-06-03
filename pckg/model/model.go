package model

import (
	"database/sql"
	"log"
)

type Question struct {
	ID          string `json:"id,omitempty"`
	Content     string `json:"content"`
	Description string `json:"description"`
	UserCreated string `json:"user_created,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UserUpdated string `json:"user_updated,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	Answer      string `json:"answer,omitempty"`
}

func EnsureTableExists(db *sql.DB) {
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS questions
(
    id SERIAL,
    content TEXT NOT NULL,
	description TEXT NOT NULL,
	answer TEXT NULL,
	createdat timestamp(0) NOT NULL,
    usercreated TEXT NOT NULL,
	updatedat timestamp(0) NULL,
	useredited TEXT NULL,
    CONSTRAINT questions_pkey PRIMARY KEY (id)
)`

func (p *Question) GetQuestion(db *sql.DB) error {
	/* return db.QueryRow(`SELECT id, 'content', description, COALESCE(CAST(answer AS VARCHAR), '') answer,
	 TO_CHAR(createdat, 'dd/mm/yyyy HH24:MI:SS'), usercreated FROM questions WHERE id = $1`,
		p.ID).Scan(&p.ID, &p.Content, &p.Description, &p.Answer, &p.CreatedAt, &p.UserCreated) */

	return db.QueryRow(`SELECT id, content, description, COALESCE(CAST(answer AS VARCHAR), '') answer,
	 TO_CHAR(createdat, 'dd/mm/yyyy HH24:MI:SS'), usercreated FROM questions WHERE id = $1`,
		&p.ID).Scan(&p.ID, &p.Content, &p.Description, &p.Answer, &p.CreatedAt, &p.UserCreated)
}

func (p *Question) UpdateQuestion(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE questions SET content = $1, description = $2, answer = $3, updatedat = NOW() WHERE id = $4",
			p.Content, p.Description, p.Answer, p.ID)

	return err
}

func (p *Question) DeleteQuestion(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM questions where id = $1", p.ID)

	return err
}

func (p *Question) CreateQuestion(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO questions(content, description, answer, createdat, usercreated) VALUES($1, $2, $3, Now(), $4) RETURNING id",
		p.Content, p.Description, p.Answer, p.UserCreated).Scan(&p.ID)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func GetQuestions(db *sql.DB, start, count int) ([]Question, error) {
	rows, err := db.Query(
		`SELECT id, content, description, 
		COALESCE(answer, '') answer , 
		TO_CHAR(createdat, 'dd/mm/yyyy HH24:MI:SS'), usercreated,
		COALESCE(CAST(updatedat AS VARCHAR), '') updatedat, 
		COALESCE(CAST(useredited AS VARCHAR), '') userupdated
		FROM public.questions LIMIT $1 OFFSET $2`,
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	questions := []Question{}

	for rows.Next() {
		var p Question
		if err := rows.Scan(&p.ID, &p.Content, &p.Description, &p.Answer, &p.CreatedAt, &p.UserCreated, &p.UpdatedAt, &p.UserUpdated); err != nil {
			return nil, err
		}
		questions = append(questions, p)
	}

	return questions, nil
}
