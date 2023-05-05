package routes

import (
	"net/http"

	_ "github.com/gabriel755/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", index)
	http.HandleFunc("/new", controllers.NewProduct)
	http.HandleFunc("/insert", controllers.Insert)
}
