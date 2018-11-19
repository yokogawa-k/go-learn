- [Integration Tests in Go · Philosophical Hacker](https://www.philosophicalhacker.com/post/integration-tests-in-go/)
- [integration](https://stackoverflow.com/questions/25965584/separating-unit-tests-and-integration-tests-in-go)

testing での flags や +build の利用

```console
$ make test
go test -tags=integration -docker
integration test: docker
ok      github.com/yokogawa-k/go-learn/testing-flags    0.001s
go test -tags=integration
integration test: machine
ok      github.com/yokogawa-k/go-learn/testing-flags    0.001s
go test
Normal test
ok      github.com/yokogawa-k/go-learn/testing-flags    0.001s
```

- `go test` を `-tags=integration` をつけて実行すると冒頭に `// +build integration` が追加されているテストだけが実行される
  - `// +build` ではハイフンは使えない
- `go test` で `-docker` のようなオプションをつけた場合 `flags` から制御できる

これらから、integration test と unit test を分けて記述することができる。

また、docker 内で行う CI や手元でのテストと本番環境などで行うテストを同じように記述できる。
