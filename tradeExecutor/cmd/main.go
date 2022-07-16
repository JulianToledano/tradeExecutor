package main

import (
	"flag"
	"fmt"
	//_ "github.com/mattn/go-sqlite3"
	"os"
	"tradeExecutor/config"
	"tradeExecutor/db"
	"tradeExecutor/engine"
	"tradeExecutor/logger"
	"tradeExecutor/order"
	"tradeExecutor/server"
)

func main() {
	configFile := flag.String("config", "config.toml", "configuration")
	flag.Parse()

	c, err := config.ReadConfig(*configFile)
	fmt.Println(c.LogFile)
	logFile, err := os.OpenFile(c.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatalf("could not set logFile %s", c.LogFile)
	}
	logger.Set(logFile, c.LogLevel)
	logger.Infof("[MAIN] [START]")

	dataBase, err := db.NewSqlite(c.Db)
	if err != nil {
		logger.Fatalf("[MAIN] [db] [NewSqlite] [%v]", err)
	}
	err = dataBase.EnsureTables()
	if err != nil {
		logger.Fatalf("[MAIN] [db] [EnsureTables] [%v]", err)
	}

	oChan := make(chan *order.Order)
	e := engine.NewTradeExecutor(oChan)
	go e.Execute()

	app := server.NewApp(dataBase)
	app.Run(c.ServerAddr, oChan)
}
