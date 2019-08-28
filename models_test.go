package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func xTestArticlesCreateSuccess(t *testing.T) {
	DB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer DB.Close()

	mock.ExpectPrepare("INSERT INTO articles").ExpectExec().WithArgs("Hello", "World", "Cliburn").WillReturnResult(sqlmock.NewResult(1, 1))
	//mock.ExpectQuery("SELECT * WHERE title='Hello' AND content='World' AND author='Cliburn'").

	db = DB

	car := CreateArticleRequest{
		Title:   "Hello",
		Content: "World",
		Author:  "Cliburn",
	}
	if _, err := ArticlesCreate(car); err != nil {
		t.Errorf("Error was not expected while creating article: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func xTestArticlesList(t *testing.T) {
	DB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer DB.Close()
	mock.NewRows([]string{"id", "title", "content", "author"}).
		AddRow(1, "title 1", "content 1", "author 1").
		AddRow(2, "title 2", "content 2", "author 2")

	rows := sqlmock.NewRows([]string{"id", "title", "content", "author"}).
		AddRow(1, "title 1", "content 1", "author 1").
		AddRow(2, "title 2", "content 2", "author 2")

	mock.ExpectQuery("SELECT * FROM articles").WillReturnRows(rows)
	mock.ExpectClose()

	db = DB

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestArticlesReadSuccess(t *testing.T) {

}
