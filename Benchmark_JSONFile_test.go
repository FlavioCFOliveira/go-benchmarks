package bench_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/FlavioCFOliveira/go-benchmarks/model"
)

func Benchmark_JSONFile_JustDecode(b *testing.B) {

	// Open the JSON file
	file, err := os.Open("testdata/lx-001.json")
	if err != nil {
		b.Fatalf("Failed to open JSON file: %v", err)
	}
	// Create a new JSON decoder
	decoder := json.NewDecoder(file)
	var obj []model.Data

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(&obj)
	}
	b.StopTimer()

	file.Close()
}

func Benchmark_JSONFile_JustUnmarshal(b *testing.B) {
	data, err := os.ReadFile("testdata/lx-001.json")
	if err != nil {
		b.Fatalf("Failed to read JSON file: %v", err)
	}

	var obj []model.Data

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal(data, &obj)
	}
	b.StopTimer()

}
