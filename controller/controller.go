package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/denis-schimidt/loja-go/dao"
)

var templateMap = template.Must(template.ParseGlob("templates/*.html"))

func Index(writer http.ResponseWriter, request *http.Request) {
	produtos, err := dao.ListarTodosOsProdutos()

	if err != nil {
		log.Fatal(err)
	}

	templateMap.ExecuteTemplate(writer, "Index", produtos)
}
