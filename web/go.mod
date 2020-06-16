module github.com/jaredwarren/rx/web

go 1.13

require (
	github.com/gorilla/mux v1.7.4
	github.com/jaredwarren/rx/user-service v0.0.0
	google.golang.org/grpc v1.29.1
)

replace github.com/jaredwarren/rx/user-service v0.0.0 => ../user-service
