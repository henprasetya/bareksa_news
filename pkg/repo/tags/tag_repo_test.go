package tags

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/henprasetya/news/pkg/model"
	"github.com/stretchr/testify/assert"
)

var u = &model.Tags{
	Id:          99,
	Code:        "code test",
	Description: "Description Test",
	IdNews:      98,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	//defer db.Close()
	if err != nil {
		log.Fatal("stub database connection", err)
	}

	return db, mock
}

func TestFindByID(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	repo := &tagsDb{
		db: db,
	}

	query := "select \\* from tags where id = ?"

	rows := sqlmock.NewRows([]string{"id", "code", "description", "news_id"}).
		AddRow(u.Id, u.Code, u.Description, u.IdNews)

	mock.ExpectQuery(query).WithArgs(u.Id).WillReturnRows(rows)

	tags, err := repo.SelectTagsList(u.Id)
	assert.NotNil(t, tags)
	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {
	db, mock := NewMock()
	repo := &tagsDb{db: db}
	defer db.Close()
	u.Id = 0
	query := "insert into users \\(id, code, description, news_id\\) VALUES \\(\\?, \\?, \\?, \\?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.Id, u.Code, u.Description, u.IdNews).WillReturnResult(sqlmock.NewResult(0, 1))

	resp, err := repo.CreateOrUpdateTags(*u)
	assert.NoError(t, err)
	assert.True(t, resp.Success, true)
}
func TestUpdate(t *testing.T) {
	db, mock := NewMock()
	repo := &tagsDb{db: db}
	defer db.Close()
	// u.Id = 0
	query := "update tags set code=\\?, description=\\?, news_id=\\? where id=\\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.Code, u.Description, u.IdNews, u.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	resp, err := repo.CreateOrUpdateTags(*u)
	assert.NoError(t, err)
	assert.True(t, resp.Success, true)
}
