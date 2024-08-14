package main

import (
	"apiserver/pkg/databases"
	"apiserver/pkg/loggers"
	"apiserver/pkg/queues"
	"apiserver/pkg/validators"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	// load env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Loading env success!")

	// init logger
	loggers.NewLogger()
	fmt.Println("Init logger success!")

	// init db
	db, err := databases.NewDB()
	if err != nil {
		loggers.Logger.Fatal(err)
	}
	fmt.Println("Connect to db success!")
	defer func() {
		db.GetConn(databases.Mysql).Close()
		fmt.Println("db connection close!")
	}()

	// init Rabbit MQ
	queue, err := queues.NewQueue()
	if err != nil {
		loggers.Logger.Fatal(err)
	}
	fmt.Println("Connect to queues success!")
	defer func() {
		queue.GetRabbitMq().GetConn().Close()
		fmt.Println("rabbitmq connection close!")
	}()

	// init validator
	err = validators.NewTransValidator()
	if err != nil {
		loggers.Logger.Fatal(err)
	}
	fmt.Println("Init the validator success!")

	// start server
	server := NewGracefulServer(db, queue)
	server.RunGracefulServer()
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
