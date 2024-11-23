package main

import (
	"encoding/xml"
	"fmt"
	"log"

	// "os"

	"validate_schema_B/models"
	"validate_schema_B/validation"
)

func main() {
	// Example JSON data
	jsonData := []byte(`{
		"name": "John Doe",
		"age": 30,
		"email": "john.doe@example.com",
		"address": {
			"street": "123 Elm Street",
			"city": "Gotham"
		}
	}`)

	// Validate JSON
	err := validation.ValidateJSON(jsonData)
	if err != nil {
		log.Fatalf("JSON validation failed: %v", err)
	}
	fmt.Println("JSON validation succeeded!")

	// Create User instance (XML example)
	user := models.User{
		Name: "Jane Smith",
		Age:  28,
		Email: "jane.smith@example.com",
		Address: &models.Address{
			Street: "456 Maple Avenue",
			City:   "Metropolis",
		},
	}

	// Convert User to XML
	xmlData, err := user.ToXML()
	if err != nil {
		log.Fatalf("Error converting user to XML: %v", err)
	}
	fmt.Println("Generated XML:")
	fmt.Println(string(xmlData))
}

// Extend User to include an XML conversion method
func (u *models.User) ToXML() ([]byte, error) {
	return xml.MarshalIndent(u, "", "  ")
}
