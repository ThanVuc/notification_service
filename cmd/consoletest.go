package cmd

import (
	"context"
	"fmt"
	"notification_service/internal/infrastructure"

	"firebase.google.com/go/v4/messaging"
)

func RunConsoleTest(
	ctx context.Context,
	infrastructure *infrastructure.InfrastructureModule,
) {
	var name string
	client, err := infrastructure.BaseModule.FirebaseApp.Messaging(ctx)
	if err != nil {
		fmt.Println("Error getting Messaging client:", err)
		return
	}

	for {
		fmt.Print("Enter your name: ")
		fmt.Scanln(&name)
		fmt.Println("Send: ", name)

		msg := &messaging.Message{
			Token: "fHdvpIz8wTzx4s5q2rxZ2j:APA91bF4KE2NJhYNfZhAikoy9WAE8DOQ1Tj52yd2bTPQxb52VICFIDQXc5NJhzwK4SA-dtYB_O8R-IBr6cwd6avA6B-jpgGQmnTUtBy-NAp1WzPuwf8X1Ng",
			Data: map[string]string{
				"title": "Hello " + name,
				"body":  "This is a test notification",
				"url":   "https://localhost:3000/schedule/daily",
				"src":   "https://jbagy.me/wp-content/uploads/2025/03/Hinh-anh-avatar-anime-nu-cute-1.jpg",
			},
		}

		response, err := client.Send(ctx, msg)
		if err != nil {
			fmt.Println("Error sending message:", err)
		} else {
			fmt.Println("Successfully sent message:", response)
		}

		if err != nil {
			fmt.Println("Error sending message:", err)
		}

		if name == "exit" {
			break
		}
	}
}
