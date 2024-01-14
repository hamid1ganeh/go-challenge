

![Logo](https://media.licdn.com/dms/image/D4D12AQGYvrJwYEpXdw/article-cover_image-shrink_720_1280/0/1665747068873?e=2147483647&v=beta&t=736Ys7pj1I7p_IJB7FviK5RjFUjideODD54Fu16L4XY)

# Go-Challenge
First of all you need to clone the project form current repository and then in order to download project dependencies please run follow command in root of project directory:
```bash
go mod tidy
```
After that copy the .envExample file and past it in same directory and rename it to .env then open it and set up you'r SERVER_PORT,DB_Name, AWS_REGION, AWS_Accesskey and AWS_SecretKey.

If you want to run project unit tests please go to test directory which is located in the root of project and then execute follow command:
```bash
go test
```
## How to deploy
First of all you need to install serverless framewokr on your local machine. If you are using linux OS just need to run follow command:
```bash
npm install -g serverless
```
Now please open th serverless.yaml in line 9 specify you'r aws lambda region and finally it's enough just run below command:
```bash
make deploy
```
After  upload you'r porject on aws lambda service successfully you will recive 2 Endpoint and you can send you'r https request to them. 
 

## Usage/Examples
In order to send https request, you can use below examples:

#### Create Endpoint:
curl --request POST \
  --url https://4bktavifb3.execute-api.us-east-1.amazonaws.com/dev/api/devices \
  --header 'Content-Type: application/json' \
  --data '{
    "id": "/devices/id1",
    "deviceModel": "/devicemodels/id1",
    "name": "Sensor",
    "note": "Testing a sensor.",
    "serial": "A020000102"}'

#### Get Endpoint:
curl --request GET \
  --url https://4bktavifb3.execute-api.us-east-1.amazonaws.com/dev/api/devices/id1

 
## Technologies used:
 - [Golang Programming language](https://golang.org)
 - AWS DynamoDB
 - [Serverless Framework](https://serverless.com)
 - AWS Lambda
 - AWS API Gateway


## Author
- [@Hamid1ganeh2st](https://github.com/hamid1ganeh)



  
 