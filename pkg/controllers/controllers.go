package controllers

import (
	"Template/pkg/kafka"
	"Template/pkg/models"
	"Template/pkg/models/response"
	"Template/pkg/utils/go-utils/database"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func XmltoJson(c *fiber.Ctx) error {
	// Example XML data
	xmlData := `
        <person>
            <name>John</name>
            <age>30</age>
            <email>john@example.com</email>
			<sex>Male</sex>
        </person>
    `

	// Unmarshal XML data into Go struct
	var person struct {
		Name  string `xml:"name"`
		Age   int    `xml:"age"`
		Email string `xml:"email"`
	}

	err := xml.Unmarshal([]byte(xmlData), &person)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Marshal Go struct into JSON data
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Print JSON data
	fmt.Println(string(jsonData))

	return c.SendString(string(jsonData))
}

func MySQL_Read(c *fiber.Ctx) error {
	account_model := []models.Account_Infos{}
	err := database.DBConn2.Debug().Raw("SELECT * FROM test_account_infos").Scan(&account_model).Error
	if err != nil {
		return err
	}

	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    account_model,
	})
}

func Consumer_Kafka(c *fiber.Ctx) error {
	// // connect to kafka broker
	// conn, _ := kg.DialLeader(context.Background(), "tcp", "localhost:9092", "topic1", 0)
	// // set timer to stop trying to send message after 8 seconds
	// conn.SetReadDeadline(time.Now().Add(time.Second * 8))

	// // message, _ := conn.ReadMessage(1e6)
	// // fmt.Println(string(message.Value))

	// // batch_message := conn.ReadBatch(1e3, 1e6)
	// // bytes := make([]byte, 1e3)
	// // for {
	// // 	_, err := batch_message.Read(bytes)
	// // 	if err != nil {
	// // 		break
	// // 	}
	// // 	fmt.Println(string(bytes))
	// // }
	// batch_message := conn.ReadBatch(1e3, 1e6)
	// var messages []string
	// bytes := make([]byte, 1e3)
	// for {
	// 	_, err := batch_message.Read(bytes)
	// 	if err != nil {
	// 		break
	// 	}
	// 	fmt.Println(string(bytes))
	// }

	kafka.InitializeNewReader("localhost:9092", "topic1", "test-group")

	for {
		message, err := kafka.Reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			break
		}
		println(string(message.Value))
	}

	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    "",
	})
}
