package services

import (
	"Homework/models"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
)

const AuthorizationOperation = `1.Авторизация,
2. Выйти
`

const LoginOperation = `Введити логин и пароль:`

func Authorization(database *sql.DB) (login, password string) {
	fmt.Println(AuthorizationOperation)
	var number int64
	_, _ = fmt.Scan(&number)
	switch number {
	case 1:
		fmt.Println(LoginOperation)
		fmt.Println("login:")
		_, _ = fmt.Scan(&login)
		fmt.Println("password:")
		_, _ = fmt.Scan(&password)
	case 2:
		fmt.Println("GOOD LUCK")
		os.Exit(0)
	case 3:
		//fmt.Println("Enter an address ")
		//var address string
		//_, _ = fmt.Scanln(&address)
		//_, err := models.AddATM(database, address)
		//if err != nil {
		//	fmt.Println("Sorry, please try again after a minute. Currently ATM does not work.")
		//}
	default:
		fmt.Println("Try again")

	}
	return
}

func Login(database *sql.DB, login, password string) {
	var User models.User
	row := database.QueryRow(`SELECT * FROM users WHERE login = ($1) AND password = ($2)`, login, password)
	err := row.Scan(
		&User.ID,
		&User.Name,
		&User.Surname,
		&User.Age,
		&User.Gender,
		&User.Login,
		&User.Password,
		&User.Access,
		&User.Remove,
	)
	if err !=nil {
		log.Fatal(err)
	}

	if User.Access {
		fmt.Println(
			`1.ADD ATM
2.EXIT'`)
		var number int
		_, _ = fmt.Scan(&number)
		switch number {
		case 1:
			address := bufio.NewReader(os.Stdin)
			fmt.Println(`enter address: `)
			text, _:= address.ReadString('\n')
			_, _ = fmt.Scan(&text)
			text2 := ""
			_, _ = fmt.Scanln(&text2)
			var ln string
			_, _ = fmt.Scanln(&ln)

			var Address string
			Address = text+` `+text2+ ` `+ln
			fmt.Println(Address)
			_, _ = models.AddATM(database, Address)
		case 2:
			fmt.Println("Goodbye")
			os.Exit(0)
		}
	}
}


