package infra

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ObterConexaoComBancoDeDados() *sql.DB {
	connectionProperties := "host=localhost user=postgres dbname=alura_loja password='123' sslmode=disable"
	conexaoComBancoDeDados, erro := sql.Open("postgres", connectionProperties)

	if erro != nil {
		log.Fatal(erro)
	}

	return conexaoComBancoDeDados
}
