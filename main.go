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

// createArticle will create an article
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

		// Output data
		displayOutput(w, getErrorString(errData))
		return
	}

	data, err := json.Marshal(article)
	if err != nil {
		log.Fatalln(err)
	}

	// Output data
	displayOutput(w, string(data))
}

// displayOutput will display the output
func displayOutput(w http.ResponseWriter, data string) {
	fmt.Fprintf(w, data)
	fmt.Fprintf(w, "\n")
}

// getErrorString will get the error string from an articleError
func getErrorString(articleError ArticleError) string {
	data, err := json.Marshal(articleError)
	if err != nil {
		log.Fatalln(err)
	}

	return string(data)
}

// listArticles will list the articles
func listArticles(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	articles, err := ArticlesList()
	if err != nil {
		errData := ArticleError{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(500),
			Data:    "",
		}

		// Output data
		displayOutput(w, getErrorString(errData))
		return
	}
	data, err := json.Marshal(articles)
	if err != nil {
		log.Fatalln(err)
	}

	// Output data
	displayOutput(w, string(data))
}

// readArticle will read an article based in the id given
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

		// Output data
		displayOutput(w, getErrorString(errData))
		return
	case err != nil:
		errData := ArticleError{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(500),
			Data:    "",
		}

		// Output data
		displayOutput(w, getErrorString(errData))
		return
	}

	data, err := json.Marshal(article)
	if err != nil {
		log.Fatalln(err)
	}

	// Output data
	displayOutput(w, string(data))
}

// main is the entry point of the program
func main() {
	mux := httprouter.New()

	mux.GET("/articles", listArticles)
	mux.GET("/articles/:article_id", readArticle)
	mux.POST("/articles", createArticle)

	http.ListenAndServe(":8080", mux)
}
