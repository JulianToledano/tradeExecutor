package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqlite(dbFile string) (db *sql.DB, err error) {
	return sql.Open("sqlite3", dbFile)
	//checkErr(err)
	//
	//create := `
	//CREATE TABLE IF NOT EXISTS order (
	//time DATETIME NOT NULL,
	//symbol VARCHAR(64) NOT NULL
	//);`
	//
	//if _, err := db.Exec(create); err != nil {
	//checkErr(err)
	//}
	//// insert
	//stmt, err := db.Prepare("INSERT INTO order(id, time, symbol) values(?,?,?)")
	//checkErr(err)
	//
	//res, err := stmt.Exec(2, "2012-12-09", "BTCUSD")
	//checkErr(err)
	//
	//id, err := res.LastInsertId()
	//checkErr(err)
	//
	//fmt.Println(id)
	//// update
	//stmt, err = db.Prepare("update order set username=? where uid=?")
	//checkErr(err)
	//
	//res, err = stmt.Exec("astaxieupdate", id)
	//checkErr(err)
	//
	//affect, err := res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println(affect)
	//
	//// query
	//rows, err := db.Query("SELECT * FROM order")
	//checkErr(err)
	//var uid int
	//var username string
	//var department string
	//var created time.Time
	//
	//for rows.Next() {
	//err = rows.Scan(&uid, &username, &department, &created)
	//checkErr(err)
	//fmt.Println(uid)
	//fmt.Println(username)
	//fmt.Println(department)
	//fmt.Println(created)
	//}
	//
	//rows.Close() //good habit to close
	//
	//// delete
	//stmt, err = db.Prepare("delete from userinfo where uid=?")
	//checkErr(err)
	//
	//res, err = stmt.Exec(id)
	//checkErr(err)
	//
	//affect, err = res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println(affect)
	//
	//db.Close()

}

func EnsureTables(db *sql.DB) (err error) {
	create := `
	CREATE TABLE IF NOT EXISTS orders(
	size REAL NOT NULL,
	price REAL NOT NULL,
	symbol VARCHAR(64) NOT NULL,
	buy INTEGER NOT NULL,
	date DATE NOT NULL
	);`
	_, err = db.Exec(create)
	return
}
