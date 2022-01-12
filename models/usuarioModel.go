package models

import (
	"github.com/pk-anderson/exercicio-crud-go/database"
)

type User struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Fone  string `json:"fone"`
}

func SelectAllUsers() []User {
	db := database.StartDatabase()
	query := "SELECT * FROM usuario"
	response, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var usersList []User = make([]User, 0)
	for response.Next() {
		var user User
		scanErr := response.Scan(&user.Id, &user.Nome, &user.Email, &user.Fone)
		if scanErr != nil {
			panic(scanErr.Error())
		}
		usersList = append(usersList, user)
	}

	defer db.Close()
	return usersList
}

func SelectUserById(id int) User {
	db := database.StartDatabase()

	query := "SELECT * FROM usuario WHERE id = ?"
	response := db.QueryRow(query, id)
	var user User
	scanErr := response.Scan(&user.Id, &user.Nome, &user.Email, &user.Fone)
	if scanErr != nil {
		panic(scanErr.Error())
	}

	defer db.Close()
	return user
}

func CreateNewUser(nome, email, fone string) int {
	db := database.StartDatabase()

	query := "INSERT INTO usuario (nome, email, fone) VALUES (?, ?, ?)"
	result, err := db.Exec(query, nome, email, fone)
	if err != nil {
		panic(err.Error())
	}

	newUserId, newUserErr := result.LastInsertId()
	if newUserErr != nil {
		panic(newUserErr.Error())
	}

	defer db.Close()
	return int(newUserId)
}

func UpdateUser(id int, nome, email, fone string) {
	db := database.StartDatabase()
	queryFind := "SELECT id FROM usuario WHERE id = ?"
	result := db.QueryRow(queryFind, id)
	var idUser int
	scanErr := result.Scan(&idUser)
	if scanErr != nil {
		panic(scanErr.Error())
	}

	queryUpdate := "UPDATE usuario SET nome = ?, email = ?, fone = ? WHERE id = ?"
	_, updateErr := db.Exec(queryUpdate, nome, email, fone, id)
	if updateErr != nil {
		panic(updateErr.Error())
	}
	defer db.Close()
}

func DeleteUser(id int) {
	db := database.StartDatabase()
	queryFind := "SELECT id FROM usuario WHERE id = ?"
	result := db.QueryRow(queryFind, id)
	var idUser int
	scanErr := result.Scan(&idUser)
	if scanErr != nil {
		panic(scanErr.Error())
	}

	queryDelete := "DELETE FROM usuario WHERE id = ?"
	_, deleteErr := db.Exec(queryDelete, id)
	if deleteErr != nil {
		panic(deleteErr.Error())
	}
	defer db.Close()
}
