package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InsertUser(user Usuario) (int64, error) {

	var Iuser Usuario

	Iuser = user

	db, erro := concect()

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO usuarios (nome, email, senha) VALUES (?, ?, ?)")

	if erro != nil {

		return 0, erro
	}

	defer statement.Close()

	insert, erro := statement.Exec(Iuser.Nome, Iuser.Email, Iuser.Senha)

	if erro != nil {

		return 0, erro
	}

	defer statement.Close()

	idInsert, erro := insert.LastInsertId()

	if erro != nil {

		return 0, erro
	}

	return idInsert, nil

}

func PutUser(user Usuario, id int) error {

	var Iuser Usuario

	Iuser = user

	db, erro := concect()

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	statement, erro := db.Prepare("UPDATE usuarios SET nome= ?, email = ?, senha= ? WHERE id = ?;")

	statement.Exec(Iuser.Nome, Iuser.Email, Iuser.Senha, id)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	return erro

}

func GetAllUsers() []Usuario {

	db, erro := concect()

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	rows, erro := db.Query("SELECT * FROM usuarios")

	var response []Usuario

	for rows.Next() {
		var structUser Usuario

		erro := rows.Scan(&structUser.ID, &structUser.Nome, &structUser.Email, &structUser.Senha)

		if erro != nil {
			panic(erro)
		}

		response = append(response, structUser)

	}

	return response

}

func GetUserId(id int) Usuario {

	db, erro := concect()

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	row, erro := db.Query("SELECT * FROM usuarios WHERE id = ?", id)

	if erro != nil {
		log.Fatal(erro)
	}

	var userId Usuario

	if row.Next() {

		if erro := row.Scan(&userId.ID, &userId.Nome, &userId.Email, &userId.Senha); erro != nil {
			log.Fatal(erro)
		}
	}

	return userId

}

func DeleteUser(id int) error {

	db, erro := concect()

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM usuarios WHERE id = ?")

	if erro != nil {
		log.Fatal(erro)
	}

	if _, erro := statement.Exec(id); erro != nil {
		log.Fatal(erro)
	}

	return erro
}
