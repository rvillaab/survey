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

type QuestionDao struct {
	DB *sql.DB
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

func (qd *QuestionDao) GetQuestion(p *Question) error {
	return qd.DB.QueryRow(`SELECT id, content, description, COALESCE(CAST(answer AS VARCHAR), '') answer,
	 TO_CHAR(createdat, 'dd/mm/yyyy HH24:MI:SS'), usercreated FROM questions WHERE id = $1`,
		&p.ID).Scan(&p.ID, &p.Content, &p.Description, &p.Answer, &p.CreatedAt, &p.UserCreated)
}

func (qd *QuestionDao) UpdateQuestion(p *Question) error {
	_, err :=
		qd.DB.Exec("UPDATE questions SET content = $1, description = $2, answer = $3, updatedat = NOW() WHERE id = $4",
			p.Content, p.Description, p.Answer, p.ID)

	return err
}

func (qd *QuestionDao) DeleteQuestion(id string) error {
	_, err := qd.DB.Exec("DELETE FROM questions where id = $1", id)

	return err
}

func (qd *QuestionDao) CreateQuestion(p *Question) error {
	err := qd.DB.QueryRow(
		"INSERT INTO questions(content, description, answer, createdat, usercreated) VALUES($1, $2, $3, Now(), $4) RETURNING id",
		p.Content, p.Description, p.Answer, p.UserCreated).Scan(&p.ID)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (qd *QuestionDao) GetQuestions(start, count int) ([]Question, error) {
	rows, err := qd.DB.Query(
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
