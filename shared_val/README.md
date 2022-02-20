## 共有変数のReadとWriteとでパフォーマンスが変わるか
* [この記事](https://pkolaczk.github.io/server-slower-than-a-laptop/)に触発されて、GOで共有変数に対するreadとwriteでパフォーマンスがどれくらい変化するのかを見てみたいと思った

## 結果
* shared_va_test.goにあるようなtestの内容では有意な差はみられなkった
```
345804              3535 ns/op               0 B/op          0 allocs/op
347917              3562 ns/op               0 B/op          0 allocs/op
```

## 考察
* (あとで気づいたけど、)このtestコードの比較だと、そもそも記事の事例との比較になっていない
  * Arcはマルチスレッドでスレッドセーフな参照カウンタとして動くので、writeする際はlockを撮ってるのと変わらない
  * ちなみにこのGoのコードの例は、threadセーフを犠牲にして、好きなタイミングで共有変数にwriteを行ってる.ので、カウンタの値も不正確になっていそう
    * javaみたいにgorouitneって、thread localなデータcacheがあるのかな?
    * javaのメモリモデル復習したい