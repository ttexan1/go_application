Golang Simple Clean Architecture Application


## Set up

* install postgresql

```
$ createdb golang_practice_development
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
