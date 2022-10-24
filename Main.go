package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func main() {
	handleRequests()
}

func (article Article) test(fn func(article Article) error) {
	fn(Article{Title: "Salam", Desc: "Some Description", Content: "Content1"})
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/log-android", getLogs).Methods("POST")
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles/{title}", getArticle).Methods("GET")
	myRouter.HandleFunc("/articles", allPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func getLogs(writer http.ResponseWriter, request *http.Request) {
	var p Log
	json.NewDecoder(request.Body).Decode(&p)
	fmt.Println("key:", p.Key, " value:", p.Value, "device:", p.Device)
	fmt.Fprintf(writer, p.Value)
}

type Log struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Device string `json:"device"`
}

func CheckDatabase() {
	db, err := sql.Open("mysql", "nima/mysql:8.0.30:secret@(127.0.0.1:5430)/dbname?parseTime=true")
	//err := db.Ping()
	if err != nil {
		fmt.Println("here")
		if db == nil {
			fmt.Println("db is null")
		} else {
			fmt.Println(db.Ping())
		}
	} else {
		fmt.Println("Error is:", err)
	}
}

func getArticle(writer http.ResponseWriter, request *http.Request) {
	articles := Articles{
		Article{Title: "Salam", Desc: "Some Description", Content: "Content1"},
		Article{Title: "Hi", Desc: "Some other Description", Content: "Content2"},
	}
	vars := mux.Vars(request)
	title := vars["title"]
	for i, article := range articles {
		if article.Title == title {
			json.NewEncoder(writer).Encode(articles[i])
			return
		}
	}
	fmt.Fprintf(writer, "NOT FOUND!")

}

func allPostArticles(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Post SuccessFully Done!")
}

func allArticles(writer http.ResponseWriter, request *http.Request) {
	token := request.URL.Query().Get("token")
	if len(token) != 0 {
		articles := Articles{
			Article{Title: "Salam", Desc: "Some Description", Content: "Content1"},
			Article{Title: "Hi", Desc: "Some other Description", Content: "Content2"},
		}
		fmt.Println("EndPoint Hit!")
		err := json.NewEncoder(writer).Encode(articles)
		if err != nil {
			fmt.Fprint(writer, err)
			return
		}
	} else {
		fmt.Fprint(writer, "Wrong Token!!")
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Home Page End Point! %s", request.URL.Path)
}
