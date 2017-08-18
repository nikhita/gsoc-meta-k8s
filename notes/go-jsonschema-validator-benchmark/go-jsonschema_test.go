package validator_test

import (
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

func BenchmarkGoJSONSchema(b *testing.B) {
	schemaLoader := gojsonschema.NewReferenceLoader("schema4.json")
	documentLoader := gojsonschema.NewReferenceLoader("valid.json")

	for i := 0; i < b.N; i++ {
		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			b.Error(err)
		}
		if !result.Valid() {
			b.Errorf("It is invalid!!!")
		}
	}
}
