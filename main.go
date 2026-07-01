package main

import (
	"fmt"
	"projeto-golang/internal/domain/campaign"

	"github.com/go-playground/validator/v10"
)

func main() {
	contacts := []campaign.Contact{{Emails: ""}}
	campaign := campaign.Campaign{Contacts: contacts}
	validate := validator.New() // instanciar através da função New()
	err := validate.Struct(campaign)
	if err == nil {
		println("Nenhum erro")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {

			switch v.Tag() {
			case "required":
				fmt.Println(v.StructField() + " is required")
			case "min":
				fmt.Println(v.StructField() + " is required with min " + v.Param())
			case "max":
				fmt.Println(v.StructField() + " is required with max " + v.Param())
			case "email":
				fmt.Println(v.StructField() + " is invalid ")
			}

			// println(v.StructField() + " is invalid: " + v.Tag())
		}
	}

}
