setup:
	go get golang.org/x/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help
	go get github.com/jessevdk/go-assets-builder

assets: setup
	go-assets-builder -p=domain -o=domain/assets.go -s="/assets/" assets/

deploy: assets
	GOOS=linux GOARCH=amd64 go build app/main.go
	ssh golang-simple "rm -f ~/simple_server"
	scp -i ~/.ssh/rails-practice.pem main golang-simple:~/simple_server
	ssh golang-simple "./simple_server &"

test:
	go test ./...
