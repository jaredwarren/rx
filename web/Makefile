
# make

all: test build
build: 
		GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o web-docker
update:
		docker-compose stop web
		docker-compose up --build -d web
