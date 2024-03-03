package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Usuario struct {
	ID    int    `json:"id,omitempty"`
	Nome  string `json:"nome,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
}

type Book struct {
	ID      int    `json:"id,omitempty"`
	Titulo  string `json:"titulo,omitempty"`
	Editora string `json:"editora,omitempty"`
	Imagem  string `json:"imagem,omitempty"`
	Autor   string `json:"autor,omitempty"`
}

// Pacote de conexão com o banco de dados
func concect() (*sql.DB, error) {

	stringConections := "burkina:burkina@/burkina_faso"

	db, erro := sql.Open("mysql", stringConections)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil

}
