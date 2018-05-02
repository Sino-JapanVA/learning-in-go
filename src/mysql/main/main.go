package main

import (
	"database/sql"
	"fmt"

	_ "github.com/Go-SQL-Driver/MySQL"
)

func main() {
	db, err := sql.Open("mysql", "root:Sunshine@/test?charset=utf8")
	defer db.Close()

	checkErr(err)

	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("Treasure", "研发部门", "2018-5-2")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	stmt, err = db.Prepare("UPDATE userinfo SET username=? WHERE uid =?")
	checkErr(err)
	res, err = stmt.Exec("MrTreasure", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	defer stmt.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
