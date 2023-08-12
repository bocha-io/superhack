package handlers

import (
	"database/sql"
	"fmt"

	"github.com/bocha-io/txbuilder/x/txbuilder"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath string = "users.sql"

func OpenDatabase(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		fmt.Printf("Error creating/opening database: %q", err)
		panic("Stop processing")
	}
	return db
}

func InitDatabase(builder *txbuilder.TxBuilder) {
	_ = builder
	db := OpenDatabase(dbPath)
	sqlStmt := `
       create table if not exists users(
        id integer not null primary key,
        username text unique,
        password text,
        address text,
        walletindex integer
    );
       `

	_, err := db.Exec(sqlStmt)
	if err != nil {
		fmt.Println(err.Error())
		panic("Could not init database")
	}

	sqlStmt = `
       create table if not exists faucet(
        id integer not null primary key,
        amount integer
    );
       `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Println(err.Error())
		panic("Could not init database")
	}

	rows, err := db.Query("select count(*) from users")
	if err != nil {
		fmt.Println("Error reading from users", err)
		panic("error reading db")
	}
	var usersCount int
	for rows.Next() {
		err := rows.Scan(&usersCount)
		if err != nil {
			panic("error reading row")
		}
	}
	rows.Close()

	rows, err = db.Query("select amount from faucet")
	if err != nil {
		fmt.Println("Error reading from users", err)
		panic("error reading db")
	}
	var faucet int
	for rows.Next() {
		err := rows.Scan(&faucet)
		if err != nil {
			panic("error reading row")
		}
	}
	rows.Close()

	// end := usersCount + 20
	// for faucet <= end {
	// 	_, errFaucet := builder.FundAnAccount(faucet)
	// 	logger.LogInfo(fmt.Sprintf("[backend] sending coins to wallet: %d", faucet))
	// 	if errFaucet != nil {
	// 		logger.LogError(
	// 			fmt.Sprintf(
	// 				"[backend] error sending coins to wallet %d, %s",
	// 				faucet,
	// 				errFaucet.Error(),
	// 			),
	// 		)
	// 	}
	// 	faucet++
	// }

	tx, _ := db.Begin()
	updateFaucet, err := tx.Prepare("INSERT OR REPLACE INTO faucet(id, amount) VALUES (?,?)")
	if err != nil {
		fmt.Printf("Error preparing transaction: %q", err)
		panic("Could not update faucet")
	}
	_, _ = updateFaucet.Exec(0, faucet-1)
	_ = tx.Commit()
	db.Close()
}

func InsertUser(username string, password []byte, address string, index int) {
	db := OpenDatabase(dbPath)
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("Error creating transaction: %q", err)
		panic("Stop proceeding")
	}
	insertAccount, err := tx.Prepare(
		"insert into users(username,password,address,walletindex) values(?,?,?,?)",
	)
	if err != nil {
		fmt.Printf("Error preparing transaction: %q", err)
		panic("Could not create insert user")
	}
	_, _ = insertAccount.Exec(username, string(password), address, index)
	_ = tx.Commit()
	db.Close()
}

func InitUsersMap() *map[string]User {
	temp := map[string]User{}
	db := OpenDatabase(dbPath)

	rows, err := db.Query(
		"select username,password,address,walletindex from users order by walletindex",
	)
	if err != nil {
		fmt.Println("Error reading addresses", err)
		panic("error reading db")
	}
	for rows.Next() {
		var username string
		var password string
		var address string
		var index int
		err := rows.Scan(&username, &password, &address, &index)
		if err != nil {
			panic("error reading row")
		}
		temp[username] = User{
			Username: username,
			Password: []byte(password),
			Address:  address,
			Index:    index,
		}
	}
	rows.Close()
	db.Close()
	return &temp
}
