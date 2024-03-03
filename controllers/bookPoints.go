package controllers

import (
	/* 	"biblioteca/model"
	   	"encoding/json"
	   	"fmt"
	   	"io"
	   	"log" */
	"net/http"
	//"strings"
)

func RegisterBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Registrando livro"))

	/* requestBody, erro := io.ReadAll(r.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var Book model.Book

	if erro = json.Unmarshal(requestBody, &Book); erro != nil {
		log.Fatal(erro)
	}

	IdBook, erro := model.InsertBook(Book)

	if erro != nil {
		log.Fatal(erro)
		w.WriteHeader(500)
	}

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Usuário cadastrado com sucesso! \n ID:%d", IdBook))) */

}

func SearchAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os livros"))
}

func SearchBookName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando livro pelo nome"))

	/* 	bookName := strings.ToLower(r.URL.Query().Get("titulo"))

	   	Response := model.GetBookByName(bookName) */

}

func AlterBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Editando livro"))
}

func ExcludeBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excluindo livro"))
}
