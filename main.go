package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/denis-schimidt/loja-go/dao"
)

var templateMap = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}

func Index(writer http.ResponseWriter, request *http.Request) {
	produtos, err := dao.ListarTodosOsProdutos()

	if err != nil {
		log.Fatal(err)
	}

	templateMap.ExecuteTemplate(writer, "Index", produtos)
}
