// package main
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/xeipuuv/gojsonschema"
)

// User represents the user data
type User struct {
	Name  string `json:"name" xml:"name"`
	Age   int    `json:"age" xml:"age"`
	Email string `json:"email" xml:"email"`
}

// validateUser validates a User struct using go-playground/validator
func validateUser(user *User) error {
	validate := validator.New()
	err := validate.Struct(user)
	return err
}
// user_schema.json
{
    "$schema" : "http://json-schema.org/draft-07/schema#",
    "type" : "object",
    "properties": {
        "name": {
            "type": "string",
            "pattern": "^[a-zA-Z ]+$",
            "minLength": 3,
            "maxLength": 100
        },
        "age": {
            "type": "integer",
            "minimum": 0,
            "maximum": 150
        },
        "email": {
            "type": "string",
            "format": "email"
        }
    },
    "required": ["name", "age", "email"]
}
func validateJSON(jsonData []byte) error {
	// Load JSON schema
	schemaLoader := gojsonschema.NewReferenceLoader(gojsonschema.MustAsset("user_schema.json"))
	schema, err := schemaLoader.LoadSchema("")
	if err != nil {
		return fmt.Errorf("error loading JSON schema: %w", err)
	}

	// Validate JSON data against schema
	document := gojsonschema.NewDocument(json.RawMessage(jsonData))
	result, err := schema.Validate(document)
	if err != nil {
		return fmt.Errorf("error validating JSON schema: %w", err)
	}

	if !result.Valid() {
		errors := result.Errors()
		return fmt.Errorf("JSON schema validation failed: %s", errors)
	}

	return nil
}
func main() {
	// JSON Data Example
	jsonData := []byte(`
	{
		"name": "John",
		"age": 30,
		"email": "john@example.com"
	}
	`)

	// Invalid JSON Data Example
	invalidJsonData := []byte(`
	{
		"name": "J",
		"age": -5,
		"email": "invalidemail"
	}
	`)
	// XML Data Example
	xmlData := []byte(`
	<user>
		<name>Jane</name>
		<age>25</age>
		<email>jane@example.com</email>
	</user>
	`)
	var user User

	// Validate JSON
	err := validateJSON(jsonData)
	if err != nil {
		log.Fatalf("JSON Validation Error: %v", err)
	}
	fmt.Println("JSON Data Validated")

	// Validate Invalid JSON
	err = validateJSON(invalidJsonData)
	if err != nil {
		log.Fatalf("Invalid JSON Validation Error: %v", err)
	}
	fmt.Println("Invalid JSON Data Validated")

	// Validate XML
	err = validateXML(xmlData, &user)
	if err != nil {
		log.Fatalf("XML Validation Error: %v", err)
	}
	fmt.Println("XML Data Validated:", user)
}
