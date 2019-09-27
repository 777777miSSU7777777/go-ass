package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"

	"github.com/777777miSSU7777777/go-ass/api"
	"github.com/777777miSSU7777777/go-ass/repository"
	"github.com/777777miSSU7777777/go-ass/service"
)

// var locationPattern = "%s/%d/%s/%s"

func main() {
	var connectionString string
	var baseLocation string

	flag.StringVar(&connectionString, "connection_string", "", "MySQL connection string")
	homePath := os.Getenv("HOME")
	flag.StringVar(&baseLocation, "storage_location", homePath+"/goass/storage", "Storage location")
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
	m := api.NewUploadManager(baseLocation)
	apiHandlers := api.NewApi(svc, m)
	// apiRouter := api.NewAPIRouter(apiHandlers)

	r := mux.NewRouter()
	// r.Handle("/api", apiRouter)
	api.NewAPIRouter(r, apiHandlers)
	r.Handle("/health-check", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))

	fmt.Println("started")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
