Golang Simple Clean Architecture Application


## Set up

* install postgresql

```
$ createdb golang_practice_development
$ createdb golang_practice_test
$ createdb golang_practice_e2etest

```

## Deploy

* set the private key and

```
$ make deploy
```

## Run
```
$ go run app/main.go
```

## Example
```
$ curl http://localhost:9000/categories -X POST -d '{"name": "category", "display_order": 1}'
```

## Test

```
$ make test
```

## Contents

* 簡単なブログアプリケーションのバックエンドAPIをクリーンアーキテクチャで実装してみました。`Article` `Writer` `Category` の三種類のモデルで設計を意識して書きました。
