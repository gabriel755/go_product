package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("template/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	TodosOsProdutos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", TodosOsProdutos)
}
func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

// CRIAR NOVOS PRODUTOS
func Insert(w http.ResponseWriter, r *http.Request) {
	// LER FORMULÁRIO
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConv, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço: ", err)
		}
		quantidadeConv, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na converção da quantidade: ", err)
		}

		// SALVAR NO DB
		models.CreateNewProduct(nome, descricao, precoConv, quantidadeConv)
	}
	http.Redirect(w, r, "/", 301)
}
