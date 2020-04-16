module github.com/zviadm/stats-go/handlers/grpcstats

go 1.14

require (
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/zviadm/stats-go v0.0.3
	google.golang.org/genproto v0.0.0-20200410110633-0848e9f44c36 // indirect
	google.golang.org/grpc v1.28.1
)

replace github.com/zviadm/stats-go => ../../
