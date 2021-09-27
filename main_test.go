package main

import (
	"SendMail/Mailer"
	"SendMail/Models"
	"testing"
)

func TestCustomerLoadFile(t *testing.T) {

	// read customer file
	customer := &Models.Customer{}
	_, err := customer.Get("./SampleData/customers.csv")

	if err != nil {
		t.Fatal("Cannot read Customer file.")
	}

}
func TestTemplateLoadFile(t *testing.T) {

	// read template file
	template := &Models.EmailTemplate{}
	_, err := template.Get("./SampleData/email_template.json")

	if err != nil {
		t.Fatal("Cannot read EmailTemplate file.")
	}

}
func TestMailerConfig(t *testing.T) {

	_, err := Mailer.GetMailer(Mailer.ConfigMailer{
		Type:     "SMTP", // set SMTP or REST API mode on .env file
		ApiKey:   "",
		Host:     "",
		Port:     "",
		User:     "",
		Password: "",
	})

	if err != nil {
		t.Fatal("GetMailer is working wrong.")
	}

}
