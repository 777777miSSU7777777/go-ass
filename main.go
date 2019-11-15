package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"

	"github.com/777777miSSU7777777/go-ass/api"
	"github.com/777777miSSU7777777/go-ass/repository"
	"github.com/777777miSSU7777777/go-ass/service"
	"github.com/777777miSSU7777777/go-ass/stream"
)

func renderIndex(w http.ResponseWriter, r *http.Request) {
	pagePath := path.Join("frontend", "index.html")

	htmlPage, err := template.ParseFiles(pagePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = htmlPage.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	var connectionString string
	var baseLocation string
	var apiOnly bool

	flag.StringVar(&connectionString, "connection_string", "", "MySQL connection string")
	homePath := os.Getenv("HOME")
	flag.StringVar(&baseLocation, "storage_location", homePath+"/goass/storage", "Storage location")
	flag.BoolVar(&apiOnly, "api_only", true, "Run only api without frontend")
	flag.Parse()

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repo := repository.New(db)
	svc := service.New(repo)
	u := api.NewFileManager(baseLocation)
	m := stream.NewMediaManager(baseLocation)
	apiHandlers := api.NewApi(svc, u)
	streamHandlers := stream.NewStreamAPI(m)

	r := mux.NewRouter()

	api.NewAPIRouter(r, apiHandlers)
	stream.NewStreamRouter(r, streamHandlers)

	r.Handle("/health-check", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))

	if (!apiOnly){
		r.Path("/").HandlerFunc(renderIndex)

		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/"))))
	}

	http.Handle("/", r)

	fmt.Println("started")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
