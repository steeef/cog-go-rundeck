NAME=cog-go-rundeck

all: install

clean:
	docker rmi ${NAME} &>/dev/null || true

build:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${NAME} main.go

install: clean build
	docker build -t ${NAME} .
