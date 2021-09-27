package main

import (
	"SendMail/Builder"
	"SendMail/Config"
	"SendMail/Mailer"
	"SendMail/Models"
	"fmt"
	"log"
	"os"
)

func main() {
	pars := getParameters()

	configVars := Config.EnvLoader()

	template := &Models.EmailTemplate{}

	emailTemplate, _ := template.Get(pars["templateFile"])

	// read customer file
	customer := &Models.Customer{}
	customers, _ := customer.Get(pars["customersFile"])

	// report customer missing email address
	export := &Builder.Export{}
	export.SetCustomers(customers)
	export.SetTemplate(emailTemplate)

	// export JSON file to a folder
	export.ToPath(pars["outputEmailFolder"])

	// export invalid customer to CSV file
	export.ErrorToPath(pars["invalidFile"])

	// set config values to Mailer
	mailer, _ := Mailer.GetMailer(Mailer.ConfigMailer{
		Type:     configVars["EMAIL_TYPE"], // set SMTP or REST API mode on .env file
		ApiKey:   configVars["EMAIL_API_KEY"],
		Host:     configVars["EMAIL_HOST"],
		Port:     configVars["EMAIL_PORT"],
		User:     configVars["EMAIL_AUTH_USER"],
		Password: configVars["EMAIL_AUTH_PASSWORD"],
	})

	// set data from input files
	mailer.SetCustomers(customers)
	mailer.SetTemplate(emailTemplate)

	// start to send with customers and content
	mailer.SendAll()

	fmt.Println("SendMail run completed.")
}
func getParameters() map[string]string {
	args := os.Args
	if len(args) <= 4 {
		log.Fatal("Parameters are required.")
	}

	return map[string]string{
		"templateFile":      args[1],
		"customersFile":     args[2],
		"outputEmailFolder": args[3],
		"invalidFile":       args[4],
	}
}
