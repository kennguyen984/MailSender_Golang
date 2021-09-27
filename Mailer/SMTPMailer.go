package Mailer

import (
	"SendMail/Lib"
	"SendMail/Models"
	"errors"
	"fmt"
	"log"
	"net/smtp"
)

func newSMTPMailer(c ConfigMailer) IMailer {
	return &SMTPMailer{
		config: c,
	}
}

type SMTPMailer struct {
	customers []Models.Customer
	template  Models.EmailTemplate
	config    ConfigMailer
}

func (m *SMTPMailer) SetCustomers(customers []Models.Customer) {
	m.customers = customers
}

func (m *SMTPMailer) SetTemplate(template Models.EmailTemplate) {
	m.template = template
}
func (m *SMTPMailer) SendAll() error {
	if m.config.Host == "" {
		return errors.New("Missing SMTP config.")
	}
	for _, customer := range m.customers {
		if customer.Email != "" {
			// fill customer info to email template
			contentBody := fillCustomerToTemplate(m.template.Body, customer)

			// send email by SMTP protocol
			err := m.Send(customer, m.template, contentBody)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

func (m *SMTPMailer) Send(customer Models.Customer, template Models.EmailTemplate, contentBody string) error {
	config := m.config
	// toList is list of email address that email is to be sent.
	to := []string{customer.Email}
	// This is the message to send in the mail
	from := Lib.Between(template.From, "<", ">")
	msg := []byte("From: " + from + "\r\n" +
		"To: " + customer.Email + "\r\n" +
		"Subject: " + template.Subject + "\r\n\r\n" +
		"Email " + contentBody + "\r\n")

	auth := smtp.PlainAuth("", config.User, config.Password, config.Host)

	err := smtp.SendMail(config.Host+":"+config.Port, auth, from, to, msg)

	// handling the errors
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Sent: ", customer.Email)
	return nil
}
