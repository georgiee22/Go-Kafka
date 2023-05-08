package main

import (
	routers "Template/pkg/routers"
	middleware "Template/pkg/utils"
	"Template/pkg/utils/go-utils/database"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"github.com/segmentio/kafka-go"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading Env File: ", err)
	}
	envi := os.Getenv("ENVIRONMENT")

	err = godotenv.Load(fmt.Sprintf(".env-%v", envi)) //
	if err != nil {
		log.Fatal("Error Loading Env File: ", err)
	}

	// Initialize DB here
	database.PostgreSQLConnect(
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSL_MODE"),
		os.Getenv("POSTGRES_TIMEZONE"),
	)

	database.MySQLConnect(
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DB_NAME"),
	)

	// Declare & initialize fiber
	app := fiber.New(fiber.Config{
		UnescapePath: true,
	})

	// For GoRoutine implementation
	// appb := fiber.New(fiber.Config{
	// 	UnescapePath: true,
	// })

	// Configure application CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// For GoRoutine implementation
	// appb.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	// }))

	// Declare & initialize logger
	app.Use(logger.New())

	// Create a new listener using the pq library
	listener := pq.NewListener("user=admin password=password dbname=postgres host=localhost sslmode=disable", 10*time.Second, time.Minute, func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Println(err.Error())
		}
	})
	defer listener.Close()

	// Listen for a specific channel
	err = listener.Listen("my_channel")
	if err != nil {
		log.Fatal(err)
	}

	// ticker := time.NewTicker(1 * time.Second)
	// quit := make(chan struct{})
	// // go routine to insert into table
	// go func() {
	// 	i := 0
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			data1 := fmt.Sprintf("admin%v", i)
	// 			data2 := fmt.Sprintf("admin%v", i)
	// 			// Insert new user
	// 			err = database.DBConn.Exec("INSERT INTO test_account_infos (username, password) VALUES ($1, $2)", data1, data2).Error
	// 			if err != nil {
	// 				panic(err)
	// 			}
	// 			i++
	// 		case <-quit:
	// 			close(quit)
	// 		}
	// 	}
	// }()

	// go routine to listen to inserts
	//For GoRoutine implementation
	go func() {
		// Wait for notifications
		for {
			select {
			case n := <-listener.Notify:
				fmt.Println("Received notification:", n.Extra)

				// connect to kafka broker
				conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "topic1", 0)
				// set timer to stop trying to send message after 10 seconds
				conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
				// write the message to be put in the specified topic
				conn.WriteMessages(kafka.Message{Value: []byte(n.Extra)})
				log.Println(n.Extra)
			case <-time.After(10 * time.Minute):
				fmt.Println("Timeout")
				os.Exit(0)
			}
		}
	}()

	// Declare & initialize routes
	routers.SetupPublicRoutes(app)
	routers.SetupPrivateRoutes(app)

	// For GoRoutine implementation
	// routers.SetupPublicRoutesB(appb)
	// go func() {
	// 	err := appb.Listen(fmt.Sprintf(":8002"))
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// }()

	fmt.Println("Port: ", middleware.GetEnv("PORT"))
	// Serve the application
	if middleware.GetEnv("SSL") == "enabled" {
		log.Fatal(app.ListenTLS(
			fmt.Sprintf(":%s", middleware.GetEnv("PORT")),
			middleware.GetEnv("SSL_CERTIFICATE"),
			middleware.GetEnv("SSL_KEY"),
		))
	} else {
		err := app.Listen(fmt.Sprintf(":%s", middleware.GetEnv("PORT")))
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
