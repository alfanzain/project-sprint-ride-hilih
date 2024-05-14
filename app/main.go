package main

import (
	"github.com/alfanzain/project-sprint-halo-suster/app/databases"
	"github.com/alfanzain/project-sprint-halo-suster/app/http"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	databases.ConnectPostgre()

	h := http.New(&http.Http{})
	h.Launch()
}
