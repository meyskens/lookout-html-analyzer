package validator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Validator allows to call a validator service
type Validator struct {
	Endpoint string
}

// NewValidator gives a new instance of Validator
func NewValidator() Validator {
	return Validator{
		Endpoint: "https://validator.w3.org/nu",
	}
}

// ValidateBytes uploads a byte slice to the validator service and returns it's output
func (v *Validator) ValidateBytes(content []byte, contentType string) (Response, error) {
	response, err := http.Post(fmt.Sprintf("%s/?out=json", v.Endpoint), contentType, bytes.NewBuffer(content))
	if err != nil {
		return Response{}, fmt.Errorf("The HTTP request failed with error %s", err)
	}
	data, _ := ioutil.ReadAll(response.Body)

	vResp := Response{}
	err = json.Unmarshal(data, &vResp)

	return vResp, err
}
