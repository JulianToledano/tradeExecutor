package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
	"time"
	db2 "tradeExecutor/db"
	"tradeExecutor/server"
)

type ticker struct {
	U        uint64  `json:"u"`
	S        string  `json:"s"`
	BidPrice float64 `json:"b"`
	BidQty   float64 `json:"B"`
	AskPrice float64 `json:"a"`
	AskQty   float64 `json:"A"`
}

func main() {
	fmt.Println("Hello")
	//websocket client connection
	//wws()
	//sqlLite()
	db, err := db2.NewSqlite("../docker/sqlite/foo.db")
	if err != nil {
		fmt.Println("ERROR %v", err)
		return
	}
	err = db2.EnsureTables(db)
	if err != nil {
		fmt.Println("ERROR %v", err)
		return
	}
	app := server.NewApp(db)
	app.Run("0.0.0.0:8003")
}

func wws() {
	c, _, err := websocket.DefaultDialer.Dial("wss://stream.binancefuture.com/ws/btcusdt@bookTicker", nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	input := make(chan ticker)

	go func() {
		// read from the websocket
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				break
			}
			// unmarshal the message
			var trade ticker
			json.Unmarshal(message, &trade)
			// send the trade to the channel
			input <- trade
		}
		close(input)
	}()

	for t := range input {
		fmt.Println(t.S)
		fmt.Println(t.BidPrice)
		fmt.Println(t.AskPrice)
		fmt.Println(t.U)
	}
}

func sqlLite() {
	db, err := sql.Open("sqlite3", "../docker/sqlite/foo.db")
	checkErr(err)

	create := `
  CREATE TABLE IF NOT EXISTS order (
  time DATETIME NOT NULL,
  symbol VARCHAR(64) NOT NULL
  );`

	if _, err := db.Exec(create); err != nil {
		checkErr(err)
	}
	// insert
	stmt, err := db.Prepare("INSERT INTO order(id, time, symbol) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(2, "2012-12-09", "BTCUSD")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	// update
	stmt, err = db.Prepare("update order set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM order")
	checkErr(err)
	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	rows.Close() //good habit to close

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
