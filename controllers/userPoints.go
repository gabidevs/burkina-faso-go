package controllers

import (
	"biblioteca/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var Usuario model.Usuario

	if erro = json.Unmarshal(requestBody, &Usuario); erro != nil {
		log.Fatal(erro)
	}

	IdUser, erro := model.InsertUser(Usuario)

	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Usuário cadastrado com sucesso! \n ID:%d", IdUser)))
}

func SearchAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários"))
}

func SearchUserID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário"))
}

func AlterUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Editar usuário"))
}

func ExcludeUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excluir usuário"))
}
