package main

import (
	"fmt"
)

// Announcement struct represents the announcement details
type Announcement struct {
	Title         string
	Description   string
	Features      []string
	ContactPerson Contact
}

// Contact struct represents contact information
type Contact struct {
	Name     string
	Telegram string
}

func main() {
	// Create an instance of Announcement
	announcement := Announcement{
		Title:       "Announcing: The BigONE Exchange Price Bot for Telegram!",
		Description: "We're thrilled to unveil our latest project: the BigONE Exchange Price Bot, built in Go Lang for Telegram! This collaboration represents a significant milestone for us, marking our inaugural partnership with the BigONE Exchange backend team. Our bot empowers cryptocurrency enthusiasts with real-time market data directly within Telegram, offering a seamless experience for users.",
		Features: []string{
			"Track Cryptocurrency Prices: Stay informed with real-time prices for hundreds of cryptocurrencies, effortlessly accessible through our bot.",
			"Receive Price Alerts: Set personalized notifications for specific price targets and never miss out on trading opportunities again.",
			"User-Friendly Interface: Enjoy a clean and intuitive design, making navigation smooth and efficient for users of all experience levels.",
		},
		ContactPerson: Contact{
			Name:     "Ashikur Rahaman",
			Telegram: "https://t.me/Alex_Ashu",
		},
	}

	// Print announcement details
	fmt.Println("Title:", announcement.Title)
	fmt.Println("Description:", announcement.Description)
	fmt.Println("Features:")
	for _, feature := range announcement.Features {
		fmt.Println("- ", feature)
	}
	fmt.Println("Contact Person:")
	fmt.Println("- Name:", announcement.ContactPerson.Name)
	fmt.Println("- Telegram:", announcement.ContactPerson.Telegram)

	// Additional actions or integrations can be performed here
}
