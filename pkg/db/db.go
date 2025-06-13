package db

import (
	"database/sql"
	"fmt"
	"log"

	"restAPI/pkg/mocks"

	"github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "nixon"
	password = "password"
	dbname   = "cardifyDB"
)

func Connect() *sql.DB {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to db!")
	return db
}

func CloseConnection(db *sql.DB) {
	defer db.Close()
}

func CreateTable(db *sql.DB) {
	var exists bool
	if err := db.QueryRow("SELECT EXISTS (SELECT FROM pg_tables WHERE  schemaname = 'public' AND tablename = 'users' );").Scan(&exists); err != nil {
		fmt.Println("failed to execute query", err)
		return
	}
	if !exists {
		results, err := db.Query("CREATE TABLE users (uid VARCHAR(36) PRIMARY KEY, username VARCHAR(100) NOT NULL, email VARCHAR(50) NOT NULL, picture TEXT, following TEXT[] NOT NULL, friends TEXT[] NOT NULL);")
		if err != nil {
			fmt.Println("failed to execute query", err)
			return
		}
		fmt.Println("Table created successfully", results)

		for _, user := range mocks.User {
			queryStmt := `INSERT INTO users (uid,username,email,picture,following,friends) VALUES ($1, $2, $3, $4, $5, $6) RETURNING uid;`

			err := db.QueryRow(queryStmt, &user.Uid, &user.Username, &user.Email, &user.Picture, pq.Array(user.Following), pq.Array(user.Friends)).Scan(&user.Uid)
			if err != nil {
				log.Println("failed to execute query", err)
				return
			}
		}
		fmt.Println("Mock Users included in Table", results)
	} else {
		fmt.Println("Table 'users' already exists ")
	}

}

func DeleteTable(db *sql.DB) {
	var exists bool
	if err := db.QueryRow("DROP TABLE users;").Scan(&exists); err != nil {
		fmt.Println("failed to execute query", err)
		return
	}
	fmt.Println("Table users deleted.")
}
