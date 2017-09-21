/* ToDo
	sql.Open working even without stuff (es. porta)
	Variabili "inline"
*/

package godb

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	//"fmt"
	"gontheline/utils"
)

type adb struct {
	meth string
	user string
	pass string
	protocol string
	host string
	name string
	char string
}

//ToDo: take from config file
func config() adb {
	thisdb := adb{
		meth: "mysql",
		user: "root",
		pass: "",
		protocol: "tcp",
		host: "localhost:3306", //path:port or nothig
		name: "go_test",
		char: "charset=utf8"}	
	return thisdb
}

func Connect() *sql.DB {
	thisdb := config()
	db, err := sql.Open(thisdb.meth, thisdb.user+":"+thisdb.pass+"@"+thisdb.protocol+"("+thisdb.host+")/"+thisdb.name+"?"+thisdb.char)
	utils.DoPanic(err)
	return db
}

func Close(db *sql.DB) {
	db.Close()
}

/*
func main() {
	db := Connect()
	Close(db)
}
*/