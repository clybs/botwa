package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"testing"
)

func TestArticlesCreateSuccess(t *testing.T) {
	DB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer DB.Close()

	mock.ExpectPrepare("INSERT INTO articles").ExpectExec().WithArgs("Hello", "World", "Cliburn").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectClose()
	//mock.ExpectQuery("SELECT * WHERE title='Hello' AND content='World' AND author='Cliburn'")

	db = DB

	// now we execute our method
	car := CreateArticleRequest{
		Title:   "Hello",
		Content: "World",
		Author:  "Cliburn",
	}
	if _, err := ArticlesCreate(car); err != nil {
		t.Errorf("Error was not expected while creating article: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestArticlesList(t *testing.T) {
	tests := []struct {
		name    string
		want    ListArticlesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ArticlesList()
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticlesList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticlesList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticlesReadSuccess(t *testing.T) {

}
