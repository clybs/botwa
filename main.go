package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func createArticle(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// Get request data
	decoder := json.NewDecoder(req.Body)
	var articleData CreateArticleRequest
	err := decoder.Decode(&articleData)
	if err != nil {
		panic(err)
	}

	// Create the article
	article, err := ArticlesCreate(articleData)
	if err != nil {
		errData := ArticleError{
			Status:  http.StatusNotAcceptable,
			Message: http.StatusText(406),
			Data:    "",
		}
		fmt.Fprintf(w, displayError(errData, w))
		fmt.Fprintf(w, "\n")
		return
	}

	data, err := json.Marshal(article)
	if err != nil {
		log.Fatalln(err)
	}

	// Output results
	fmt.Fprintf(w, string(data))
	fmt.Fprintf(w, "\n")
}

func displayError(articleError ArticleError, w http.ResponseWriter) string {
	data, err := json.Marshal(articleError)
	if err != nil {
		log.Fatalln(err)
	}

	return string(data)
}

func listArticles(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	articles, err := ArticlesList()
	if err != nil {
		errData := ArticleError{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(500),
			Data:    "",
		}
		fmt.Fprintf(w, displayError(errData, w))
		fmt.Fprintf(w, "\n")
		return
	}
	data, err := json.Marshal(articles)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, string(data))
	fmt.Fprintf(w, "\n")
}

func readArticle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// Get article id to read
	article, err := ArticlesRead(ps.ByName("article_id"))
	switch {
	case err == sql.ErrNoRows:
		errData := ArticleError{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    "",
		}
		fmt.Fprintf(w, displayError(errData, w))
		fmt.Fprintf(w, "\n")
		return
	case err != nil:
		errData := ArticleError{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(500),
			Data:    "",
		}
		fmt.Fprintf(w, displayError(errData, w))
		fmt.Fprintf(w, "\n")
		return
	}

	data, err := json.Marshal(article)
	if err != nil {
		log.Fatalln(err)
	}

	// Output data
	fmt.Fprintf(w, string(data))
	fmt.Fprintf(w, "\n")
}

func main() {
	mux := httprouter.New()

	mux.GET("/articles", listArticles)
	mux.GET("/articles/:article_id", readArticle)
	mux.POST("/articles", createArticle)

	http.ListenAndServe(":8080", mux)
}
