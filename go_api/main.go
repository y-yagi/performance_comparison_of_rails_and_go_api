package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"

	"goji.io"
	"goji.io/pat"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	db *gorm.DB
)

func getUsers(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var users []User
	var buffer bytes.Buffer
	db.Find(&users)

	mappage, _ := json.Marshal(users)
	buffer.WriteString(string(mappage))

	fmt.Fprint(w, buffer.String())
}

func route(mux *goji.Mux) {
	mux.HandleFuncC(pat.Get("/users"), getUsers)
}

func main() {
	var err error

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "dbname=performance_comparison"
	}

	db, err = gorm.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	mux := goji.NewMux()
	route(mux)

	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalln(err)
	}
}
