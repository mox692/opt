## Hashmapを使いましょう
* 検索する場面が多い場合はhashmap使いましょう的な話(まあ基本だけど)
* 配列と比較して、検索が速いことを確認する



## Run
```
$ make b

BenchmarkWinding-12             869851530                1.19 ns/op            0 B/op          0 allocs/op
BenchmarkUnWinding10-12         1000000000               1.12 ns/op            0 B/op          0 allocs/op
BenchmarkUnWinding30-12         1000000000               0.978 ns/op           0 B/op          0 allocs/op
```