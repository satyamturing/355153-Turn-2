// user.go
package models

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

type User struct {
	Name string `json:"name" xml:"name,attr" validate:"required"`
	Age  int    `json:"age" xml:"age,attr" validate:"gte=0,lte=150"`
	Email string `json:"email" xml:"email,attr" validate:"required,email"`
	Address *Address `json:"address" xml:"address"`
}

type Address struct {
	Street string `json:"street" xml:"street,attr"`
	City   string `json:"city" xml:"city,attr"`
}
 