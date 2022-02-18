package main

import (
	"net/http"

	"github.com/hygorm10/aluraLoja/routes"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
