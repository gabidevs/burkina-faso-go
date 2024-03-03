package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InsertBook(livro Book) (int64, error) {

	var Ibook Book

	Ibook = livro

	db, erro := concect()

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO livros (titulo, editora, imagem, autor) VALUES (?, ?, ?, ?)")

	if erro != nil {

		return 0, erro
	}

	defer statement.Close()

	insert, erro := statement.Exec(Ibook.Titulo, Ibook.Editora, Ibook.Imagem, Ibook.Autor)

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

func PutBook(livro Book, id int) error {

	var Ibook Book

	Ibook = livro

	db, erro := concect()

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	statement, erro := db.Prepare("UPDATE livros SET titulo= ?, editora = ?, imagem= ?, autor =  ? WHERE id = ?")

	statement.Exec(Ibook.Titulo, Ibook.Editora, Ibook.Imagem, Ibook.Autor, id)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	return erro

}

func GetAllBooks() []Book {

	db, erro := concect()

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	rows, erro := db.Query("SELECT * FROM livros")

	var response []Book

	for rows.Next() {
		var structBook Book

		erro := rows.Scan(&structBook.ID, &structBook.Titulo, &structBook.Editora, &structBook.Imagem, &structBook.Autor)

		if erro != nil {
			panic(erro)
		}

		response = append(response, structBook)

	}

	return response

}

func GetBookByName(name string) Book {
	db, erro := concect()

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	row, erro := db.Query("SELECT * FROM livros WHERE titulo LIKE '% %s %'", name)

	if erro != nil {
		log.Fatal(erro)
	}

	var bookName Book

	if row.Next() {

		if erro := row.Scan(&bookName.ID, &bookName.Titulo, &bookName.Editora, &bookName.Imagem, &bookName.Autor); erro != nil {
			log.Fatal(erro)
		}
	}

	return bookName
}

func DeleteBook(id int) error {

	db, erro := concect()

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM livros WHERE id = ?")

	if erro != nil {
		log.Fatal(erro)
	}

	if _, erro := statement.Exec(id); erro != nil {
		log.Fatal(erro)
	}

	return erro

}
