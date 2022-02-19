## Hashmapを使いましょう
* 検索する場面が多い場合はhashmap使いましょう的な話(まあ基本だけど)
* 配列の検索と比較して、hashmapの検索の方が速いことを確認する
  * 配列の検索は、線形探索と2分探索どっちもやってみた



## Run
```
$ make b

BenchmarkWinding-12             869851530                1.19 ns/op            0 B/op          0 allocs/op
BenchmarkUnWinding10-12         1000000000               1.12 ns/op            0 B/op          0 allocs/op
BenchmarkUnWinding30-12         1000000000               0.978 ns/op           0 B/op          0 allocs/op
```