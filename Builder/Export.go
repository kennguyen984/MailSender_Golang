package Builder

import (
	"SendMail/Lib"
	"SendMail/Models"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Export struct {
	customers []Models.Customer
	template  Models.EmailTemplate
}

func (e *Export) ToPath(outputPath string) {
	// create folder if not existing
	err := os.MkdirAll(outputPath, 0755)
	if err != nil {
		log.Fatal(err)
	}
	jsonArr := e.Json()
	for _, item := range jsonArr {
		fileName := outputPath + "/" + strings.ReplaceAll(item.To, "@", "") + ".json"

		file, _ := JSONMarshal(item)
		ioutil.WriteFile(fileName, file, 0644)
	}

}
func (e *Export) ErrorToPath(pathFile string) {
	// create folder if not existing
	e.mkdirByPathFile(pathFile)
	// get the invalid customers
	customers := e.getInvalidCustomers()
	file, err := os.Create(pathFile)
	checkError("Cannot create file: ", err)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"TITLE", "FIRST_NAME", "LAST_NAME", "EMAIL"})

	for _, item := range customers {
		writer.Write([]string{item.Title, item.First_name, item.Last_name})

	}
}
func (e *Export) mkdirByPathFile(pathFile string) {
	err := os.MkdirAll(filepath.Dir(pathFile), 0755)
	if err != nil {
		log.Fatal(err)
	}
}
func (e *Export) getInvalidCustomers() []Models.Customer {
	var invalidItems []Models.Customer
	for _, customer := range e.customers {
		if customer.Email == "" {
			invalidItems = append(invalidItems, customer)
		}
	}
	return invalidItems
}
func (e *Export) SetCustomers(customers []Models.Customer) {
	e.customers = customers
}
func (e *Export) SetTemplate(template Models.EmailTemplate) {
	e.template = template
}

func (e *Export) Json() []Models.CustomerEmail {
	var jsonArr []Models.CustomerEmail
	for _, customer := range e.customers {
		// fill customer info to email template
		if customer.Email != "" {

			jsonArr = append(jsonArr, Models.CustomerEmail{
				From:     e.template.From,
				To:       customer.Email,
				Subject:  e.template.Subject,
				MineType: e.template.MineType,
				Body:     fillCustomerToTemplate(e.template.Body, customer),
			})
		}
	}
	return jsonArr
}
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
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
