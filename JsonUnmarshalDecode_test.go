package bench_test

import (
	"encoding/json"
	"io"
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

func BenchmarkJsonDecode(b *testing.B) {

	// Open the JSON file
	file, err := os.Open("testdata/lx-001.json")
	if err != nil {
		b.Fatalf("Failed to open JSON file: %v", err)
	}

	// Create a new JSON decoder
	decoder := json.NewDecoder(file)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		file.Seek(0, io.SeekStart)

		var obj []model.Data
		decoder.Decode(&obj)
	}
	b.StopTimer()

	file.Close()
}

func BenchmarkJsonUnmarshal(b *testing.B) {

	data, err := os.ReadFile("testdata/lx-001.json")
	if err != nil {
		b.Fatalf("Failed to read JSON file: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var obj []model.Data
		json.Unmarshal(data, &obj)
	}
	b.StopTimer()

}
