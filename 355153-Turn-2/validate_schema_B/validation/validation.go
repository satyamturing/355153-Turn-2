// validation.go
package validation

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/xeipuuv/gojsonschema"
)

func ValidateJSON(jsonData []byte) error {
	// Load the JSON schema from file
	schemaData, err := ioutil.ReadFile("user.schema.json")
	if err != nil {
		return fmt.Errorf("error reading JSON schema file: %w", err)
	}

	// Create a schema loader
	loader := gojsonschema.NewReferenceLoader(gojsonschema.NewBytesLoader(schemaData))

	// Parse the schema document
	schema, err := gojsonschema.NewSchema(loader)
	if err != nil {
		return fmt.Errorf("error parsing JSON schema: %w", err)
	}

	// Create a document loader
	documentLoader := gojsonschema.NewBytesLoader(jsonData)

	// Validate the instance against the schema
	result, err := schema.Validate(documentLoader)
	if err != nil {
		return fmt.Errorf("error validating JSON: %w", err)
	}

	if !result.Valid() {
		errList := result.Errors()
		return errors.New(string(errList.JSON()))
	}

	return nil
}
