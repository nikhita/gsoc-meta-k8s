package validator_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"io/ioutil"

	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

type schemaTestT struct {
	Description string       `json:"description"`
	Schema      *spec.Schema `json:"schema"`
	Tests       []struct {
		Description string      `json:"description"`
		Data        interface{} `json:"data"`
		Valid       bool        `json:"valid"`
	}
}

func TestOpenAPI(t *testing.T) {
	file, _ := ioutil.ReadFile("openapi-data.json")
	var testDescriptions []schemaTestT
	json.Unmarshal(file, &testDescriptions)

	for _, testDescription := range testDescriptions {
		// Expand the refs
		err := spec.ExpandSchema(testDescription.Schema, nil, nil /*new(noopResCache)*/)
		if err != nil {
			t.Errorf("should expland clearly: %v", err)
		}

		// create the validator
		validator := validate.NewSchemaValidator(testDescription.Schema, nil, "data", strfmt.Default)

		// validate against each test
		for _, test := range testDescription.Tests {
			result := validator.Validate(test.Data)
			fmt.Println(test.Data)
			if result.AsError() != nil {
				t.Error(result.AsError())
			}
		}

	}
}
