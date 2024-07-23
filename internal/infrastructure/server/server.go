package server

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Run(db *sql.DB) {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, World!")
    })

    fmt.Println("Server is running at :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
