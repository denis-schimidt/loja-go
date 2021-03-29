package main

import (
	"net/http"

	"github.com/denis-schimidt/loja-go/route"
)

func main() {
	http.ListenAndServe(":8000", nil)
	route.CarregaRotas()
}
