package Mailer

import (
	"SendMail/Models"
	"fmt"
)

type APIMailer struct {
	customers []Models.Customer
	template  Models.EmailTemplate
	config    ConfigMailer
}

func (m *APIMailer) SetCustomers(customers []Models.Customer) {
	m.customers = customers
}

func (m *APIMailer) SetTemplate(template Models.EmailTemplate) {
	m.template = template
}
func (m *APIMailer) SendAll() error {
	for _, customer := range m.customers {
		m.Send(customer, m.template, m.template.Body)
	}
	return nil
}
func (m *APIMailer) Send(customer Models.Customer, template Models.EmailTemplate, contentBody string) error {

	fmt.Println("APIMailer sent: ", customer.Email)
	return nil
}
