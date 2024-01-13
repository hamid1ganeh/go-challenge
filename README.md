

![Logo](https://media.licdn.com/dms/image/D4D12AQGYvrJwYEpXdw/article-cover_image-shrink_720_1280/0/1665747068873?e=2147483647&v=beta&t=736Ys7pj1I7p_IJB7FviK5RjFUjideODD54Fu16L4XY)

#  Go-Challenge
First of all you need to configure you'r AWS cli and set up you'r Access key,Secret key and Region on you'r locall machine.

After clone the project form current repository, in order to download project dependencies please run follow command in root of project directory:
```bash
go mod tidy
```
Then  copy the .envExample file and rename it to .env then open it and setup you'r SERVER_PORT,DB_REGION and DB_Name.

If you want to run project unit tests please go to test directory wich is located in root of project and then execute follow comand:
```bash
go test
```
if you want to run project and send http request in root of project directory run this command:
```bash
go run .
```
Now please wait until project execute on the SERVER_PORT that you specified in .evn file.

## Usage/Examples
In order to send http request you can use below examples:

#### Create Endpoint:
curl --request POST \
  --url localhost:1370/api/devices \
  --header 'Content-Type: application/json' \
  --data '{
    "id": "/devices/id1",
    "deviceModel": "/devicemodels/id1",
    "name": "Sensor",
    "note": "Testing a sensor.",
    "serial": "A020000102"}'

#### Get Endpoint:
curl --request GET \
  --url localhost:1370/api/devices/id1

## How to deploy

First of all you need to install serverless framework on your local machine. If you are using linux OS just need to run follow comdand:
```bash
npm install -g serverless
```
Now that's enogh just you run below command:
```bash
make deploy
```
After uploaded you'r porject on aws lambda service successfully you will recive 2 Endpoint and you can send you'r http request to them. 

 
## Technologies used:
 - [Golang Programming language](https://golang.org)
 - AWS DynamoDB
 - [Serverless Framework](https://serverless.com)
 - AWS Lambda
 - AWS API Gateway


  
 