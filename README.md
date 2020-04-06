Convert .proto to .pb.go (run in project's root directory). 

protoc --proto_path=./proto --proto_path=vendor --go_out=plugins=grpc:proto ./proto/*.proto

