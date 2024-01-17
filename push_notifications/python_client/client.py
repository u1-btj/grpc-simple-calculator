import grpc
import topic_pb2
import topic_pb2_grpc
from datetime import datetime

def run():
    try:
        count = int(input('Enter total facts per response : '))
        second = int(input('Enter time delay (in second) : '))
        limit = int(input('Enter total response : '))

        channel = grpc.insecure_channel("localhost:8080")
        stub = topic_pb2_grpc.TopicSelectionStub(channel)
        
        request = topic_pb2.FactRequest(count=count, second=second, limit=limit)
        i = 1

        for meow_facts in stub.StreamMeowFacts(request):
            print(f"{datetime.now()} Response - {i} :")
            print(meow_facts)
            i += 1

    except ValueError:
        print('Please input the correct value!')
    
    except grpc.RpcError as e:
        print(f"Rpc Error : {e.details()}")

if __name__ == '__main__':
    run()