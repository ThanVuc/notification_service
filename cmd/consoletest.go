package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"notification_service/internal/application/dto"
	"notification_service/internal/infrastructure"
	interface_constant "notification_service/internal/interface/constant"
	"os"
	"strings"

	"firebase.google.com/go/v4/messaging"
	"github.com/google/uuid"
	"github.com/wagslane/go-rabbitmq"
)

func RunConsoleTest(
	ctx context.Context,
	infrastructure *infrastructure.InfrastructureModule,
) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Console Test Menu ===")
		fmt.Println("1) Test Firebase push")
		fmt.Println("2) Test Team consumer via RabbitMQ (supports direct email worker case)")
		fmt.Println("3) Exit")
		fmt.Print("Choose option: ")

		choice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading option:", err)
			return
		}

		switch strings.TrimSpace(choice) {
		case "1":
			runFirebaseConsoleTest(ctx, infrastructure, reader)
		case "2":
			runTeamNotificationRabbitConsoleTest(ctx, reader)
		case "3", "exit":
			fmt.Println("Bye")
			return
		default:
			fmt.Println("Unknown option")
		}
	}
}

func runFirebaseConsoleTest(
	ctx context.Context,
	infrastructure *infrastructure.InfrastructureModule,
	reader *bufio.Reader,
) {
	var name string
	client, err := infrastructure.BaseModule.FirebaseApp.Messaging(ctx)
	if err != nil {
		fmt.Println("Error getting Messaging client:", err)
		return
	}

	for {
		fmt.Print("Enter your name: ")
		line, readErr := reader.ReadString('\n')
		if readErr != nil {
			fmt.Println("Error reading input:", readErr)
			return
		}

		name = strings.TrimSpace(line)
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

		if name == "exit" {
			break
		}
	}
}

func runTeamNotificationRabbitConsoleTest(ctx context.Context, reader *bufio.Reader) {
	fmt.Print("RabbitMQ URI (default amqp://guest:guest@localhost:5672): ")
	rawURI, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading RabbitMQ URI:", err)
		return
	}

	rabbitURI := strings.TrimSpace(rawURI)
	if rabbitURI == "" {
		rabbitURI = "amqp://guest:guest@localhost:5672"
	}

	conn, err := rabbitmq.NewConn(rabbitURI, rabbitmq.WithConnectionOptionsLogging)
	if err != nil {
		fmt.Println("Error creating RabbitMQ connection:", err)
		return
	}
	defer conn.Close()

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName(string(interface_constant.TEAM_EXCHANGE)),
		rabbitmq.WithPublisherOptionsExchangeKind("direct"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
		rabbitmq.WithPublisherOptionsExchangeDurable,
	)
	if err != nil {
		fmt.Println("Error creating RabbitMQ publisher:", err)
		return
	}
	defer publisher.Close()

	fmt.Println("RabbitMQ ready")
	fmt.Println("Exchange:", interface_constant.TEAM_EXCHANGE)
	fmt.Println("Routing key:", interface_constant.TEAM_ROUTING_KEY)

	for {
		fmt.Println("\n--- Team Message Publisher ---")
		title := readLine(reader, "Title (or 'exit' to stop): ")
		if title == "exit" {
			return
		}

		message := readLine(reader, "Message: ")
		linkInput := readLine(reader, "Link (optional): ")
		imageInput := readLine(reader, "Image URL (optional): ")
		senderID := readLine(reader, "Sender ID: ")
		receiverIDsInput := readLine(reader, "Receiver IDs CSV (leave empty to test direct email worker): ")
		nonExistentReceiversInput := readLine(reader, "Non-existent receiver emails CSV: ")
		isSentMailInput := readLine(reader, "is_sent_mail (true/false, default true): ")

		receiverIDs := splitCSV(receiverIDsInput)
		nonExistentReceivers := splitCSV(nonExistentReceiversInput)
		isSentMail := parseBoolWithDefault(isSentMailInput, true)

		var link *string
		if linkInput != "" {
			link = &linkInput
		}

		var imageURL *string
		if imageInput != "" {
			imageURL = &imageInput
		}

		teamMessage := dto.TeamNotificationMessage{
			EventType:   "team_notification",
			SenderID:    senderID,
			ReceiverIDs: receiverIDs,
			Payload: dto.TeamNotificationMessagePayload{
				Title:           title,
				Message:         message,
				Link:            link,
				ImageURL:        imageURL,
				CorrelationID:   uuid.NewString(),
				CorrelationType: 0,
			},
			Metadata: dto.TeamNotificationMessageMetadata{
				IsSentMail:           isSentMail,
				NonExistentReceivers: nonExistentReceivers,
			},
		}

		payload, err := json.Marshal(teamMessage)
		if err != nil {
			fmt.Println("Error marshaling message:", err)
			continue
		}

		err = publisher.PublishWithContext(
			ctx,
			payload,
			[]string{interface_constant.TEAM_ROUTING_KEY},
			rabbitmq.WithPublishOptionsExchange(string(interface_constant.TEAM_EXCHANGE)),
			rabbitmq.WithPublishOptionsContentType("application/json"),
		)
		if err != nil {
			fmt.Println("Publish failed:", err)
			continue
		}

		fmt.Println("Message published successfully")
		fmt.Println("- exchange:", interface_constant.TEAM_EXCHANGE)
		fmt.Println("- routing key:", interface_constant.TEAM_ROUTING_KEY)
		fmt.Println("- receiver_ids length:", len(receiverIDs))
		fmt.Println("- non_existent_receivers length:", len(nonExistentReceivers))
	}
}

func readLine(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	line, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(line)
}

func splitCSV(value string) []string {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return []string{}
	}

	parts := strings.Split(trimmed, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		item := strings.TrimSpace(part)
		if item != "" {
			result = append(result, item)
		}
	}

	return result
}

func parseBoolWithDefault(value string, defaultValue bool) bool {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "", "default":
		return defaultValue
	case "true", "t", "1", "yes", "y":
		return true
	case "false", "f", "0", "no", "n":
		return false
	default:
		return defaultValue
	}
}
