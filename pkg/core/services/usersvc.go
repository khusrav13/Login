package services

import (
	"Homework/models"
	"database/sql"
	"fmt"
)

const AuthorizationOperation = `1.Авторизация,
2. Выйти`

const LoginOperation = `Введити логин и пароль:`

func Authorization() (Login, Password string) {
	fmt.Println(AuthorizationOperation)
	var number int64
	fmt.Scan(&number)
	switch number {
	case 1:
		fmt.Println(LoginOperation)
		fmt.Println("login:")
		fmt.Scan(&Login)
		fmt.Println("password:")
		fmt.Scan(&Password)
		return Password, Login
	case 2:
		fmt.Println("GOOD LUCK")
	default:
		fmt.Println("Try again")

	}
	return
}

func Login(database *sql.DB, Login, Password string) { //(ok bool)
	var User models.User
	_ = database.QueryRow(`SELECT * FROM USERS WHERE Login = ($1) AND password = ($2)`, Login, Password).Scan(
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
