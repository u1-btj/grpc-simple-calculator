## **Assignment Task: Build a Simple Calculator gRPC Service!**

More detailed objective and requirement from this task, [follow this link](https://gist.github.com/herizsh/042d92f7c489860a08d1b01e669eadcf).  
If you want to build this project from scratch, [follow this step](https://github.com/u1-btj/grpc-simple-calculator/tree/main/grpc_calculator#how-to-build-this-project-from-scratch).  
If you want to clone and run this repo, [follow this step](https://github.com/u1-btj/grpc-simple-calculator/tree/main/grpc_calculator#how-to-setup-this-grpc-project).  

## To Do
1. [How To Build This Project From Scratch](https://github.com/u1-btj/grpc-simple-calculator/tree/main/grpc_calculator#how-to-build-this-project-from-scratch)
2. [How To Setup This gRPC Project](https://github.com/u1-btj/grpc-simple-calculator/tree/main/grpc_calculator#how-to-setup-this-grpc-project)
3. [How To Run The Server & Client Services](https://github.com/u1-btj/grpc-simple-calculator/tree/main/grpc_calculator#how-to-run-the-server--client-services)
4. [Test Case Example](https://github.com/u1-btj/grpc-simple-calculator/tree/main/grpc_calculator#test-case-example)

## How To Build This Project From Scratch
### Make sure golang already installed and properly set
Check [this documentation](https://go.dev/doc/install) to install golang  
Verify if golang already properly set

     go version

### Make sure protocol buffer compiler (protoc) already installed and properly set
Check [this documentation](https://grpc.io/docs/protoc-installation/) to install protoc  
Verify if protoc already properly set

    protoc --version

### Create project directory and go inside that directory

    mkdir grpc-simple-calculator
    cd grpc-simple-calculator/grpc_calculator

### Initialize golang project

    go mod init github.com/u1-btj/grpc-simple-calculator

### Configure protocol compiler and proto file
Install go plugins for protocol compiler

    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

Update path so that protocol compiler can find the plugin

    export  PATH="$PATH:$(go env GOPATH)/bin"

Fill your proto file like [this example](https://github.com/u1-btj/grpc-simple-calculator/blob/main/grpc_calculator/calculator_protoc/calculator.proto)  
Generate protobuffer file 

    protoc --go_out=. --go-grpc_out=. calculator_protoc/calculator.proto

**Note :** When editing the proto file, need to regenerate the protobuffer by execute command above

### Create server and client file
[Server file example](https://github.com/u1-btj/grpc-simple-calculator/blob/main/grpc_calculator/server/main.go)  
[Client file example](https://github.com/u1-btj/grpc-simple-calculator/blob/main/grpc_calculator/client/main.go)  
Make sure the final structure of the project will be like this :
```
+-- calculator_protoc
|   +-- calculator_grpc.pb.go
|   +-- calculator.pb.go
|   +-- calculator.proto
+-- client
|   +-- main.go
+-- server
|   +-- main.go
+-- go.mod
+-- go.sum
```
### Download and validate dependency

    go mod tidy

### Run services
[Run the server & client services](https://github.com/u1-btj/grpc-simple-calculator/tree/main/grpc_calculator#how-to-run-the-server--client-services)

## How To Setup This gRPC Project
### Make sure golang and protoc already installed and properly set
[Verify golang](https://github.com/u1-btj/grpc-simple-calculator/tree/main/grpc_calculator#make-sure-golang-already-installed-and-properly-set)  
[Verify protoc](https://github.com/u1-btj/grpc-simple-calculator/tree/main/grpc_calculator#make-sure-protocol-buffer-compiler-protoc-already-installed-and-properly-set)

### Clone this repo

    git clone https://github.com/u1-btj/grpc-simple-calculator.git

### Go inside this repo directory

    cd grpc-simple-calculator

### Run services
[Run the server & client services](https://github.com/u1-btj/grpc-simple-calculator/tree/main/grpc_calculator#how-to-run-the-server--client-services)

## How To Run The Server & Client Services
### Run the server first

    go run server/main.go
If firewall notification appear, click allow access

### Run the client services
Use default value num = 1 for both number

    go run client/main.go

Or you can use your favorite number as input for num1 and num2

    go run client/main.go --num1=5 --num2=2

## Test Case Example
### num1 = 2 ; num2 = 10

    go run client/main.go --num1=2 --num2=10

Client Output

    $ go run client/main.go --num1=2 --num2=10
    2024/01/12 20:56:09 Addition: 12
    2024/01/12 20:56:09 Substraction: -8
    2024/01/12 20:56:09 Multiplication: 20
    2024/01/12 20:56:09 Division: 0.2

Server Output

    2024/01/12 20:56:09 Num 1: 2
    2024/01/12 20:56:09 Num 2: 10
    2024/01/12 20:56:09 Addition: 12
    2024/01/12 20:56:09 Substraction: -8
    2024/01/12 20:56:09 Multiplication: 20
    2024/01/12 20:56:09 Division: 0.2

### num1 = 5 ; num2 = 2

    go run client/main.go --num1=5 --num2=2

Client Output

    $ go run client/main.go --num1=5 --num2=2
    2024/01/12 20:56:49 Addition: 7
    2024/01/12 20:56:49 Substraction: 3
    2024/01/12 20:56:49 Multiplication: 10
    2024/01/12 20:56:49 Division: 2.5

Server Output

    2024/01/12 20:56:49 Num 1: 5
    2024/01/12 20:56:49 Num 2: 2
    2024/01/12 20:56:49 Addition: 7
    2024/01/12 20:56:49 Substraction: 3
    2024/01/12 20:56:49 Multiplication: 10
    2024/01/12 20:56:49 Division: 2.5

### num1 = 10.5 ; num2 = 4.5

    go run client/main.go --num1=10.5 --num2=4.5

Client Output

    $ go run client/main.go --num1=10.5 --num2=4.5
    2024/01/12 20:57:24 Addition: 15
    2024/01/12 20:57:24 Substraction: 6
    2024/01/12 20:57:24 Multiplication: 47.25
    2024/01/12 20:57:24 Division: 2.3333333333333335

Server Output

    2024/01/12 20:57:24 Num 1: 10.5
    2024/01/12 20:57:24 Num 2: 4.5
    2024/01/12 20:57:24 Addition: 15
    2024/01/12 20:57:24 Substraction: 6
    2024/01/12 20:57:24 Multiplication: 47.25
    2024/01/12 20:57:24 Division: 2.3333333333333335

### num1 = 4 ; num2 = 0

    go run client/main.go --num1=4 --num2=0

Client Output

    $ go run client/main.go --num1=4 --num2=0
    2024/01/12 20:56:59 Addition: 4
    2024/01/12 20:56:59 Substraction: 4
    2024/01/12 20:56:59 Multiplication: 0
    2024/01/12 20:56:59 Division Error: rpc error: code = InvalidArgument desc = Cannot divided by zero
    exit status 1

Server Output

    2024/01/12 20:56:59 Num 1: 4
    2024/01/12 20:56:59 Num 2: 0
    2024/01/12 20:56:59 Addition: 4
    2024/01/12 20:56:59 Substraction: 4
    2024/01/12 20:56:59 Multiplication: 0
    2024/01/12 20:56:59 Division Error: rpc error: code = InvalidArgument desc = Cannot divided by zero

### num1 = abc ; num2 = xyz

    go run client/main.go --num1=2 --num2=10

Client Output

    $ go run client/main.go --num1=abc --num2=xyz
    invalid value "abc" for flag -num1: parse error
    Usage of C:\Users\ASUS\AppData\Local\Temp\go-build3505390063\b001\exe\main.exe:
      -addr string
            Address To Connect (default "localhost:8080")
      -num1 float
            Num 1 (default 1)
      -num2 float
            Num 2 (default 1)
    exit status 2

Server Output

    (Empty, because parsing error already handled in client)

