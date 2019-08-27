package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
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

func Test_displayError(t *testing.T) {
	type args struct {
		articleError ArticleError
		w            http.ResponseWriter
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
				w: nil,
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
				w: nil,
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
			if got := displayError(tt.args.articleError, tt.args.w); reflect.DeepEqual(got, tt.want) {
				t.Errorf("displayError() = %v, don't want %v", got, tt.want)
			}
		})
	}
}

func Test_listArticles(t *testing.T) {
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

func Test_readArticle(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
		ps  httprouter.Params
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