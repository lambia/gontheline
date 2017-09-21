/* todo
	insert, update ecc.. return id not print
	-> sono methodi del db
	-> valori e campi = unico struct

	query può accettare query specifica o costruirla da parametri
	-> restituisce struct coi dati
	-> argomenti da usare come filtri (per query e per echo successivo)

	delete accetta struct {ID: 4}, restituisce bool

	sostituire nomi funzioni troppo generici
	mettere tutte queste funzioni in un package unico
*/

package main

import (
	//"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"gontheline/utils"
	"gontheline/godb"
)

func InsertUser(db *sql.DB, user map[string]string ) int64 {
	stmt, err := db.Prepare("INSERT users SET Nome=?,Mail=?,Password=?")
	utils.DoPanic(err)

	res, err := stmt.Exec(user["user"], user["mail"], user["pass"])
	utils.DoPanic(err)

	id, err := res.LastInsertId()
	utils.DoPanic(err)
	
	return id
}

func UpdateUser(db *sql.DB, what string, value string, id int64) int64 {
	stmt, err := db.Prepare("update users set ?=? where ID=?")
	utils.DoPanic(err)

	res, err := stmt.Exec(what, value, id)
	utils.DoPanic(err)

	affect, err := res.RowsAffected()
	utils.DoPanic(err)

	return affect
}

func QuerySelectUser(db *sql.DB, what string, where string) {
	var queryString string
	if where == "" {
		queryString = "SELECT "+what+" FROM users"
	} else {
		queryString = "SELECT "+what+" FROM users WHERE "+where		
	}
	
	rows, err := db.Query(queryString)
	utils.DoPanic(err)

	for rows.Next() {
		var uid int
		var name string
		var mail string
		var pass string
		err = rows.Scan(&uid, &name, &mail, &pass)
		utils.DoPanic(err)
		fmt.Println(uid)
	}
}

func DeleteUser(db *sql.DB, id int64) {
	stmt, err := db.Prepare("delete from users where ID=?")
	utils.DoPanic(err)

	res, err := stmt.Exec(id)
	utils.DoPanic(err)

	affect, err := res.RowsAffected()
	utils.DoPanic(err)

	fmt.Print(affect)
	fmt.Print(" row(s) eliminated")
}

func main() {
	db := godb.Connect()
	
	userData := map[string]string{"user": "Luca", "mail": "luca@lambia.it", "pass": "pass]#[word"}
	id := InsertUser(db, userData )
	println(id)
	godb.Close(db)
}