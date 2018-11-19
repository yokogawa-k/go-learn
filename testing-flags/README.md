- [Integration Tests in Go · Philosophical Hacker](https://www.philosophicalhacker.com/post/integration-tests-in-go/)
- [integration](https://stackoverflow.com/questions/25965584/separating-unit-tests-and-integration-tests-in-go)

testing での flags や +build の利用

`// +build` ではハイフンは使えない


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


## [Mage](https://magefile.org/) を使った場合

### target を指定しない場合

default は未指定

```console
$ mage
Targets:
  test                            is run all test
  testIntegrationDocker           is integration-test with docker
  testIntegrationWithoutDocker    is integration-test without docker
  testNormal                      is normal golang test
```

### test target を指定

```console
$ mage test
Running target: Test
Running dependency: TestNormal
exec: go test
Running dependency: TestIntegrationDocker
exec: go test -tags=integration -docker
Running dependency: TestIntegrationWithoutDocker
exec: go test -tags=integration
integration test: machine
ok      github.com/yokogawa-k/go-learn/testing-flags    0.012s
integration test: docker
ok      github.com/yokogawa-k/go-learn/testing-flags    0.001s
Normal test
ok      github.com/yokogawa-k/go-learn/testing-flags    0.001s
```

