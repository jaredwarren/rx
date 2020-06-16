
# make

all: test build
build: 
		GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o user-docker
update:
		docker-compose stop user
		docker-compose up --build -d user
