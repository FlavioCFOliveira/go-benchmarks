package bench_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/FlavioCFOliveira/go-benchmarks/model"
)

func TestUnmarshalJSONFile(t *testing.T) {
	// Read the JSON file
	data, err := ioutil.ReadFile("testdata/lx-001.json")
	if err != nil {
		t.Fatalf("Failed to read JSON file: %v", err)
	}

	// Unmarshal the JSON data
	var obj []model.Data
	err = json.Unmarshal(data, &obj)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON data: %v", err)
	}

}
