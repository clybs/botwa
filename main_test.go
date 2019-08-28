package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_createArticle(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_getErrorString(t *testing.T) {
	type args struct {
		articleError ArticleError
	}
	tests := []struct {
		name string
		args args
		want ArticleError
	}{
		{
			name: "should display error",
			args: args{
				articleError: ArticleError{
					Status:  http.StatusOK,
					Message: http.StatusText(200),
					Data:    "Some data",
				},
			},
			want: ArticleError{
				Status:  http.StatusOK,
				Message: http.StatusText(200),
				Data:    "",
			},
		},
		{
			name: "should display error",
			args: args{
				articleError: ArticleError{
					Status:  http.StatusBadRequest,
					Message: http.StatusText(400),
					Data:    "Some data",
				},
			},
			want: ArticleError{
				Status:  http.StatusOK,
				Message: http.StatusText(400),
				Data:    "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getErrorString(tt.args.articleError); reflect.DeepEqual(got, tt.want) {
				t.Errorf("getErrorString() = %v, don't want %v", got, tt.want)
			}
		})
	}
}

func Test_listArticles(t *testing.T) {
	router := httprouter.New()
	router.GET("/articles", listArticles)

	req, _ := http.NewRequest("GET", "/articles", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
}

func xTest_readArticle(t *testing.T) {
	DB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer DB.Close()

	mock.NewRows([]string{"id", "title", "content", "author"}).
		AddRow(1, "title 1", "content 1", "author 1").
		AddRow(2, "title 2", "content 2", "author 2")

	mock.ExpectQuery("SELECT * FROM articles WHERE id = 1")
	mock.ExpectClose()

	db = DB

	router := httprouter.New()
	router.GET("/articles/:article_id", readArticle)

	req, _ := http.NewRequest("GET", "/articles/1", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
