.DEFAULT_GOAL := all

all: buildgo builddocker

buildgo:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cog-go-rundeck main.go

builddocker:
	docker build -t cog-go-rundeck .
