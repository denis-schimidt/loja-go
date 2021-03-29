package route

import (
	"net/http"

	"github.com/denis-schimidt/loja-go/controller"
)

func CarregaRotas() {
	http.HandleFunc("/", controller.Index)
}
