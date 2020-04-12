module github.com/zviadm/stats-go/handlers/grpcstats

go 1.14

require (
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/zviadm/stats-go/metrics v0.0.0-20200412122026-72fc1da5a98f
	google.golang.org/genproto v0.0.0-20200410110633-0848e9f44c36 // indirect
	google.golang.org/grpc v1.28.1
)

replace github.com/zviadm/stats-go/metrics => ../../metrics