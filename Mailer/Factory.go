package Mailer

import (
	"SendMail/Lib"
	"SendMail/Models"
	"errors"
	"fmt"
	"time"
)

type ConfigMailer struct {
	Type      string
	ApiKey    string
	Host      string
	Port      string
	User      string
	Password  string
	From      string
	From_name string
}
type IMailer interface {
	SetCustomers(customers []Models.Customer)
	SetTemplate(template Models.EmailTemplate)
	SendAll() error
	Send(customer Models.Customer, template Models.EmailTemplate, contentTemplate string) error
}

func fillCustomerToTemplate(template string, customer Models.Customer) string {
	vars := map[string]string{
		"TITLE":      customer.Title,
		"FIRST_NAME": customer.First_name,
		"LAST_NAME":  customer.Last_name,
		"TODAY":      Lib.DateFormat(time.Now()),
	}
	return Lib.FillVarsToTemplate(template, vars)
}

/* Implement Factory Pattern
Designed so that it can be updated in the future to send emails via SMTP or REST API
without changing the current code base (can add new code)
*/
func GetMailer(c ConfigMailer) (IMailer, error) {
	switch c.Type {
	case "SMTP":
		return &SMTPMailer{config: c}, nil
	case "API":
		return &APIMailer{config: c}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Mailer type %s not recognized.", c.Type))
	}
}
