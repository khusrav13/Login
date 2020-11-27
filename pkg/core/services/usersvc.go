package services

import (
	"Homework/models"
	"database/sql"
	"fmt"
)

const AuthorizationOperation = `1.Авторизация,
2. Выйти`

const LoginOperation = `Введити логин и пароль:`

func Authorization() (login, password string) {
	fmt.Println(AuthorizationOperation)
	var number int64
	fmt.Scan(&number)
	switch number {
	case 1:
		fmt.Println(LoginOperation)
		fmt.Println("login:")
		fmt.Scan(&login)
		fmt.Println("password:")
		fmt.Scan(&password)
	case 2:
		fmt.Println("GOOD LUCK")
	default:
		fmt.Println("Try again")

	}
	return
}

func Login(database *sql.DB, login, password string) { //(ok bool)
	var User models.User
	_ = database.QueryRow(`SELECT * FROM users WHERE login = ($1) AND password = ($2)`, login, password).Scan(
		&User.ID,
		&User.Name,
		&User.Surname,
		&User.Age,
		&User.Gender,
		&User.Login,
		&User.Password,
		&User.Remove,
	)

	//if User.ID > 0 {
	//	return true
	//}
	//return false
	fmt.Println(User)

}
