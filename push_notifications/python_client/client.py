import grpc
import topic_pb2
import topic_pb2_grpc
from datetime import datetime

channel = grpc.insecure_channel("localhost:8080")
stub = topic_pb2_grpc.TopicSelectionStub(channel)

def main_service(option):
    try:
        if option == 1:
            meow_facts_service()
        elif option == 2:
            color_info_service()

    # catch value error
    except ValueError:
        print('Please input the correct value type!')
        return 
    
    # catch grpc error
    except grpc.RpcError as e:
        print(f"Rpc Error : {e.details()}")
        return 
    
    # catch unknown exception
    except Exception as e:
        print(f"Exception : {e}")
        return

def meow_facts_service():
    print("\nYou choose Random Interesting Cat Facts Topic")
    count = int(input('Enter total facts per response : '))
    second = int(input('Enter time delay (in second) : '))
    limit = int(input('Enter total response : '))

    # handle input value less than 1
    if count < 1 or second < 1 or limit < 1:
        print('Total facts, time delay, or total responses cannot less than 1!')
        return

    request = topic_pb2.FactRequest(count=count, second=second, limit=limit)
    i = 1

    for meow_facts in stub.StreamMeowFacts(request):
        print(f"{datetime.now()} Response - {i} :")
        print(meow_facts)
        i += 1

def color_info_service():
    print("\nYou choose Color Format Information Topic")
    all_color = []
    total = int(input('Enter total color : '))

    for i in range(1, total + 1):
        color = str(input(f"Enter name for color - {i} : "))
        all_color.append(color)

    second = int(input('Enter time delay (in second) : '))

    # handle input value less than 1
    if total < 1 or second < 1:
        print('Total color or time delay cannot less than 1!')
        return

    request = topic_pb2.ColorRequest(name=all_color, second=second)
    i = 0

    for color_info in stub.StreamColorInfo(request):
        print(f"{datetime.now()} Color - {all_color[i]} Information :")
        print(color_info)
        i += 1

def main_menu():
    print('Choose your requested topic :')
    print('1. Random Interesting Cat Facts')
    print('2. Color Format Information (Hex, RGB, HSL)')
    print('3. Exit')
    num = int(input('Input number: '))

    if num == 1 or num == 2:
        main_service(num)
    elif num == 3:
        print('Thank You')
        return
    else:
        print('Invalid Menu')
        return

if __name__ == '__main__':
    main_menu()