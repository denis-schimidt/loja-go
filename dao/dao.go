package dao

import (
	"github.com/denis-schimidt/loja-go/infra"
	"github.com/denis-schimidt/loja-go/model"
)

func ListarTodosOsProdutos() ([]model.Produto, error) {
	conexaoComBancoDeDados := infra.ObterConexaoComBancoDeDados()
	registrosDeProdutos, erro := conexaoComBancoDeDados.Query("SELECT * from produtos")

	if erro != nil {
		return nil, erro
	}

	var listaProdutos []model.Produto

	for registrosDeProdutos.Next() {
		produto := model.Produto{}
		erro := registrosDeProdutos.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)

		if erro != nil {
			return nil, erro
		}

		listaProdutos = append(listaProdutos, produto)
	}

	defer conexaoComBancoDeDados.Close()

	return listaProdutos, nil
}
