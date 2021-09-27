package Models

import (
	// "errors"
	"encoding/json"
	"io/ioutil"
	"os"
)

type EmailTemplate struct {
	From     string
	Subject  string
	MineType string
	Body     string
}

func (t *EmailTemplate) Get(path string) (EmailTemplate, error) {
	var result EmailTemplate

	file, err := os.Open(path)
	if err != nil {
		return result, err
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	_ = json.Unmarshal(byteValue, &result)

	return result, nil
}
