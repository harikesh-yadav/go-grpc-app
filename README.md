# go-grpc-app
go-grpc-app which contain client and server app in their own directory.


PROTOC Command:

  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  proto/inventory.proto

