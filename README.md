# bk-end-lambda-api


Prereqs:

Install Go
Insatll AWS CLI(https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)


Stack Used:

AWS CLI
go
AWS Lambda
AWS API Gateway

How to Run:

    Step 1:
    Configure AWS CLI (aws configure)
        Info Needed for configuration:

            AWS Access Key ID: Your access key.
            AWS Secret Access Key: Your secret access key.
            Default Region: The region for your AWS services (used us-east-1).
            Output Format: Choose json, text, or table (default: json).
    Step 2:
    Build an executable of the code from main
        GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
        zip function.zip bootstrap
    
    Step 3:
    Create an AWS Lambda Function and ccollect the arn of the created function "FunctionArn"
       go-lambda % aws lambda create-function \
            --function-name go-lambda-api \
            --runtime provided.al2 \ 
            --role arn:aws:iam::********:role/lambda-admin \
            --handler main \
            --zip-file fileb://function.zip \
            --region us-east-1
<img width="654" alt="Screenshot 2024-11-27 at 2 04 21 PM" src="https://github.com/user-attachments/assets/e47c3419-eac6-4825-8181-4b7a629b6bd1">



    Step 4:
    Create an API Gateway 
        aws apigatewayv2 create-api \
            --name "GoLambdaAPI" \
            --protocol-type HTTP \
            --target arn:aws:lambda:us-east-1:********:function:go-lambda-api

<img width="683" alt="Screenshot 2024-11-27 at 2 01 41 PM" src="https://github.com/user-attachments/assets/c500dd54-2bd4-44ce-ac73-5a2fecec82e2">

            

    Step 5:
    Test using Postman
    endpoint will be created by API Gateway addd the parameter name and your value


    curl --location 'https://ajwc1m7wy0.execute-api.us-east-1.amazonaws.com/greet?name=Pradeep'

    

<img width="655" alt="Screenshot 2024-11-27 at 1 47 29 PM" src="https://github.com/user-attachments/assets/2d60133b-75e4-4372-befc-d4912d9125fb">



