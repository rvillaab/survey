package test

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"survey/pckg/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestGetQuestion(t *testing.T) {

	db, mock := NewMock()
	quest := &model.Question{ID: "5"}
	defer func() {
		db.Close()
	}()
	query := "SELECT id, content, description, COALESCE\\(CAST\\(answer AS VARCHAR\\), ''\\) answer, TO_CHAR\\(createdat, 'dd/mm/yyyy HH24:MI:SS'\\), usercreated FROM questions WHERE id = \\$1"

	rows := sqlmock.NewRows([]string{"id", "content", "description", "answer", "createdat", "usercreated"}).
		AddRow("5", "Where is the Torre Eifel?", "City", "The Torre Eifel is in Paris", "02/06/2021 15:30:55", "Peter")

	mock.ExpectQuery(query).WithArgs(quest.ID).WillReturnRows(rows)

	questionDao := model.QuestionDao{DB: db}

	err := questionDao.GetQuestion(quest)

	if err != nil {
		t.Fatalf("an error '%s'", err.Error())
	}

	assert.EqualValues(t, "Where is the Torre Eifel?", quest.Content)
	assert.EqualValues(t, "City", quest.Description)
	assert.EqualValues(t, "The Torre Eifel is in Paris", quest.Answer)

}

func TestDelete(t *testing.T) {
	db, mock := NewMock()
	quest := &model.Question{ID: "5"}
	defer func() {
		db.Close()
	}()

	query := "DELETE FROM questions where id = $1"

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(quest.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	questionDao := model.QuestionDao{DB: db}

	err := questionDao.DeleteQuestion(quest.ID)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	db, mock := NewMock()
	quest := &model.Question{ID: "5"}
	defer func() {
		db.Close()
	}()

	query := "UPDATE questions SET content = $1, description = $2, answer = $3, updatedat = NOW() WHERE id = $4"

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(quest.Content, quest.Description, quest.Answer, quest.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	questionDao := model.QuestionDao{DB: db}

	err := questionDao.UpdateQuestion(quest)
	assert.NoError(t, err)
}
