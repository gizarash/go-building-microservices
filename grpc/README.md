1. Install latest version of Protobuf (https://protobuf.dev/downloads/)  
2. Generate code based on product.proto using the following command:  
    > protoc -I=. --go_out=. ./product.proto
