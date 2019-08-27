package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	var err error

	connString := "host=db user=postgres password=secret dbname=dev port=5432 sslmode=disable"
	db, err = sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	fmt.Println("DB connected.")
}

type ArticleData struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type ArticleDataIdOnly struct {
	Id int64 `json:"id"`
}

type ArticleError struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type CreateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type CreateArticleResponse struct {
	Status  int64             `json:"status"`
	Message string            `json:"message"`
	Data    ArticleDataIdOnly `json:"data"`
}

type ReadArticleResponse struct {
	Status  int64       `json:"status"`
	Message string      `json:"message"`
	Data    ArticleData `json:"data"`
}

type ListArticlesResponse struct {
	Status  int64         `json:"status"`
	Message string        `json:"message"`
	Data    []ArticleData `json:"data"`
}

func ArticlesCreate(car CreateArticleRequest) (CreateArticleResponse, error) {
	var data CreateArticleResponse

	article := ArticleData{}
	articleId := ArticleDataIdOnly{}

	// Pass the request values
	article.Title = car.Title
	article.Content = car.Content
	article.Author = car.Author

	// Check if any are blank
	if article.Title == "" || article.Content == "" || article.Author == "" {
		// Not acceptable
		return data, errors.New("400. Bad request. All fields must be complete.")
	}

	// Prepare query
	query := "INSERT INTO articles (title, content, author) VALUES ($1, $2, $3) RETURNING id"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(article.Title, article.Content, article.Author).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	articleId.Id = id

	// Prepare data
	data.Status = http.StatusOK
	data.Message = "Success"
	data.Data = articleId

	return data, nil
}

func ArticlesList() (ListArticlesResponse, error) {
	var data ListArticlesResponse

	// Execute query
	rows, err := db.Query("SELECT * FROM articles")
	if err != nil {
		return data, err
	}
	defer rows.Close()

	// Collect results
	articles := make([]ArticleData, 0)
	for rows.Next() {
		article := ArticleData{}
		err := rows.Scan(&article.Id, &article.Title, &article.Content, &article.Author)
		if err != nil {
			return data, err
		}
		articles = append(articles, article)
	}

	if err = rows.Err(); err != nil {
		return data, err
	}

	// Prepare data
	data.Status = http.StatusOK
	data.Message = "Success"
	data.Data = articles

	return data, nil
}

func ArticlesRead(id string) (ReadArticleResponse, error) {
	var data ReadArticleResponse
	articleData := ArticleData{}

	// Make sure there is data
	if id == "" {
		return data, errors.New("400. Bad Request.")
	}

	// Execute query
	row := db.QueryRow("SELECT * FROM articles WHERE id = $1", id)

	err := row.Scan(&articleData.Id, &articleData.Title, &articleData.Content, &articleData.Author)
	if err != nil {
		return data, err
	}

	// Prepare data
	data.Status = http.StatusOK
	data.Message = "Success"
	data.Data = articleData

	return data, nil
}
