package main

import (
	"godb"
)

func main() {
	var tabella = "users"
	campi := [3]string{
		"Nome",
		"Mail",
		"Password"}
	valori := [3]string{
		"Luca",
		"luca@lambia.com",
		"mammt"}
	valori2 := [3]string{
		"Lambi",
		"luca@lambia.it",
		"mammt2"}

		/*
	valoriNew := [2][3]string{
		{"Nome", "Mail", "Password"} ,
		{"Tizio", "me@g.it", "sga42x"}}
		*/

	db := godb.Connect()
	id := godb.Insert(db,tabella,campi[:],valori[:])
	godb.Update(db,tabella,campi,valori2,id)
	godb.QuerySelect(db,"",tabella,"ID=28")
	//godb.Query(db,"SELECT * FROM"+tabella+" WHERE 1=1")
	godb.Delete(db,tabella,id)
	godb.Close(db)

}