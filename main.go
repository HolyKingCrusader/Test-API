package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Article struct {
	Id      string `json:Id`
	Title   string `json:Title`
	Desc    string `json:Desc`
	Content string `json:Content`
}

var Articles []Article

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "dupadupa123"
	dbname   = "articles"
)

var db *sql.DB

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page.")
	fmt.Println("Endpoint hit: homePage")
}

func returnArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllArticles")

	arts := make([]*Article, 0)
	query := "SELECT * FROM articlestable"

	if r.URL.Path != "/articles" {
		vars := mux.Vars(r)
		key := vars["id"]
		query = query + " WHERE id = " + key
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		art := new(Article)
		err := rows.Scan(&art.Id, &art.Title, &art.Desc, &art.Content)
		if err != nil {
			log.Fatal(err)
		}
		arts = append(arts, art)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(arts)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: createNewArticle")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)

	Id := article.Id
	Title := article.Title
	Desc := article.Desc
	Content := article.Content
	if Id == "" || Title == "" || Desc == "" || Content == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	result, err := db.Exec("INSERT INTO articlestable VALUES($1, $2, $3, $4)", Id, Title, Desc, Content)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Printf("Cannot create an article: %s\n", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "Article %s created successfully. Rows Affected: %d\n", Id, rowsAffected)

}

func deleteArticle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)

	result, err := db.Exec("DELETE FROM articlestable WHERE id= " + key)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Printf("Cannot delete an article: %s\n", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	fmt.Fprintf(w, "Article %s deleted successfully. Rows affected: %d \n", key, rowsAffected)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)

	Title := article.Title
	Desc := article.Desc
	Content := article.Content
	if Title == "" || Desc == "" || Content == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	result, err := db.Exec("UPDATE articlestable SET title=$1, description=$2, content=$3 WHERE id= "+key, Title, Desc, Content)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Printf("Cannot modify an article: %s\n", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "Article %s modified successfully. Rows affected: %d \n", key, rowsAffected)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnArticles)
	myRouter.HandleFunc("/articles/create", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/articles/{id}", returnArticles)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	handleRequests()
}
