## loop展開による最適化
loopの際に、都度終了条件を走らせるのではなく、loopの単位を大きくして条件分岐命令が入る回数を小さくしよう的なやつ

## Run
```
$ make b

BenchmarkWinding-12             869851530                1.19 ns/op            0 B/op          0 allocs/op
BenchmarkUnWinding10-12         1000000000               1.12 ns/op            0 B/op          0 allocs/op
BenchmarkUnWinding30-12         1000000000               0.978 ns/op           0 B/op          0 allocs/op
```

びみょーにだけど早くなる.全部展開したらそれなりに早くなりそう

## Ref
* https://ja.wikipedia.org/wiki/%E3%83%AB%E3%83%BC%E3%83%97%E5%B1%95%E9%96%8B