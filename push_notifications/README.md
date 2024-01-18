## **Assignment Task: Simple gRPC-based Push Notification Application**

More detailed objective and requirement from this task, [follow this link](https://github.com/herizsh/btj-forecast/blob/main/README.md).  

## API Used
1. Meow Facts : https://meow-facts.netlify.app/
2. Color Information : https://color.serialif.com/

## To Do
1. [Setup Golang Server](https://github.com/u1-btj/grpc-simple-calculator/tree/main/push_notifications#setup-golang-server)
2. [Setup Python Client](https://github.com/u1-btj/grpc-simple-calculator/tree/main/push_notifications#setup-python-client)
3. [Test Case Example](https://github.com/u1-btj/grpc-simple-calculator/tree/main/push_notifications#test-case-example)

## Clone repo and go to project directory
    git clone https://github.com/u1-btj/grpc-simple-calculator.git
    cd grpc-simple-calculator/push_notifications

## Setup Golang Server
**Note:** If you looking to build project from scratch you can follow [this step](https://github.com/u1-btj/grpc-simple-calculator/tree/main/grpc_calculator#how-to-build-this-project-from-scratch)  


### Make sure golang and protoc already installed and properly set
[Verify golang](https://github.com/u1-btj/grpc-simple-calculator/tree/main/grpc_calculator#make-sure-golang-already-installed-and-properly-set)  
[Verify protoc](https://github.com/u1-btj/grpc-simple-calculator/tree/main/grpc_calculator#make-sure-protocol-buffer-compiler-protoc-already-installed-and-properly-set)  

### Generate protobuf
If you need to change the .proto file, you need to regenerate protobuffer file on go_server. On push_notifications directory, use this command:

    protoc --go_out=. --go-grpc_out=. proto/topic.proto

### Run server
On go_server directory, run this command:
    
    go run server.go

## Setup Client Server
### Make sure python and pyenv already installed and properly set
[Install python](https://wiki.python.org/moin/BeginnersGuide/Download)  
[Install pyenv](https://github.com/pyenv/pyenv?tab=readme-ov-file#installation)

### Make sure python version same as project
Check python version
    
    cd python_client
    python -V

If not same, install that version and set on local env

    pyenv install 3.10.5
    pyenv local 3.10.5

### Setup environment

    virtualenv venv

If virtualenv not installed, follow [this link](https://virtualenv.pypa.io/en/latest/installation.html) for installation  

Activate virtualenv

    source venv/Scripts/activate

Install required module based on requirements.txt

    pip install -r requirements.txt

### Generate protobuf on client
If you need to change the .proto file, you need to regenerate protobuffer file on python_client. On python_client directory, use this command:

    python -m grpc_tools.protoc --proto_path=../proto --python_out=. --grpc_python_out=. ../proto/topic.proto

### Run client
On python_client directory, run this command:
    
    go run client.py

## Test Case Example
### Choose Topic - 1 (Random Interesting Cat Facts)
Client Output

    $ python client.py
    Choose your requested topic :
    1. Random Interesting Cat Facts
    2. Color Format Information (Hex, RGB, HSL)
    3. Exit
    Input number: 1

    You choose Random Interesting Cat Facts Topic
    Enter total facts per response : 3
    Enter time delay (in second) : 2
    Enter total response : 3
    2024-01-18 15:46:33.721870 Response - 1 :
    facts: "Cat families usually play best in even numbers. Cats and kittens should be aquired in pairs whenever possible."
    facts: "Cats can not taste sweetness."
    facts: "A group of cats is called a clowder."

    2024-01-18 15:46:36.015218 Response - 2 :
    facts: "The chlorine in fresh tap water irritates sensitive parts of the cat\'s nose. Let tap water sit for 24 hours before giving it to a cat."
    facts: "The Maine Coone is the only native American long haired breed."
    facts: "All cats have three sets of long hairs that are sensitive to pressure - whiskers, eyebrows,and the hairs between their paw pads."      

    2024-01-18 15:46:38.555813 Response - 3 :
    facts: "Some common houseplants poisonous to cats include: English Ivy, iris, mistletoe, philodendron, and yew."
    facts: "Cats that live together sometimes rub each others heads to show that they have no intention of fighting. Young cats do this more often, especially when they are excited."
    facts: "When well treated, a cat can live twenty or more years but the average life span of a domestic cat is 14 years."

Server Output

    $ go run server.go
    2024/01/18 15:46:08 Server listening at [::]:8080
    2024/01/18 15:46:32 Received request to send 3 cat facts every 2 second for 3 times

### Choose Topic - 2 (Color Format Information)
Client Output

    $ python client.py
    Choose your requested topic :
    1. Random Interesting Cat Facts
    2. Color Format Information (Hex, RGB, HSL)
    3. Exit
    Input number: 2

    You choose Color Format Information Topic
    Enter total color : 3
    Enter name for color - 1 : Purple
    Enter name for color - 2 : black
    Enter name for color - 3 : Red
    Enter time delay (in second) : 3
    2024-01-18 15:51:01.702311 Color - Purple Information :
    hex: "#800080"
    rgb: "rgb(128, 0, 128)"
    hsl: "hsl(300, 100%, 25%)"

    2024-01-18 15:51:05.110073 Color - black Information :
    hex: "#000000"
    rgb: "rgb(0, 0, 0)"
    hsl: "hsl(0, 0%, 0%)"

    2024-01-18 15:51:08.503449 Color - Red Information :
    hex: "#ff0000"
    rgb: "rgb(255, 0, 0)"
    hsl: "hsl(0, 100%, 50%)"

Server Output

    $ go run server.go
    2024/01/18 15:46:08 Server listening at [::]:8080
    2024/01/18 15:50:59 Received request to send color information every 3 second for total of 3 colors. Colors : [Purple black Red].
