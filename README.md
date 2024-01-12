## **Assignment Task: Build a Simple Calculator gRPC Service - (Yuma Anugrah Virya Gunawan)**

## To Do
1. [How To Setup gRPC Project](https://github.com/u1-btj/grpc-simple-calculator/tree/main?tab=readme-ov-file#how-to-run-the-server--client-services)
2. [How To Run The Server & Client Services](https://github.com/u1-btj/grpc-simple-calculator/tree/main?tab=readme-ov-file#how-to-run-the-server--client-services)
3. [Test Case Example](https://github.com/u1-btj/grpc-simple-calculator/tree/main?tab=readme-ov-file#test-case-example)

## How To Setup gRPC Project
### Clone this repo

    git clone https://github.com/u1-btj/grpc-simple-calculator.git

### Make sure golang already installed and properly set
Check [this documentation](https://go.dev/doc/install) to install golang
Verify if golang already properly set

     go version

### Make sure protocol buffer compiler (protoc) already installed and properly set
Check [this documentation](https://grpc.io/docs/protoc-installation/) to install protoc
Verify if protoc already properly set

    protoc --version

### Go inside this repo directory

    cd grpc-simple-calculator


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

