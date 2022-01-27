## CPUのキャッシュ
* シーケンシャルアクセスの方が早いことを確認する
* HostのL1キャッシュのサイズを見て、L1を当てるか外すかの境界を見る

## Run
L1キャッシュのサイズを調べる
```
$ sysctl hw.l1dcachesize
hw.l1dcachesize: 32768
```
つまり、int64(8byte)の変数は512個までL1キャッシュにのる.
これを踏まえた上でbench.

benchのrun
```
$ make b

BenchmarkMatrixCombination_Sequential_128-12               59952             21025 ns/op               4 B/op          0 allocs/op
BenchmarkMatrixCombination_NotSequential_128-12            54304             22322 ns/op               4 B/op          0 allocs/op
BenchmarkMatrixCombination_Sequential_256-12               15368             77467 ns/op              69 B/op          0 allocs/op
BenchmarkMatrixCombination_NotSequential_256-12             7657            152953 ns/op             138 B/op          0 allocs/op
```

* 128と256の違いから、256あたりでL1 cacheを外してることがわかる
  * 256飛ばしのアクセスでも理論的にはキャッシュに入りそうだけど、なんでだろう？

## Ref
* https://teivah.medium.com/go-and-cpu-caches-af5d32cc5592