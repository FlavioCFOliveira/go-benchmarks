# go-benchmarks

```shell
go test -benchmem -run=^$ -bench ^Benchmark github.com/FlavioCFOliveira/go-benchmarks
goos: windows
goarch: amd64
pkg: github.com/FlavioCFOliveira/go-benchmarks
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
Benchmark_JSONFile_JustDecode-8         669732553                1.763 ns/op           0 B/op          0 allocs/op
Benchmark_JSONFile_JustUnmarshal-8           1114               987213 ns/op       27669 B/op       2244 allocs/op
PASS
ok      github.com/FlavioCFOliveira/go-benchmarks       2.770s
```
