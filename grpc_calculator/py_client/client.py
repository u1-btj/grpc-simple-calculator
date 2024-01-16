import grpc
import calculator_pb2
import calculator_pb2_grpc

def run():
    try:
        num1 = float(input('Enter first number : '))
        num2 = float(input('Enter second number : '))
        with grpc.insecure_channel("localhost:8080") as channel:
            stub = calculator_pb2_grpc.CalculationStub(channel)
            
            add = stub.Addition(calculator_pb2.NumRequest(num_a=num1, num_b=num2))
            print("Addition : " + remove_trailing_zero(str(add.result)))
            
            sub = stub.Substraction(calculator_pb2.NumRequest(num_a=num1, num_b=num2))
            print("Substraction : " + remove_trailing_zero(str(sub.result)))
            
            mul = stub.Multiplication(calculator_pb2.NumRequest(num_a=num1, num_b=num2))
            print("Multiplication : " + remove_trailing_zero(str(mul.result)))
            
            div = stub.Division(calculator_pb2.NumRequest(num_a=num1, num_b=num2))
            print("Division : " + remove_trailing_zero(str(div.result)))

    except ValueError:
        print('Please input only number')
    
    except grpc.RpcError as e:
        print(f"Rpc Error : {e.details()}")

def remove_trailing_zero(str_num):
    return str_num.rstrip('0').rstrip('.') if '.' in str_num else str_num

if __name__ == '__main__':
    run()