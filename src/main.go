package main

import (
	"net/http"

	_ "github.com/gabriel755/models"
	_ "github.com/gabriel755/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
