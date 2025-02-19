package bench_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
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

func TestDecodeJSONFile(t *testing.T) {
	// Open the JSON file
	file, err := os.Open("testdata/lx-001.json")
	if err != nil {
		t.Fatalf("Failed to open JSON file: %v", err)
	}
	defer file.Close()

	// Create a new JSON decoder
	decoder := json.NewDecoder(file)

	// Decode the JSON data
	var obj []model.Data
	err = decoder.Decode(&obj)
	if err != nil {
		t.Fatalf("Failed to decode JSON data: %v", err)
	}
}
