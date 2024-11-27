1. Install latest version of Protobuf (https://protobuf.dev/downloads/)  

2. Generate code based on product.proto using the following command:  
    > protoc --proto_path=. --go_out=. ./product.proto

    It will generate file productpb/product.pb.go  

3. Generate code based on productservice.proto using the following command:  
    > protoc --proto_path=. --go_out=. --go-grpc_out=. ./product.proto ./productservice.proto

    It will generate files productpb/productservice.pb.go and productpb/productservice_grpc.pb.go
